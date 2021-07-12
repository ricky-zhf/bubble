package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

//routers
//用于处理路由

func SetupRouter() *gin.Engine {

	gin.ForceConsoleColor()

	gin.DefaultWriter = colorable.NewColorableStdout()
	r := gin.Default()
	//告诉gin去哪拉去模板文件引用的静态文件,即请求路径中/static是去static中去找。
	r.Static("/static", "static")
	//告诉gin去哪里找模板文件。
	r.LoadHTMLGlob("templates/*")
	//括号内的是变量，所以IndexHandler后不能加括号。
	//在go中，当我们将函数本身赋值给某个变量的时候，是不能加括号的
	r.GET("/", controller.IndexHandler)
	//v1版本-使用路由组
	v1Group := r.Group("v1")
	{
		/*添加待办*/
		v1Group.POST("/todo", controller.CreateATodo)
		/*查看待办*/
		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//更新待办
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除待办
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)

	}
	return r
}
