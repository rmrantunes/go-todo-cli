/*
Copyright © 2025 RAFAEL ANTUNES rmrantunes.dev@gmail.com
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

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Wipes all todos",
	Long:  `Be careful while using it! It wipes all todos!`,
	Run: func(cmd *cobra.Command, args []string) {
		if !confirmInputAction("Are you sure you want to remove all todos?") {
			fmt.Println("Action cancelled. No todos removed.")
			return
		}

		db := database.New()
		defer db.Close()

		todoRepository := repository.NewTodoRepository(&repository.TodoRepositoryInject{
			DB: db.DB,
		})

		ctx := context.Background()

		err := todoRepository.DangerouslyDeleteAll(ctx)

		log.Println("All todos deleted com storage")

		util.DieOnError(err)
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
