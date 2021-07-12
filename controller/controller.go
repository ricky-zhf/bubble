package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//controller-路由解析与转发
//里面不会有对数据的具体操作，只会调用logic层（类似于service层）中的方法来进行操作。
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	//前端页面填写一个待办事项，提交到此路由。
	//(1)从请求中拉取数据
	var todo models.Todo
	//绑定json
	if err2 := c.ShouldBind(&todo); err2 != nil {
		panic(err2)
	}
	//(2)存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else { //(3)返回响应
		c.JSON(http.StatusOK, todo) //在公司这里的返回需要有一些附加信息。
		//c.JSON(http.StatusOK, gin.H{
		//	"code":2000,
		//	"msg": "success",
		//	"data":todo,
		//})
	}
}

func GetTodoList(c *gin.Context) {
	todoList, err := models.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	//（1）先获取id并判断是否合法
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		//c.json后如果不想让代码继续执行一定要返回
		return
	}
	//2.根据id获取对应的记录，然后赋值给todo变量
	todo, err := models.GetATodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//3。进行更新操作
	if todo.Status {
		todo.Status = false
	} else {
		todo.Status = true
	}
	err = models.UpdateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	//（1）先获取id并判断是否合法
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		//c.json后如果不想让代码继续执行一定要返回
		return
	}
	err := models.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"deletedId": id})
	}
}
