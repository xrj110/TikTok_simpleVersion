package tools

import (
	"github.com/RaymondCode/simple-demo/Entry"
)

func CreateTable() {

	DbCon.AutoMigrate(&Entry.User{})
	DbCon.AutoMigrate(&Entry.UserFavorite{})
	DbCon.AutoMigrate(&Entry.Video{})
	DbCon.AutoMigrate(&Entry.LoginInfor{})
	DbCon.AutoMigrate(&Entry.Comment{})

}
