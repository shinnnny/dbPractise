package main

import (
	"dbPractise/models"
	"dbPractise/setting"
	"dbPractise/util"
)

func main() {
	setting.Setup()
	models.Setup()

	article := map[string]interface{}{
		"tag_id":     1,
		"title":      "test",
		"desc":       "add for testing",
		"content":    "whatever",
		"created_by": "lzt",
		"state":      1,
	}
	util.AddArticle(article)
}
