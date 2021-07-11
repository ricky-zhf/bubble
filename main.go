package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mattn/go-colorable"
	"net/http"
)

//Todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//为了项目使用方便，有些变量可以定义为全局。
var (
	DB *gorm.DB
)

//链接mysql
func initMySQL() (err error) {
	dsn := "root:12321@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	//链接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	//延迟关闭数据库
	defer DB.Close()
	//绑定模型 - 同时也会创建表
	DB.AutoMigrate(&Todo{})
	//至此，引入数据库与初始化完成。
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
	//v1版本-使用路由组
	v1Group := r.Group("/v1")
	{
		/*添加待办*/
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写一个待办事项，提交到此路由。
			//(1)从请求中拉取数据
			var todo Todo
			//绑定json
			if err2 := c.ShouldBind(&todo); err2 != nil {
				panic(err2)
			}
			//(2)存入数据库
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else { //(3)返回响应
				c.JSON(http.StatusOK, todo) //在公司这里的返回需要有一些附加信息。
				//c.JSON(http.StatusOK, gin.H{
				//	"code":2000,
				//	"msg": "success",
				//	"data":todo,
				//})
			}

		})
		/*查看待办*/
		//一、查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//二、查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//更新待办
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			//（1）先获取id并判断是否合法
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
				//c.json后如果不想让代码继续执行一定要返回
				return
			}
			//2.根据id获取对应的记录，然后赋值给todo变量
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			//3。进行更新操作
			if todo.Status {
				todo.Status = false
			} else {
				todo.Status = true
			}
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除待办
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			//（1）先获取id并判断是否合法
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				//c.json后如果不想让代码继续执行一定要返回
				return
			}
			if err = DB.Where("id = ?", id).Delete(&Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"deletedId": id})
			}
		})

	}
	// 待办事项
	//添加待办
	//查看待办
	//更新带边
	//删除待办
	r.Run()

}
