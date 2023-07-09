package models

import (
	"log"
	"time"
)

type Todo struct {
	ID int
	Content string
	UserId int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (content, user_id, created_at) VALUES (?, ?, ?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserId,
		&todo.CreatedAt,
	)
	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserId, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	return todos, err
}

func (u *User) GetTodosByUserId(user_id int) (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos WHERE user_id = ?`
	rows, err := Db.Query(cmd, user_id)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserId, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	return todos, err
}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `UPDATE todos SET content = ? WHERE id = ?`
	_, err = Db.Exec(cmd, t.Content, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `DELETE FROM todos WHERE id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
