package controller

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id" gorm:"primarykey"` // 视频唯一标识
	AuthorID      int64  `json:"-"`
	Author        User   `json:"author" gorm:"foreignKey:AuthorID;references:Id"`
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

type Comment struct {
	Id         int64  `json:"id"` // 评论id
	UserID     int64  `json:"-"`
	User       User   `json:"user" gorm:"foreignKey:UserID;references:Id"`
	VideoID    int64  `json:"-"`
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

type User struct {
	Id              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

type Account struct {
	ID       int64
	UserID   int64
	Username string
	Password string
	User     User `gorm:"foreignKey:UserID;references:Id"`
}

type Favorite struct {
	Token   string
	VideoID int64
}

type Relation struct {
	Id       int64
	UserId   int64
	ToUserId int64
}

type Message struct {
	ID         int64  `json:"id" gorm:"primarykey"` // 消息id
	FromUserID int64  `json:"from_user_id"`         // 消息发送者id
	ToUserID   int64  `json:"to_user_id"`           // 消息接收者id
	Content    string `json:"content"`              // 消息内容
	CreateTime int64  `json:"create_time"`          // 消息发送时间yyyy-MM-dd HH:MM:ss
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

var DemoVideos = make([]Video, 0)

var DemoComments = make([]Comment, 0)
