package main

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/tools"
)

func main() {
	tools.Init()

	tools.DbCon.AutoMigrate(&Entry.User{})

	//tools.DbCon.AutoMigrate(&controller.Video{})
	//tools.DbCon.AutoMigrate(&controller.Comment{})
	tools.DbCon.AutoMigrate(&Entry.LoginInfor{})

}
