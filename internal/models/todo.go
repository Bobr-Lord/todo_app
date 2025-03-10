package models

import "errors"

type TodoList struct {
	Id          int64  `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int64
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int64
	UserId int
	ListId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("missing required field")
	}
	return nil
}
