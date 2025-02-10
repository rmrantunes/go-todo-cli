/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
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

			createdAtColumn := csvLine[3]

			if createdAtColumn != "CreatedAt" {
				createdAtTime, err := time.Parse(defaultTimeFormat, createdAtColumn)

				if err != nil {
					createdAtColumn = "<<MALFORMED DATE>>"
				} else {
					createdAtColumn = timediff.TimeDiff(createdAtTime)
				}
			}

			lineTextFormat := "%s\t%s\t%s\n"
			lineData := []any{csvLine[0], csvLine[1], createdAtColumn}

			if all {
				lineTextFormat = "%s\t%s\t%s\t%s\n"
				lineData = []any{csvLine[0], csvLine[1], csvLine[2], createdAtColumn}
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
