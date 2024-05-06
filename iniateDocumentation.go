package autodocumentation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitiateDocumentation(routes gin.RoutesInfo) {
	for _, routeInfo := range routes {
		fmt.Printf("Method: %s, Path: %s, Handler: %s\n", routeInfo.Method, routeInfo.Path, routeInfo.Handler)

	}

	walkThroughFiles()
}

func searchInFile(fileName string, searchText string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err // handle the error, maybe continue with next file
	}

	// Check if content contains searchText
	if strings.Contains(string(content), searchText) {
		fmt.Println("Found in:", fileName)
	}
	return nil
}

func walkThroughFiles() {
	root := "./"                            // directory to start searching from
	searchText := "GetChatRoomListForAdmin" // text to search for

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // prevent panic by handling failure accessing a path
		}
		if !info.IsDir() {
			err := searchInFile(path, searchText)
			if err != nil {
				fmt.Println("Error reading file:", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, err)
	}
}
