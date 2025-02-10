/*
Copyright © 2025 RAFAEL ANTUNES rmrantunes.dev@gmail.com
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"todo-cli/util"

	"github.com/spf13/cobra"
)

var title string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := util.LoadFile(storageFilePath)

		util.DieOnError(err)

		defer util.CloseFile(file)

		csvReader := csv.NewReader(file)

		csvFile, err := csvReader.ReadAll()

		util.DieOnError(err)

		fmt.Println(csvFile)

		csvWriter := csv.NewWriter(file)

		util.DieOnError(err)

		maxIdFile, err := os.ReadFile(maxIdFilePath)

		util.DieOnError(err)

		maxIdString := strings.ReplaceAll(string(maxIdFile), "\n", "")

		if maxIdString == "" {
			maxIdString = "0"
		}

		maxId, err := strconv.ParseInt(maxIdString, 10, 64)

		util.DieOnError(err)

		todoId := maxId + 1

		err = csvWriter.Write([]string{
			strconv.FormatInt(todoId, 10),
			cmd.Flag("description").Value.String(),
			"false",
			time.Now().Format(time.RFC3339),
		})

		util.DieOnError(err)

		csvWriter.Flush()

		err = csvWriter.Error()

		util.DieOnError(err)

		newMaxIdFileBytes := []byte(strconv.FormatInt(todoId, 10))

		err = os.WriteFile(maxIdFilePath, newMaxIdFileBytes, os.ModeAppend)

		util.DieOnError(err)

		fmt.Println("Todo created with id", todoId)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&title, "description", "d", "", "Todo description")
}
