package cmd

import (
	"path/filepath"
	"time"
)

var storageFilePath = filepath.Join("storage", "storage.csv")
var maxIdFilePath = filepath.Join("storage", "max-id.txt")
var defaultTimeFormat = time.RFC3339
