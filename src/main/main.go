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
	//	"content": "modified by Edit",
	//}
	//util.EditArticle(20,article)

	// Delete
	//util.DeleteArticle(1)

	// Get
	//article:=util.GetArticle(20)
	//log.Println("\t[GET]:\t",article.ID,"\t",article.CreatedAt)

	// Get all
	articles := util.GetArticles()
	log.Println("\t[GET ALL]")
	for _, article := range articles {
		log.Println("\t", article.ID, "\t", article.CreatedAt)
	}

	// Clean
	//err:=models.CleanAllArticle()
	//if err != nil {
	//	log.Println(err)
	//}
}
