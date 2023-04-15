package common

import (
	"os"
	"regexp"
)

func RootPath() string {
	projectDirName := APP_NAME
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}
