package controller

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Entry.Response
	UserList []Entry.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	user, _ := c.MustGet("user").(Entry.User)
	ToUserId := c.Query("to_user_id")

	if ToUserId == "" {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 1, StatusMsg: "to user id is null"})
	}
	toUserId, _ := strconv.Atoi(ToUserId)
	if user.Id == int64(toUserId) {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 2, StatusMsg: "you can't follow yourself"})
		return
	}
	status, err := service.RelationAction(user, int64(toUserId))
	if err == nil {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: status})
	} else {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 1, StatusMsg: err.Error()})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Entry.Response{
			StatusCode: 0,
		},
		UserList: []Entry.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Entry.Response{
			StatusCode: 0,
		},
		UserList: []Entry.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Entry.Response{
			StatusCode: 0,
		},
		UserList: []Entry.User{DemoUser},
	})
}
