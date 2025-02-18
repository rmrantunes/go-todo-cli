/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"todo-cli/internal/database"
	"todo-cli/internal/database/repository"
	"todo-cli/util"

	"github.com/spf13/cobra"
)

var id int

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.New()
		defer db.Close()

		todoRepository := repository.NewTodoRepository(&repository.TodoRepositoryInject{
			DB: db.DB,
		})

		ctx := context.Background()

		existingTodo, err := todoRepository.FindOneById(ctx, id)

		util.DieOnError(err)

		if existingTodo == nil {
			fmt.Println("Todo not found with id", id)
			return
		}

		_, err = todoRepository.UpdateOneById(ctx, id, map[string]string{
			"done": "true",
		})

		util.DieOnError(err)

		log.Println("Todo update with id", id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().IntVarP(&id, "id", "i", 0, "Todo ID")
}
