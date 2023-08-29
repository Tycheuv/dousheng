package controller

import (
	"dousheng/gormdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	if user, exist := usersLoginInfo[token]; exist {
		var video = Video{}
		gormdb.DB.Preload("Author").Where("id = ?", videoId).Find(&video)
		err := gormdb.DB.AutoMigrate(&Favorite{})
		if err != nil {
			return
		}
		vId, _ := strconv.ParseInt(videoId, 10, 64)
		if actionType == "1" {
			//更新用户喜欢列表
			gormdb.DB.Create(&Favorite{
				Token:   token,
				VideoID: vId,
			})
			//更新视频获赞数
			gormdb.DB.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
			//更新用户喜欢数
			gormdb.DB.Model(&user).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
			//更新作者获赞数
			gormdb.DB.Model(&video.Author).UpdateColumn("total_favorited", gorm.Expr("total_favorited + ?", 1))
		} else {
			//更新视频获赞数
			gormdb.DB.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
			//更新用户喜欢列表
			gormdb.DB.Where("token = ? AND video_id = ?", token, vId).Delete(&Favorite{})
			//更新用户喜欢数
			gormdb.DB.Model(&user).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
			//更新作者获赞数
			gormdb.DB.Model(&video.Author).UpdateColumn("total_favorited", gorm.Expr("total_favorited - ?", 1))
		}
		if err != nil {
			fmt.Printf("Favorite failure, err:%v\n", err)
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Favorite failure"})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Favorite success"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func FavoriteList(c *gin.Context) {
	//userId := c.Query("user_id")
	token := c.Query("token")
	var FavoriteVideoList = make([]Video, 0, 0)
	var FavoriteVideoID []int64
	gormdb.DB.Model(&Favorite{}).Where("token = ?", token).Pluck("video_id", &FavoriteVideoID)
	//gormdb.DB.Raw("SELECT video_id from favorites where token = ?", token).Scan(&FavoriteVideoID)
	gormdb.DB.Preload("Author").Where("id in ?", FavoriteVideoID).Find(&FavoriteVideoList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: FavoriteVideoList,
	})
}
