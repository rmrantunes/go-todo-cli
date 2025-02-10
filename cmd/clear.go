/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
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

		csvFile, err := util.LoadFile(storageFilePath)

		util.DieOnError(err)

		defer util.CloseFile(csvFile)

		maxIdFile, err := util.LoadFile(maxIdFilePath)

		util.DieOnError(err)

		defer util.CloseFile(maxIdFile)

		os.WriteFile(storageFilePath, []byte(csvHeader), os.ModeAppend)
		os.WriteFile(maxIdFilePath, []byte(""), os.ModeAppend)
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
