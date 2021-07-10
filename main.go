package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()
	r := gin.Default()
	//告诉gin去哪拉去模板文件引用的静态文件,即请求路径中/static是去static中去找。
	r.Static("/static", "static")
	//告诉gin去哪里找模板文件。
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 待办事项
	//添加待办
	//查看待办
	//更新带边
	//删除待办
	r.Run()

}
