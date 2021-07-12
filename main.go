package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

//main -> router -> controller -> logic(service) -> model -> dao(只是数据库的初始化、模型绑定等。)
func main() {
	//链接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	//延迟关闭数据库
	defer dao.Close()
	//绑定模型 - 同时也会创建表
	models.InitModel()
	//至此，引入数据库与初始化完成。
	//开始注册路由
	r := routers.SetupRouter()
	// 待办事项
	//添加待办
	//查看待办
	//更新带边
	//删除待办
	r.Run()

}
