package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Entry.Response
	VideoList []Entry.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTime := c.Query("latest_time")
	fmt.Print("---kk-------lastTime", lastTime)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Entry.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
