/*
Copyright © 2025 RAFAEL ANTUNES rmrantunes.dev@gmail.com
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"todo-cli/internal/database"
	"todo-cli/internal/database/repository"
	"todo-cli/util"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var all bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long: `By default it list the uncompleted todos only. 

If you want to list all, use the flag -a or --all
	`,
	Run: func(cmd *cobra.Command, args []string) {
		db := database.New()
		defer db.Close()

		todoRepository := repository.NewTodoRepository(&repository.TodoRepositoryInject{
			DB: db.DB,
		})

		ctx := context.Background()

		equals := map[string]string{}

		if !all {
			equals = map[string]string{
				"done": "false",
			}
		}

		todos, err := todoRepository.List(ctx, equals)

		util.DieOnError(err)

		tbWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		lineTextFormat := "%s\t%s\t%s\n"
		headerData := []any{"ID", "Description", "CreatedAt"}

		if all {
			lineTextFormat = "%s\t%s\t%s\t%s\n"
			headerData = []any{"ID", "Description", "Done", "CreatedAt"}
		}

		_, err = fmt.Fprintf(tbWriter, lineTextFormat, headerData...)

		util.DieOnError(err)

		for _, todo := range todos {

			createdAtColumn := timediff.TimeDiff(todo.CreatedAt)

			lineData := []any{fmt.Sprintf("%d", todo.Id), todo.Description, createdAtColumn}

			if all {
				lineData = []any{fmt.Sprintf("%d", todo.Id), todo.Description, fmt.Sprintf("%v", todo.Done), createdAtColumn}
			}

			_, err = fmt.Fprintf(tbWriter, lineTextFormat, lineData...)

			util.DieOnError(err)
		}

		err = tbWriter.Flush()
		util.DieOnError(err)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "List all todos")
}
