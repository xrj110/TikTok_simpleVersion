package service

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/tools"
)

// valid 1-favorite 2-cancel favorite
func FavoriteAction(favorite Entry.UserFavorite) bool {
	tools.DbCon.Where("user_id=? AND video_id=?", favorite.UserID, favorite.VideoID).First(&favorite)

	if favorite.Valid == 0 { //create
		favorite.Valid = 1
		result := tools.DbCon.Create(&favorite)
		if result.Error != nil {
			panic("crate a favorite user error")
			return false
		}
	} else if favorite.Valid == 1 {
		result := tools.DbCon.Where("user_id=? AND video_id=?", favorite.UserID, favorite.VideoID).Model(&favorite).Update("valid", 2)
		if result.Error != nil {
			panic("model favorite user error")
			return false
		}
	} else {
		result := tools.DbCon.Where("user_id=? AND video_id=?", favorite.UserID, favorite.VideoID).Model(&favorite).Update("valid", 1)
		if result.Error != nil {
			panic("model favorite user error")
			return false
		}
	}
	return true
}
func FavoriteList(userId int64) ([]Entry.Video, error) {
	var favorites []Entry.UserFavorite
	result := tools.DbCon.Where("user_id=? AND valid=?", userId, 1).Find(&favorites)
	if result.Error != nil {
		return nil, result.Error
	}
	var videosIds []int64
	for _, favorites := range favorites {
		videosIds = append(videosIds, favorites.VideoID)
	}
	var videos []Entry.Video
	tools.DbCon.Where("id IN (?)", videosIds).Find(&videos)
	for i := range videos {
		SetAuthor(&videos[i])
	}
	return setVideoURL(videos), nil

}
