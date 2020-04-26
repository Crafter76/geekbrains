package main

import (
	"fmt"
	"log"
	"orm/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// db, err := gorm.Open("mysql", "mysql:root@/beego?parseTime=true")
	db, err := gorm.Open("mysql", "mysql:root@/gorm?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()

	// task := models.Task{
	// 	Text: "Text from GORM",
	// }
	// db.Create(&task)

	// db.Delete(&models.Task{}, "id = 4")

	task := models.Task{}
	db.First(&task, "id = 3")
	fmt.Println(task.Text)

	// tasks := []models.Task{}
	// db.Find(&tasks)
	// for _, t := range tasks {
	// 	fmt.Println(t)
	// }

}
