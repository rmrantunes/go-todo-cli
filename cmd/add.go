/*
Copyright © 2025 RAFAEL ANTUNES rmrantunes.dev@gmail.com
*/
package cmd

import (
	"context"
	"log"
	"todo-cli/internal/database"
	"todo-cli/internal/database/repository"
	"todo-cli/util"

	"github.com/spf13/cobra"
)

var title string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.New()
		defer db.Close()

		todoRepository := repository.NewTodoRepository(&repository.TodoRepositoryInject{
			DB: db.DB,
		})

		description := cmd.Flag("description").Value.String()

		ctx := context.Background()

		todoId, err := todoRepository.CreateTodo(ctx, map[string]interface{}{
			"description": description,
		})

		util.DieOnError(err)

		log.Println("Todo created with id", todoId)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&title, "description", "d", "", "Todo description")
}
