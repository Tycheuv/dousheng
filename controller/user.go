package controller

import (
	"dousheng/gormdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts

var usersLoginInfo = make(map[string]User, 10)

var userIdSequence int64

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func AccountUserInfo() {
	var AccountInfo []Account
	gormdb.DB.Preload("User").Find(&AccountInfo)
	for _, v := range AccountInfo {
		usersLoginInfo[v.Username+v.Password] = v.User
	}
	userIdSequence = int64(len(AccountInfo))
	//fmt.Println(usersLoginInfo)
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	AccountUserInfo() //读取所有用户信息

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		signature := fmt.Sprintf("测试用户-%d", userIdSequence)
		newUser := User{
			Id:              userIdSequence,
			Name:            username,
			Avatar:          "https://cdn.pixabay.com/photo/2023/08/12/13/59/cat-8185712_640.jpg",
			BackgroundImage: "https://alifei01.cfp.cn/creative/vcg/nowater800/new/VCG211451041048.jpg",
			FavoriteCount:   0,
			TotalFavorited:  "0",
			WorkCount:       0,
			Signature:       signature,
		}
		usersLoginInfo[token] = newUser
		err := gormdb.DB.AutoMigrate(&User{}, &Account{})
		if err != nil {
			return
		}
		gormdb.DB.Create(&Account{ID: userIdSequence, Username: username, Password: password, User: newUser})
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Registered successfully!"},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password
	AccountUserInfo()
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Successful landing!"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "OK"},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
