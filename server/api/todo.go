package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/saiki/golang-echo-sqlx-todo/server/database"
	"github.com/saiki/golang-echo-sqlx-todo/server/middleware"
	"github.com/saiki/golang-echo-sqlx-todo/server/model"
)

func CreateTodo(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	repository := database.NewTodoRepository(cc.DB)
	todo, err := model.NewTodo(c)
	if err != nil {
		return err
	}
	todo, err = repository.Create(todo)
	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	repository := database.NewTodoRepository(cc.DB)
	todo, err := model.NewTodo(c)
	if err != nil {
		return err
	}
	todo, err = repository.Update(todo)
	return c.JSON(http.StatusCreated, todo)
}

func ListTodo(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	repository := database.NewTodoRepository(cc.DB)
	todo, err := repository.Read()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c echo.Context) error {
	cc := c.(*middleware.CustomContext)
	repository := database.NewTodoRepository(cc.DB)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = repository.Delete(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
