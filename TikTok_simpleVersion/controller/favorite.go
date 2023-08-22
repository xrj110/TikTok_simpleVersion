package controller

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	user, _ := c.MustGet("user").(Entry.User)
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 2, StatusMsg: "invalid video id"})
	}
	res := service.FavoriteAction(Entry.UserFavorite{
		UserID:  user.Id,
		VideoID: videoId,
	})
	if res {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 1, StatusMsg: "favorite failed"})
	}

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	user, _ := c.MustGet("user").(Entry.User)
	videoList, err := service.FavoriteList(user.Id)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Entry.Response{
				StatusCode: 1,
				StatusMsg:  "get favorite list error"}})

	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Entry.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
