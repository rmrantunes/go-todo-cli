package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var storageFilePath = filepath.Join("storage", "storage.csv")
var maxIdFilePath = filepath.Join("storage", "max-id.txt")
var defaultTimeFormat = time.RFC3339
var csvHeader = "ID,Task,Done,CreatedAt\n"

func confirmInputAction(promt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(promt + " (yes/no): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	return input == "yes"
}
