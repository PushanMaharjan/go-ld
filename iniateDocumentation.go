package autodocumentation

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitiateDocumentation(routes gin.RoutesInfo) {
	for _, routeInfo := range routes {
		fmt.Printf("Method: %s, Path: %s, Handler: %s\n", routeInfo.Method, routeInfo.Path, routeInfo.Handler)

		text, err := extractMethodName(routeInfo.Handler)
		if err != nil {
			fmt.Println("Error extracting method name:", err)
			continue
		}
		walkThroughFiles(text)

	}

}

func searchInFile(fileName string, searchText string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	// Check if content contains searchText
	if strings.Contains(string(content), searchText) {
		fmt.Println("Found in:", fileName)
	}
	return nil
}

func walkThroughFiles(searchText string) {
	root := "./" // directory to start searching from

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

func extractMethodName(input string) (string, error) {
	// Compile the regular expression to extract the method name
	re, err := regexp.Compile(`\.\(([^)]+)\)\.([^-\s]+)`)
	if err != nil {
		return "", err
	}

	// Find the match
	matches := re.FindStringSubmatch(input)
	if matches != nil && len(matches) > 2 {
		return matches[2], nil // Return the method name
	}

	return "", fmt.Errorf("no method name found")
}
