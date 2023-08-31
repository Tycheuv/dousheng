package controller

import (
	"dousheng/config"
	"dousheng/gormdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"strconv"
)

var count = 0
var CoverImg = [...]string{
	"https://cdn.pixabay.com/photo/2023/08/03/19/34/catholic-church-8167850_640.png",
	"https://cdn.pixabay.com/photo/2023/08/11/18/35/flowers-8184126_640.jpg",
	"https://alifei01.cfp.cn/creative/vcg/nowater800/new/VCG211451100675.jpg",
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	err = gormdb.DB.AutoMigrate(&Video{})
	if err != nil {
		return
	}
	user := usersLoginInfo[token]
	//gormdb.DB.Raw("SELECT work_count from users where id = ?", usersLoginInfo[token].WorkCount).Scan(&FavoriteVideoID)
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	// 本地存储路径，视频真实路径
	saveFile := filepath.Join("./public/", finalName)
	// web运行，视频url地址
	simulatorFile := fmt.Sprintf("https://%s/static/%s", config.WebUrl, finalName)
	//coverFile := fmt.Sprintf("/sdcard/Pictures/coverimg/work%d",user.Id) //封面
	err = gormdb.DB.Create(&Video{
		AuthorID:      user.Id,
		Author:        user,
		PlayURL:       simulatorFile,
		CoverURL:      CoverImg[count],
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}).Error
	if err != nil {
		fmt.Println("视频发布失败！")
	}
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		fmt.Println("视频发布失败！")
		return
	}
	count = (count + 1) % 3
	gormdb.DB.Model(&user).UpdateColumn("work_count", gorm.Expr("work_count + ?", 1))
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")
	uid, _ := strconv.ParseInt(userId, 10, 64)
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var myVideos = make([]Video, 0)

	gormdb.DB.Preload("Author").Where("author_id = ?", uid).Find(&myVideos)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: myVideos,
	})
}
