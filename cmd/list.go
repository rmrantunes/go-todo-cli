/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
	"todo-cli/util"

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
		all := cmd.Flag("all").Value.String() == "true"

		file, err := util.LoadFile(storageFilePath)

		util.DieOnError(err)

		defer util.CloseFile(file)

		csvReader := csv.NewReader(file)

		tbWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		for {
			csvLine, err := csvReader.Read()

			if err != nil {
				if err.Error() == "EOF" {
					break
				}

				util.DieOnError(err)
			}

			if !all && csvLine[2] == "true" {
				continue
			}

			lineTextFormat := "%s\t%s\t%s\n"
			lineData := []any{csvLine[0], csvLine[1], csvLine[3]}

			if all {
				lineTextFormat = "%s\t%s\t%s\t%s\n"
				lineData = []any{csvLine[0], csvLine[1], csvLine[2], csvLine[3]}
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
