package model

import "fmt"

type Todo struct {
	UID       int    `json:"uid"`
	ID        int    `json:"id" gorm:"praimary_key"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

func CreateTodo(todo *Todo) {
	db.Create(todo)
}

func FindTodos(t *Todo) Todos {
	var todos Todos
	db.Where(t).Find(&todos)
	return todos
}

func DeleteTodo(t *Todo) error {
	// 削除件数が0件だったら引数で指定されたIDが不正だった
	if rows := db.Where(t).Delete(&Todo{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not fild Todo (%v) to delete", t)
	}
	return nil
}

func UpdateTodo(t *Todo) error {
	rows := db.Model(t).Update(map[string]interface{}{
		"name":      t.Name,
		"completed": t.Completed,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", t)
	}
	return nil
}
