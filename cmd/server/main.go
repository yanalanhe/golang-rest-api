package main

import (	
	"fmt"

	"github.com/yanalanhe/golang-rest-api/internal/comment"
	"github.com/yanalanhe/golang-rest-api/internal/db"
	transportHttp "github.com/yanalanhe/golang-rest-api/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting up the application")

	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfull connected database")

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		fmt.Println(err)
	}
	
	/* cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID: "7f3b7221-e326-4c77-902d-aa74169e915c",
			Slug: "manual-test",
			Author: "Alan",
			Body: "A test comment",
		},
	) */
	
	/* fmt.Println(cmtService.GetComment(
		context.Background(), 
		"c068bd76-74cb-11ed-a1eb-0242ac120002",
	)) */
/* 
	fmt.Println(cmtService.UpdateComment(
		context.Background(),
		"c206606a-361c-4942-ba97-4f2de7f8464a",
		comment.Comment{		
			Slug: "manual-test-updated",
			Author: "Alan-updated",
			Body: "A test comment - updated",
		},
	)) */

	/* cmtService.DeleteComment(
		context.Background(),
		"c068bd76-74cb-11ed-a1eb-0242ac120002",
	) */

	return nil
}

func main() {
	fmt.Println("Golang REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
