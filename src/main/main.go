package main

import (
	"dbPractise/models"
	"dbPractise/setting"
	"dbPractise/util"
	"log"
)

func main() {
	setting.Setup()
	models.Setup()

	// Add
	//article := map[string]interface{}{
	//	"tag_id":  1,
	//	"title":   "test",
	//	"desc":    "add for testing",
	//	"content": "whatever",
	//	"state":   1,
	//}
	//util.AddArticle(article)

	// Edit
	//article:=map[string]interface{}{
	//	"content": "modified",
	//}
	//util.EditArticle(18,article)

	// Delete
	//util.DeleteArticle(1)

	// Get
	//article:=util.GetArticle(20)
	//log.Println("[GET]:\t",article.ID,"\t",article.CreatedAt)

	// Get all
	articles := util.GetArticles()
	log.Println("[GET ALL]:\t", articles)
}
