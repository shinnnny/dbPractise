package models

import (
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	TagID int `json:"tag_id" gorm:"index"`

	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	State   int    `json:"state"`
}

func (a *Article) GetMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	//maps["deleted_on"] = 0
	maps["delete_at"] = "NULL"
	if a.State != -1 {
		maps["state"] = a.State
	}
	//if a.TagID != -1 {
	//	maps["tag_id"] = a.TagID
	//}

	return maps
}

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_at = ? ", id, "NULL").First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
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
	if err := db.Model(&Article{}).Where("id = ? AND deleted_at = ? ", id, "NULL").Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// Delete delete a single article by id
func Delete(id int) error {
	//if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
	//	return err
	//}
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
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

// GetAll gets a list of articles based on paging constraints
func GetAll(pageNum int, pageSize int, maps interface{}) ([]Article, error) {
	var articles []Article
	err := db.Find(&articles).Error
	//err := db.Preload("Article").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// CleanAllArticle clear all article
func CleanAllArticle() error {
	if err := db.Unscoped().Where("deleted_at != ? ", "NULL").Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}
