package database

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/saiki/golang-echo-sqlx-todo/server/model"
)

type TodoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (repository *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	tx := repository.db.MustBegin()

	result, err := tx.NamedExec("insert into todo (task, finished) values (:task, :finished)", todo)
	if err != nil {
		log.Fatal(tx.Rollback())
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(tx.Rollback())
		return nil, err
	}
	todo.Id = int(id)

	err = tx.Commit()
	if err != nil {
		log.Fatal(tx.Rollback())
		return nil, err
	}
	return todo, nil
}

func (repository *TodoRepository) Read() ([]model.Todo, error) {
	todo := []model.Todo{}
	err := repository.db.Select(&todo, "select id, task, finished from todo")
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (repository *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	tx := repository.db.MustBegin()

	result, err := tx.NamedExec("update todo set task = :task, finished = :finished where id = :id", todo)
	if err != nil {
		log.Fatal(tx.Rollback())
		return nil, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(tx.Rollback())
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("not affected.")
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(tx.Rollback())
		return nil, err
	}
	return todo, nil
}

func (repository *TodoRepository) Delete(id int) error {
	tx := repository.db.MustBegin()

	result, err := tx.NamedExec("delete from todo where id = ?", id)
	if err != nil {
		log.Fatal(tx.Rollback())
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(tx.Rollback())
		return err
	}
	if affected == 0 {
		return errors.New("not affected.")
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(tx.Rollback())
		return err
	}
	return nil
}
