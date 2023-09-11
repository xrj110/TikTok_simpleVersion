package main

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	tools.ReadConfig()
	go service.RunMessageServer()

	r := gin.Default()

	InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
