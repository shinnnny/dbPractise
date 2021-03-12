package util

import (
	"dbPractise/models"
	"log"
)

func AddArticle(data map[string]interface{}) {
	article := models.Article{
		TagID:   data["tag_id"].(int),
		Title:   data["title"].(string),
		Desc:    data["desc"].(string),
		Content: data["content"].(string),
		State:   data["state"].(int),
	}

	if err := models.Add(article); err != nil {
		log.Println(err.Error())
	}
}

func EditArticle(id int, data map[string]interface{}) {
	if err := models.Edit(id, data); err != nil {
		log.Println(err.Error())
	}
}

func DeleteArticle(id int) {
	if err := models.Delete(id); err != nil {
		log.Println(err.Error())
	}
}

func GetArticle(id int) *models.Article {
	if article, err := models.Get(id); err != nil {
		log.Println(err.Error())
		return nil
	} else {
		return article
	}
}

func GetArticles() []*models.Article {
	article := models.Article{
		State: 1,
	}
	if articles, err := models.GetAll(1, 10, article.GetMaps()); err != nil {
		log.Println(err.Error())
		return nil
	} else {
		return articles
	}
}
