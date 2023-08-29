package controller

import (
	"dousheng/gormdb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//var tempChat = map[string][]Message{}

//var messageIdSequence = int64(0)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		//chatKey := genChatKey(user.Id, int64(userIdB))

		curMessage := Message{
			FromUserID: user.Id,
			ToUserID:   int64(userIdB),
			Content:    content,
			CreateTime: time.Now().UnixNano() / 1e6,
		}

		err := gormdb.DB.AutoMigrate(&Message{})
		if err != nil {
			return
		}
		gormdb.DB.Create(&curMessage)
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	premsg := c.Query("pre_msg_time")
	preMsg, _ := strconv.ParseInt(premsg, 10, 64)

	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.ParseInt(toUserId, 10, 64)
		var newMsg = make([]Message, 0)
		if preMsg == 0 {
			gormdb.DB.Where("from_user_id = ? AND to_user_id = ? OR from_user_id = ? AND to_user_id = ? ", userIdB, user.Id, user.Id, userIdB).Find(&newMsg)
		} else {
			gormdb.DB.Where("create_time > ? AND from_user_id = ? AND to_user_id = ?", preMsg, userIdB, user.Id).Find(&newMsg)
		}

		c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0, StatusMsg: "消息加载完成"}, MessageList: newMsg})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
