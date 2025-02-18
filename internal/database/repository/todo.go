package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"todo-cli/internal/database/model"
)

type TodoRepositoryInject struct {
	DB *sql.DB
}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(inject *TodoRepositoryInject) *TodoRepository {
	return &TodoRepository{
		db: inject.DB,
	}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, data map[string]interface{}) (int, error) {
	var id int
	timestamp := time.Now().Format(time.RFC3339)

	err := r.db.QueryRowContext(ctx, "INSERT INTO todo(description, created_at) values(?, ?) RETURNING id", data["description"], timestamp).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TodoRepository) List(ctx context.Context, equals map[string]string) ([]model.Todo, error) {
	var todos = make([]model.Todo, 0)

	query := "SELECT id, description, done, created_at FROM todo"

	if len(equals) != 0 {
		query = fmt.Sprintf("%s %s", query, "WHERE")
	}

	i := 0
	for key, item := range equals {
		query = fmt.Sprintf("%s %s = %s", query, key, item)

		if i < len(equals)-1 {
			query = fmt.Sprintf("%s %s", query, "AND")
		}
		i++
	}

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var todo model.Todo

		rows.Scan(&todo.Id, &todo.Description, &todo.Done, &todo.CreatedAt)

		todos = append(todos, todo)
	}

	return todos, nil
}
