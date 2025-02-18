/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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

var deleteId int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.New()
		defer db.Close()

		todoRepository := repository.NewTodoRepository(&repository.TodoRepositoryInject{
			DB: db.DB,
		})

		ctx := context.Background()

		existingTodo, err := todoRepository.FindOneById(ctx, deleteId)

		util.DieOnError(err)

		if existingTodo == nil {
			log.Println("Todo not found with id", deleteId)
			return
		}

		_, err = todoRepository.DeleteOneById(ctx, deleteId)

		util.DieOnError(err)

		log.Println("Todo deleted with id", deleteId)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVarP(&deleteId, "id", "i", 0, "Todo ID")
}
