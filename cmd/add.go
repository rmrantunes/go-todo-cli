/*
Copyright © 2025 RAFAEL ANTUNES rmrantunes.dev@gmail.com
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"path/filepath"
	"time"
	"todo-cli/util"

	"github.com/spf13/cobra"
)

var title string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := filepath.Join("storage", "storage.csv")
		file, err := util.LoadFile(filePath)

		util.DieOnError(err)

		defer util.CloseFile(file)

		csvReader := csv.NewReader(file)

		csvFile, err := csvReader.ReadAll()

		util.DieOnError(err)

		fmt.Println(csvFile)

		csvWriter := csv.NewWriter(file)

		util.DieOnError(err)

		err = csvWriter.Write([]string{
			"1",
			cmd.Flag("description").Value.String(),
			"false",
			time.Now().Format(time.RFC3339),
		})

		util.DieOnError(err)

		csvWriter.Flush()

		err = csvWriter.Error()

		util.DieOnError(err)

		fmt.Println("Todo created with id", "1")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&title, "description", "t", "", "Todo description")
}
