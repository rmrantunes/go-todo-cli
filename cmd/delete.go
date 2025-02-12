/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strconv"
	"todo-cli/util"

	"github.com/spf13/cobra"
)

var id int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := util.LoadFile(storageFilePath)

		util.DieOnError(err)

		defer util.CloseFile(file)

		csvReader := csv.NewReader(file)

		csvFile, err := csvReader.ReadAll()

		util.DieOnError(err)

		possibleIndex, found := slices.BinarySearchFunc(csvFile, id, func(line []string, id int) int {
			csvRowID, err := strconv.ParseInt(line[0], 10, 64)

			if line[0] == "ID" {
				return -1
			}

			if err != nil {
				fmt.Println(err.Error())
				return -1
			}

			return int(csvRowID) - id
		})

		if !found {
			fmt.Println("Todo not found with id", id)
			return
		}

		csvFile = slices.Delete(csvFile, possibleIndex, possibleIndex+1)

		fmt.Println(csvFile)

		err = os.WriteFile(storageFilePath, []byte(""), fs.ModeAppend)
		util.DieOnError(err)

		csvWriter := csv.NewWriter(file)
		csvWriter.WriteAll(csvFile)

		csvWriter.Flush()

		err = csvWriter.Error()

		util.DieOnError(err)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "Todo ID")
}
