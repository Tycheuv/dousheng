package controller

import (
	"dousheng/gormdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserID := c.Query("to_user_id")
	actionType := c.Query("action_type")
	toUserid, _ := strconv.ParseInt(toUserID, 10, 64)
	if user, exist := usersLoginInfo[token]; exist {
		err := gormdb.DB.AutoMigrate(&Relation{})
		if err != nil {
			return
		}
		if actionType == "1" {
			//更新关系列表
			gormdb.DB.Create(&Relation{
				UserId:   user.Id,
				ToUserId: toUserid,
			})
			//更新用户关注数
			gormdb.DB.Model(&user).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1))
			//更新作者粉丝数
			gormdb.DB.Model(&User{Id: toUserid}).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1))
		} else {
			//更新关系列表
			gormdb.DB.Where("user_Id = ? AND to_user_id = ?", user.Id, toUserid).Delete(&Relation{})
			//更新用户关注数
			gormdb.DB.Model(&user).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1))
			//更新作者粉丝数
			gormdb.DB.Model(&User{Id: toUserid}).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1))
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func FollowList(c *gin.Context) {
	userID := c.Query("user_id")
	var toUserIDList []int64 //关注的用户id
	gormdb.DB.Model(&Relation{}).Where("user_id = ?", userID).Pluck("to_user_id", &toUserIDList)
	var toUserList []User
	err := gormdb.DB.Where("id in ?", toUserIDList).Find(&toUserList).Error
	if err != nil {
		fmt.Println(err)
	}
	for k := range toUserList {
		toUserList[k].IsFollow = true
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: toUserList,
	})
}

func FollowerList(c *gin.Context) {
	userID := c.Query("user_id")
	userid, _ := strconv.ParseInt(userID, 10, 64)
	//token := c.Query("token")
	var relationUser []Relation
	gormdb.DB.Where("user_id = ? OR to_user_id = ?", userID, userID).Find(&relationUser)
	var toUserIDList []int64 //关注的用户id
	var UserIDList []int64   //粉丝id
	for k := range relationUser {
		if relationUser[k].UserId == userid {
			toUserIDList = append(toUserIDList, relationUser[k].ToUserId)
		} else {
			UserIDList = append(toUserIDList, relationUser[k].UserId)
		}
	}
	var UserList []User
	gormdb.DB.Where("id in ?", UserIDList).Find(&UserList)
	for k := range UserList {
		for _, id := range toUserIDList {
			if UserList[k].Id == id {
				UserList[k].IsFollow = true
			}
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: UserList,
	})
}

func FriendList(c *gin.Context) {
	userID := c.Query("user_id")
	var frendIdList []int64
	gormdb.DB.Model(&Relation{}).Where("to_user_id = ?", userID).Pluck("user_id", &frendIdList)
	var friendList []User
	gormdb.DB.Where("id in ?", frendIdList).Find(&friendList)
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "好友列表加载完成",
		},
		UserList: friendList,
	})
}
