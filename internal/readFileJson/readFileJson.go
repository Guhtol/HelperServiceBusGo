package readFileJson

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFileJsonToGetMessageBody(queueName, dirName string) (message string, success bool) {
	jsonPath := filepath.Join(dirName, "*.json")
	jsonFiles, err := filepath.Glob(jsonPath)
	if err != nil {
		panic(err)
	}

	for _, file := range jsonFiles {
		ext := filepath.Ext(file)
		fileBase := filepath.Base(file)

		nameFile := strings.TrimSuffix(fileBase, ext)
		if nameFile != queueName {
			continue
		}

		messageBytes, err := os.ReadFile(file)
		if err != nil {
			break
		}
		message = string(messageBytes)
		success = true
	}
	return
}
