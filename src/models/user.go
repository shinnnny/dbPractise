package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserId       string `json:"user_id"`
	Password 	 []byte `json:"password"`
	Email   	 string `json:"email"`
	State 	 	 int 	`json:"state"`
}

func (b *User) GetUserMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["delete_at"] = "Null"
	if b.State != -1 {
		maps["state"] = b.State
	}
	return maps
}

func AddUser(user User) error {
	if err := db.Create(&user).Error; err != nil{
		return err
	}
	return nil
}

func GetUser(UserId string) (*User, error) {
	var user User

	err := db.Where("user_id = ?", UserId).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, err
	}
	return &user, nil
}

func ExitUserByNameAndEmail(username, email string)(bool, error){
	var user User

	err := db.Select("user_id, email").Where("user_id = ? OR email = ?", username, email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return false, err
	}
	//fmt.Println(user)
	if user.UserId != "" || user.Email != ""{
		return true, nil
	}
	return false, nil
}

func ExitUserByName(username string)(bool, error){
	var user User

	err := db.Select("user_id").Where("user_id = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return false, err
	}

	if user.UserId != ""{
		return true, nil
	}
	return false, nil
}


