package util

import (
	"dbPractise/models"
	"log"
)

func AddArticle(data map[string]interface{}) {
	article := models.Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}

	if err := models.Add(article); err != nil {
		log.Println(err.Error())
	}
}
