package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/petprojects9964409/todo_app/internal/models"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, list models.TodoList) (int, error) {
	const op = "Repository.Create"
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	var id int
	createListQuerry := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuerry, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListsTable)
	if _, err := tx.Exec(createUsersListQuery, userID, id); err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userID int) ([]models.TodoList, error) {
	const op = "Repository.GetAll"
	var list []models.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, userListsTable)
	err := r.db.Select(&list, query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return list, nil
}

func (r *TodoListPostgres) GetByID(userID, listID int) (models.TodoList, error) {
	var list models.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, userListsTable)
	err := r.db.Get(&list, query, userID, listID)

	return list, err
}

func (r *TodoListPostgres) Delete(userID, listID int) error {
	const op = "Repository.GetByID"
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, userListsTable)
	_, err := r.db.Exec(query, userID, listID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *TodoListPostgres) Update(userID int, listID int, input models.UpdateListInput) error {
	const op = "Repository.Update"
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, setQuery, userListsTable, argID, argID+1)
	args = append(args, listID, userID)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
