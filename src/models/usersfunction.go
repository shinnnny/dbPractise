package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

func RegisterUser()(map[string]interface{}){
	var user = make(map[string]interface{})
	var username, email string
	var repeatpassword []byte
	var password []byte
	//密钥
	key := []byte("12345678")

	//建立用户
	for{
		fmt.Printf("Please input your username and email: ")
		fmt.Scan(&username, &email)
		//验证用户是否存在
		flag, _ := ExitUserByNameAndEmail(username, email)
		//fmt.Println(flag)
		if flag == true{
			fmt.Printf("This username or email is registered, please choose another one. ")
			continue
		}
		//用户不存在，设置密码
		fmt.Printf("Please input your password: ")
		fmt.Scan(&password)
		//跳出循环
		break
	}

	//重复确认密码
	for  {
		fmt.Printf("Please repeat your password: ")
		fmt.Scan(&repeatpassword)
		match := strings.Compare(string(password), string(repeatpassword))
		if match == 0{
			break
		}
	}
	fmt.Printf("Password match successfully")

	//加密
	cle := []byte(password)
	cip := DesEncrypt(cle, key)
	//fmt.Println(cip)
	c ,_:= json.Marshal(cip)
	password = c

	user["user_id"], user["password"], user["email"], user["state"]= username, password, email, 1
	return user
}

//用户登录
func LoginUser(){
	//var user User
	var username string
	var password []byte
	key := []byte("12345678")

	for{
		fmt.Printf("Please input your username: ")
		fmt.Scan(&username)
		flag, _ := ExitUserByName(username)
		if flag == true{
			user, _ := GetUser(username)
			fmt.Printf("Please input your password: ")
			fmt.Scan(&password)

			//解密
			cip := user.Password
			var a []byte
			json.Unmarshal(cip,&a)
			//fmt.Println(cip, user.Password)
			cle := DesDecrypt(a, key)
			user.Password = cle

			match :=  strings.Compare(string(password), string(user.Password))
			if match == 0 {
				fmt.Printf("Land Successfully")
			}else{
				fmt.Printf("Wrong Password")
			}
		}else{
			//用户未注册
			fmt.Printf("Connot find your username, please input again. ")
			continue
		}
		break
	}

}