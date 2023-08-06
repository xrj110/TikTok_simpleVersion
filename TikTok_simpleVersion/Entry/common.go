package Entry

import (
	"gorm.io/gorm"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	gorm.Model
	Id            int64 `json:"id,omitempty"`
	Author        User  `json:"author" gorm:"foreignKey:UserID"`
	UserID        int64
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	gorm.Model
	Id         int64  `json:"id,omitempty"`
	UserID     int64  `gorm:"foreignKey:UserID"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	gorm.Model
	Id              int64  `json:"id,omitempty" gorm:"primaryKey"`
	Name            string `json:"name,omitempty"`
	FollowCount     int64  `json:"follow_count,omitempty"`
	FollowerCount   int64  `json:"follower_count,omitempty"`
	IsFollow        bool   `json:"is_follow,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	Signature       string `json:"signature,omitempty"`
	TotalFavorited  int64  `json:"total_favorited,omitempty"`
	WorkCount       int64  `json:"work_count,omitempty"`
	FavoriteCount   int64  `json:"favorite_count,omitempty"`
}

type Message struct {
	gorm.Model
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}
type LoginInfor struct {
	UserID    int64  `json:"user_id,omitempty"gorm:"primary_key"`
	UserName  string `json:"userName,omitempty" "`
	Password  string `json:"password,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" sql:"index"`
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