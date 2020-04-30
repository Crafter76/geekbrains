package main

import (
	"context"
	"fmt"
	"log"
	"serv/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("geekbrains")
	// запись
	// post := &models.Post{
	// 	Title:   "title2",
	// 	Desc:    "desc2",
	// 	Content: "content2",
	// }
	// _ = post.Insert(ctx, db)
	// fmt.Println(post)

	чтение
	id, _ := primitive.ObjectIDFromHex("5eaa718cf2609ef3f671f478")
	post, err := models.GetPost(ctx, db, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(post.Title)

	// id, _ := primitive.ObjectIDFromHex("5ea9bfd1b58d3761d2ce3783")
	// postDelete := models.Post{
	// 	Mongo: models.Mongo{
	// 		ID: id,
	// 	},
	// }
	// if err := postDelete.Delete(ctx, db); err != nil {
	// 	log.Fatal(err)
	// }

	// id, _ = primitive.ObjectIDFromHex("5eaa718cf2609ef3f671f478")
	// postUpdate := models.Post{
	// 	Mongo: models.Mongo{
	// 		ID: id,
	// 	},
	// 	Title:   "title3",
	// 	Desc:    "desc3",
	// 	Content: "content3",
	// }
	// if err := postUpdate.Update(ctx, db); err != nil {
	// 	log.Fatal(err)
	// }

	// posts, _ := models.GetPosts(ctx, db)
	// fmt.Println(posts)

	// postsFind, _ := models.Find(ctx, db, "title", "title2")
	// fmt.Println(postsFind)
}
