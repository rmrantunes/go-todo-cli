/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"todo-cli/util"

	"github.com/spf13/cobra"
)

var id int

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := util.LoadFile(storageFilePath)

		util.DieOnError(err)

		defer util.CloseFile(file)

		csvReader := csv.NewReader(file)

		csvFile, err := csvReader.ReadAll()

		util.DieOnError(err)

		found := false

		for _, row := range csvFile {
			if row[0] == "ID" {
				continue
			}

			rowId, err := strconv.ParseInt(row[0], 10, 64)

			util.DieOnError(err)

			if int64(id) == rowId {
				row[2] = "true"
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Todo not found with id", id)
			return
		}

		emptyFile, err := os.Create(storageFilePath)

		util.DieOnError(err)

		defer emptyFile.Close()

		csvWriter := csv.NewWriter(emptyFile)
		csvWriter.WriteAll(csvFile)

		err = csvWriter.Error()

		util.DieOnError(err)

		csvWriter.Flush()

		err = csvWriter.Error()

		util.DieOnError(err)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
	completeCmd.Flags().IntVarP(&id, "id", "i", 0, "Todo ID")
}
