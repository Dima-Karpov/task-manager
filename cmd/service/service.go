package service

import (
	"fmt"
	"log"
	"os"
	"task-manager/pkg/storage/memdb"
	postgres "task-manager/pkg/storage/postgres"
)

func main() {
	var err error
	pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}
	// TODO
	connstr := "postgres://" + pwd + "@localhost:5434/tasks"
	_, err = postgres.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	dbM := memdb.DB{}
	tasks, err := dbM.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d tasks\n", tasks)
}
