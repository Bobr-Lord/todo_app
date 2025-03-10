package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list models.TodoList) (int, error) {
	const op = "Repository.Create"
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	var id int
	createListQuerry := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuerry, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListsTable)
	if _, err := tx.Exec(createUsersListQuery, userId, id); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]models.TodoList, error) {
	const op = "Repository.GetAll"
	var list []models.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, userListsTable)
	err := r.db.Select(&list, query, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return list, nil
}

func (r *TodoListPostgres) GetById(userId, listId int) (models.TodoList, error) {
	const op = "Repository.GetById"
	var list models.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, userListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
