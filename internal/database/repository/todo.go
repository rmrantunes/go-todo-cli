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

func (r *TodoRepository) FindOneById(ctx context.Context, id int) (*model.Todo, error) {
	todo := &model.Todo{}

	query := "SELECT id, description, done, created_at FROM todo WHERE id = ?"

	err := r.db.QueryRowContext(ctx, query, id).Scan(&todo.Id, &todo.Description, &todo.Done, &todo.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) UpdateOneById(ctx context.Context, id int, update map[string]string) (*model.Todo, error) {
	todo := &model.Todo{}

	query := "UPDATE todo"

	if len(update) != 0 {
		query = fmt.Sprintf("%s %s", query, "SET")
	}

	i := 0
	for key, item := range update {
		query = fmt.Sprintf("%s %s = %s", query, key, item)

		if i < len(update)-1 {
			query = fmt.Sprintf("%s%s", query, ",")
		}
		i++
	}

	query = fmt.Sprintf("%s %s", query, "RETURNING id, description, done, created_at")

	err := r.db.QueryRowContext(ctx, query).Scan(&todo.Id, &todo.Description, &todo.Done, &todo.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) DeleteOneById(ctx context.Context, id int) (*model.Todo, error) {
	todo := &model.Todo{}

	query := "DELETE FROM todo WHERE id = ? RETURNING id, description, done, created_at"

	err := r.db.QueryRowContext(ctx, query, id).Scan(&todo.Id, &todo.Description, &todo.Done, &todo.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return todo, nil
}

func (r *TodoRepository) DangerouslyDeleteAll(ctx context.Context) error {
	query := "DELETE FROM todo"

	_, err := r.db.ExecContext(ctx, query)

	if err != nil {
		return err
	}

	return nil
}
