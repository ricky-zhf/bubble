package models

import (
	"bubble/dao"
)

//将定义的模型放在这个目录中
//其次，对于模型所有的原子操作，即增删改查等操作也放在此处
//Todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func InitModel() {
	dao.DB.AutoMigrate(&Todo{})
}

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return err
}

func GetTodoList() (todoList []Todo, err error) { //
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	} else {
		return
	}
}

func GetATodoById(id string) (todo *Todo, err error) {
	todo = new(Todo) //此处要加这句话，否则无法更新。？？
	err = dao.DB.Where("id=?", id).First(todo).Error
	if err != nil {
		return nil, err
	} else {
		return
	}
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return nil
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}
