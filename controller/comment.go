package controller

import (
	"dousheng/gormdb"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

var commentIdSequence int64

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	vID := c.Query("video_id")
	videoID, _ := strconv.ParseInt(vID, 10, 64)
	if user, exist := usersLoginInfo[token]; exist {
		err := gormdb.DB.AutoMigrate(&Comment{})
		if err != nil {
			return
		}
		gormdb.DB.Find(&Comment{}).Count(&commentIdSequence)
		atomic.AddInt64(&commentIdSequence, 1)
		var Comm Comment
		if actionType == "1" {
			text := c.Query("comment_text")
			//向数据库表中添加评论
			Comm = Comment{
				Id:         commentIdSequence,
				VideoID:    videoID,
				UserID:     user.Id,
				User:       user,
				Content:    text,
				CreateDate: time.Now().Format("01-02"),
			}
			gormdb.DB.Create(&Comm)
			//更新视频评论数
			gormdb.DB.Preload("Author").Model(&Video{}).Where("id = ?", videoID).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0},
				Comment:  Comm,
			})
			return
		} else {
			cid := c.Query("comment_id")
			gormdb.DB.Preload("User").Where("id = ?", cid).Delete(&Comment{}) //删除评论
			//更新视频评论数
			gormdb.DB.Preload("Author").Model(&Video{}).Where("id = ?", videoID).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
		}
		c.JSON(http.StatusOK, CommentActionResponse{Response{StatusCode: 0, StatusMsg: "Comment action success"}, Comm})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func CommentList(c *gin.Context) {
	videoID := c.Query("video_id")
	gormdb.DB.Preload("User").Where("video_id = ?", videoID).Find(&DemoComments)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
