package model

import "github.com/labstack/echo"

type Todo struct {
	Id       int    `db:"id" json:"id"`
	Task     string `db:"task" json:"task"`
	Finished bool   `db:"finished" json:"finished"`
}

func NewTodo(c echo.Context) (*Todo, error) {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return nil, err
	}
	return todo, nil
}
