package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	TagID int `json:"tag_id" gorm:"index"`

	Title   string    `json:"title"`
	Desc    string    `json:"desc"`
	Content string    `json:"content"`
	State   int       `json:"state"`
	AddTime time.Time `gorm:"autoCreateTime"`
}

func (a *Article) GetMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.State != -1 {
		maps["state"] = a.State
	}

	return maps
}

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil || err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

// Add adds a single article
func Add(article Article) error {
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

// Edit modify a single article
func Edit(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// Delete delete a single article by id
func Delete(id int) error {
	if err := db.Delete(&Article{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Get Get a single article based on ID
func Get(id int) (*Article, error) {
	var article Article
	//err := db.Where("id = ? AND deleted_at = ? ", id, "NULL").First(&article).Error
	err := db.Where(id).First(&article).Error
	if err != nil || err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func getPage(pageNum int) int {
	if pageNum > 0 {
		result := (pageNum - 1) * pageSize
		return result
	} else {
		log.Println("\t[Error]:\tPage Number error.")
	}
	return 0
}

// GetAll gets a list of articles based on paging constraints
func GetAll(pageNum int, maps interface{}) ([]*Article, error) {
	var articles []Article
	//err := db.Find(&articles).Error
	err := db.Where(maps).Limit(pageSize).Offset(getPage(pageNum)).Find(&articles).Error
	//err := db.Preload("Article").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil || err != gorm.ErrRecordNotFound {
		return nil, err
	}

	temp := make([]*Article, len(articles))
	for aIndex, a := range articles {
		temp[aIndex] = &a
	}
	return temp, nil
}

// CleanAllArticle clear all article
func CleanAllArticle() error {
	if err := db.Unscoped().Where("deleted_at != ? ", "0000-00-01 00:00:01").Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
