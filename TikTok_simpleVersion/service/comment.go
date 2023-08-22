package service

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/tools"
	"time"
)

func CommentAddAction(userId int64, videoId int64, context string) (int32, Entry.Comment) {

	com := Entry.Comment{
		VideoId:    videoId,
		UserID:     userId,
		Content:    context,
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	result := tools.DbCon.Create(&com)
	if result.Error != nil {
		panic("create comment failed")
		return 1, com
	}
	return 0, com
}
func CommentDeleteAction(commentId int64) int32 {
	result := tools.DbCon.Delete(Entry.Comment{}, commentId)
	if result.Error != nil {
		panic("delete comment failed")
		return 1
	} else {
		return 0
	}

}
func CommentList(videoId int64) ([]Entry.Comment, error) {
	var comments []Entry.Comment
	result := tools.DbCon.Where("video_id=?", videoId).Find(&comments)
	for i := range comments {
		var user Entry.User
		result := tools.DbCon.Where("id=?", comments[i].UserID).Find(&user)
		if result.Error != nil {
			panic("get user error when get the comment list ")
			return comments, result.Error
		}
		comments[i].User = user

	}
	if result.Error != nil {
		return comments, result.Error
	} else {
		return comments, nil
	}
}
