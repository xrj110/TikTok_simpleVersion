package service

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/tools"
	"github.com/jinzhu/gorm"
)

func RelationAction(user Entry.User, ToUserId int64) (int32, error) {
	id := user.Id
	var follow Entry.Follow
	result := tools.DbCon.Where("user_id=? and to_user_id=?", id, ToUserId).Find(&follow)
	tx := tools.DbCon.Begin()
	var ToUser Entry.User
	ToUser.Id = ToUserId
	r := tools.DbCon.Where("id=?", ToUser.Id).Find(&ToUser)
	tools.DbCon.Where("id=?", user.Id).Find(&user)
	if r.Error != nil {
		return 1, r.Error
	}
	if tx.Error != nil {
		panic("transaction error")
	}
	if result.RowsAffected == 0 { //表中无人
		follow.UserId = id
		follow.ToUserId = int64(ToUserId)
		follow.Status = 1
		result = tx.Create(&follow)
		if result.Error != nil {
			return 1, result.Error
		}
		err := updateFollow(user, ToUser, 1, tx)
		if err != nil {
			tx.Rollback()
			panic("transaction rollback")
			return 1, err
		}
	} else { //有人改状态
		follow.Status = -follow.Status
		result = tx.Where("follow_id=?", follow.FollowId).Model(&follow).Update("status", follow.Status)
		if result.Error != nil {
			return 1, result.Error
		}
		err := updateFollow(user, ToUser, int(follow.Status), tx)
		if err != nil {
			tx.Rollback()
			panic("transaction rollback")
			return 1, err
		}
	}
	tx.Commit()
	return 0, nil

}
func updateFollow(user Entry.User, ToUser Entry.User, operation int, tx *gorm.DB) error {
	if operation > 0 {
		user.FollowCount++
		ToUser.FollowerCount++
	} else {
		user.FollowCount--
		ToUser.FollowerCount--
	}
	r1 := tx.Where("id=?", user.Id).Model(&user).Update("follow_count", user.FollowCount)
	if r1.Error != nil {
		return r1.Error
	}
	r1 = tx.Where("id=?", ToUser.Id).Model(&ToUser).Update("follower_count", ToUser.FollowerCount)
	if r1.Error != nil {
		return r1.Error
	} else {
		return nil
	}

}
