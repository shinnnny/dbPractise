package main

import (
	"dbPractise/models"
	"dbPractise/pkg/logging"
	"dbPractise/setting"
	"dbPractise/util"
	"log"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()

	// Add
	//article := map[string]interface{}{
	//	"tag_id":  1,
	//	"title":   "test",
	//	"desc":    "add for testing",
	//	"content": "whatever",
	//	"state":   1,
	//}
	//util.AddArticle(article)

	//Edit
	//for i:=1;i<30;i++{
	//	article:=map[string]interface{}{
	//		"content": "modified by Edit",
	//	}
	//	util.EditArticle(i,article)
	//}

	// Delete
	//util.DeleteArticle(26)

	//Get
	article := util.GetArticle(4)
	log.Println("\t[GET]:\t", article.ID, "\t", article.CreatedAt)

	//Get all
	//articles := util.GetArticles(1)
	//log.Println("\t[GET ALL]")
	//for _, article := range articles {
	//	log.Println("\t", article.ID, "\t", article.CreatedAt)
	//}

	// Clean
	//util.CleanArticles()

	// Exist
	//exist := util.ExistArticleByID(25)
	//log.Println(exist)
}
