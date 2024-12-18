package postrges

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"task-manager/internal/entities"
	customErrors "task-manager/pkg/errors"
	"time"
)

type PostListPostgres struct {
	db *sqlx.DB
}

func NewPostListPostgres(db *sqlx.DB) *PostListPostgres {
	return &PostListPostgres{db: db}
}

func (r *PostListPostgres) Create(userId int, post entities.PostList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	currentTime := time.Now()
	createListQuery := fmt.Sprintf(
		"INSERT INTO %s (title, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		postsListsTable,
	)
	row := tx.QueryRow(createListQuery, post.Title, post.Content, currentTime, currentTime)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, post_id) VALUES ($1, $2)",
		usersListsTable,
	)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *PostListPostgres) GetAll(userId int) ([]entities.PostList, error) {
	var posts []entities.PostList

	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.content, tl.created_at, tl.updated_at 
					FROM %s tl INNER JOIN %s ul on tl.id = ul.post_id WHERE ul.user_id = $1`,
		postsListsTable,
		usersListsTable,
	)
	err := r.db.Select(&posts, query, userId)

	return posts, err
}

func (r *PostListPostgres) GetById(userId, id int) (entities.PostList, error) {
	var post entities.PostList
	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.content, tl.created_at, tl.updated_at FROM %s tl 
					INNER JOIN %s ul on tl.id = ul.post_id WHERE ul.user_id = $1 AND ul.post_id = $2`,
		postsListsTable,
		usersListsTable,
	)
	err := r.db.Get(&post, query, userId, id)

	return post, err
}

func (r *PostListPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s tl USING %s ul WHERE tl.id = ul.post_id AND ul.user_id = $1 AND ul.post_id = $2",
		postsListsTable,
		usersListsTable,
	)
	result, err := r.db.Exec(query, userId, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return customErrors.ErrPostNotFound
	}

	return err
}

func (r *PostListPostgres) Update(userId, id int, input entities.UpdatePostInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	setValues = append(setValues, "updated_at=NOW()")
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(
		"UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.post_id AND ul.post_id=$%d AND ul.user_id=$%d",
		postsListsTable,
		setQuery,
		usersListsTable,
		argId,
		argId+1,
	)

	args = append(args, id, userId)

	logrus.Warnf("updateQuery: %s", query)
	logrus.Warnf("args: %v", args)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		logrus.Error("No rows were updated")
		return err
	}

	return err
}
