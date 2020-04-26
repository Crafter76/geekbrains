package main

import (
	"boiler/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	db, _ := sql.Open("mysql", "mysql:root@/beego")
	boil.SetDB(db)
	boil.DebugMode = true

	ctx := context.Background()

	tasks, _ := models.Tasks().All(ctx, db)
	for _, t := range tasks {
		fmt.Println(t)
	}

	task := &models.Task{
		Text: "boil text user",
	}

	user := &models.User{
		Name: "Ivan",
	}

	if err := task.SetUser(ctx, db, true, user); err != nil {
		log.Fatal(err)
	}

	if err := task.Insert(ctx, db, boil.Infer()); err != nil {
		log.Fatal(err)
	}

}
