package controller

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Entry.Response
	CommentList []Entry.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Entry.Response
	Comment Entry.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	user, _ := c.MustGet("user").(Entry.User)

	actionType := c.Query("action_type")
	if actionType == "1" {
		text := c.Query("comment_text")
		videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{Response: Entry.Response{StatusCode: 2, StatusMsg: "get video ID failed"}})
		}
		status, com := service.CommentAddAction(user.Id, videoId, text)
		c.JSON(http.StatusOK, CommentActionResponse{Response: Entry.Response{StatusCode: status},
			Comment: com})
	} else if actionType == "2" {
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{Response: Entry.Response{StatusCode: 2, StatusMsg: "get comment ID failed"}})
		}
		status := service.CommentDeleteAction(commentId)
		if status != 0 {
			c.JSON(http.StatusOK, CommentActionResponse{Response: Entry.Response{StatusCode: status, StatusMsg: "delete failed"}})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{Response: Entry.Response{StatusCode: status}})
		}
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{Response: Entry.Response{StatusCode: 2, StatusMsg: "get video ID failed"}})
	}
	var comments []Entry.Comment
	comments, err = service.CommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Entry.Response{StatusCode: 1, StatusMsg: "get comment list failed"},
		})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Entry.Response{StatusCode: 0},
			CommentList: comments,
		})
	}

}
