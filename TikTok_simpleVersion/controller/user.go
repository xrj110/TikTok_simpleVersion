package controller

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]Entry.User{}

var userIdSequence = int64(1)

/*
tony:后期使用redis作为登录缓存
*/
type UserLoginResponse struct {
	Entry.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Entry.Response
	User Entry.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	result := service.Register(username, password)
	if result == 0 {
		userIdSequence++
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})

	} else if result == 1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 1, StatusMsg: "Username already exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 1, StatusMsg: "unknown error"},
		})

	}

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	result := service.Login(username, password)
	if result.Id != 0 {
		token := username + password
		usersLoginInfo[token] = result
		//c.SetCookie("token", token, 3600, "/", "localhost", false, false)

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 0},
			UserId:   result.Id,
			Token:    token})

	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Entry.Response{StatusCode: 1, StatusMsg: "Username or Password Incorrect"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Entry.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Entry.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
