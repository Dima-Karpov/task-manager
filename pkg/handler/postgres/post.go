package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task-manager/internal/entities"
	customErrors "task-manager/pkg/errors"
)

// @Summary Create post
// @Security ApiKeyAuth
// @Tags posts
// @Description create post
// @ID create-post
// @Accept  json
// @Produce  json
// @Param input body entities.CreatePostListBody true "list info"
// @Success 200 {object} responses.CreatePostResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/posts [post]
func (h *Handler) create(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input entities.PostList
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.PostList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllPostResponse struct {
	Data []entities.PostList `json:"data"`
}

// @Summary Get All Posts
// @Security ApiKeyAuth
// @Tags posts
// @Description get all posts
// @ID get-all-posts
// @Accept  json
// @Produce  json
// @Success 200 {integer} getAllPostResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/posts [get]
func (h *Handler) getAll(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	posts, err := h.services.PostList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllPostResponse{
		Data: posts,
	})
}

// @Summary Get Post By ID
// @Security ApiKeyAuth
// @Tags posts
// @Description get post by id
// @ID get-post-by-id
// @Accept  json
// @Produce  json
// @Param postId path number true  "ID list"
// @Success 200 {object} entities.PostList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/posts/{postId} [get]
func (h *Handler) getById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	post, err := h.services.PostList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

// @Summary Delete post
// @Security ApiKeyAuth
// @Tags posts
// @Description delete post
// @ID delete-post
// @Accept  json
// @Produce  json
// @Param postId path number true  "ID post"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/posts/{postId} [delete]
func (h *Handler) delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	err = h.services.PostList.Delete(userId, id)
	if err != nil {
		if errors.Is(err, customErrors.ErrPostNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Update post
// @Security ApiKeyAuth
// @Tags posts
// @Description update post
// @ID update-post
// @Accept  json
// @Produce  json
// @Param postId path number true  "ID list"
// @Param data body entities.UpdatePostListBody true "Data for list"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/posts/{postId} [put]
func (h *Handler) update(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	var input entities.UpdatePostInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.PostList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
