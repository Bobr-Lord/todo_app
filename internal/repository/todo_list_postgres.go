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

func (r *TodoListPostgres) Delete(userId, listId int) error {
	const op = "Repository.GetById"
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, userListsTable)
	_, err := r.db.Exec(query, userId, listId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *TodoListPostgres) Update(userId int, listId int, input models.UpdateListInput) error {
	const op = "Repository.Update"
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, setQuery, userListsTable, argId, argId+1)
	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
