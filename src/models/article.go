package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// Add adds a single article
func Add(article Article) error {
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

// GetAll gets a list of articles based on paging constraints
func GetAll(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Article").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// CleanAllArticle clear all article
func CleanAllArticle() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}
