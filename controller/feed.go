package controller

import (
	"dousheng/config"
	"dousheng/gormdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	//FindVideos()
	//fmt.Println(DemoVideos)
	token := c.Query("token")
	if user, exist := usersLoginInfo[token]; exist {
		gormdb.DB.Preload("Author").Find(&DemoVideos)
		//处理视频是否喜欢
		var FavoriteVideoID []int64
		gormdb.DB.Model(&Favorite{}).Where("token = ?", token).Pluck("video_id", &FavoriteVideoID) //查找喜欢的视频ID
		for k := range DemoVideos {
			url := fmt.Sprintf("https://%s/%s", config.WebUrl, DemoVideos[k].PlayURL)
			DemoVideos[k].PlayURL = url
			for _, fid := range FavoriteVideoID {
				if DemoVideos[k].Id == fid {
					DemoVideos[k].IsFavorite = true
				}
			}
		}
		//处理视频作者是否关注
		var toUserIDList []int64
		gormdb.DB.Model(&Relation{}).Where("user_id = ?", user.Id).Pluck("to_user_id", &toUserIDList)
		for k := range DemoVideos {
			if DemoVideos[k].AuthorID == user.Id { //自己必定关注自己
				DemoVideos[k].Author.IsFollow = true
			} else {
				for _, fid := range toUserIDList {
					if DemoVideos[k].AuthorID == fid {
						DemoVideos[k].Author.IsFollow = true
					} else {
						DemoVideos[k].Author.IsFollow = false
					}
				}
			}
		}
	} else {
		gormdb.DB.Preload("Author").Find(&DemoVideos)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
