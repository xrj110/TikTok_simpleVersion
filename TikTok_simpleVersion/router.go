package main

import (
	"github.com/RaymondCode/simple-demo/Middleware"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	r.LoadHTMLGlob("templates/*")

	// home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", Middleware.TokenAuthMiddleware(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", Middleware.TokenAuthMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", Middleware.TokenAuthMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", Middleware.TokenAuthMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", Middleware.TokenAuthMiddleware(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", Middleware.TokenAuthMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", Middleware.TokenAuthMiddleware(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
