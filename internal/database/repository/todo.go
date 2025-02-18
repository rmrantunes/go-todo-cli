package repository

import (
	"context"
	"database/sql"
	"time"
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
