package autodocumentation

import (
	"bufio"
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

		rootPath, handlerName, err := extractMethodName(routeInfo.Handler)
		if err != nil {
			fmt.Println("Error extracting method name:", err)
			continue
		}
		walkThroughFiles(rootPath, handlerName)

	}

}

func searchInFile(fileName string, searchText string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	pattern := fmt.Sprintf(`%s\((\w+)\s+\*gin\.Context\)`, searchText)

	r, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return err
	}

	matches := r.FindAllString(string(content), -1)
	if matches != nil {
		fmt.Println("Found in:", fileName)
		findFunctionScope(fileName, r)
	}
	return nil
}

func walkThroughFiles(rootPath, searchText string) {
	root := rootPath // directory to start searching from

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

func extractMethodName(input string) (string, string, error) {
	// Compile the regular expression to extract the method name and path
	re, err := regexp.Compile(`[^/]*?/(.*?)/?(\(\*[^)]+\))?\.?([^.-]+)`)
	if err != nil {
		return "", "", err
	}

	// Find the match
	matches := re.FindStringSubmatch(input)
	if len(matches) > 3 {
		// Return the path and the method name
		return matches[1], matches[3], nil
	}

	return "", "", fmt.Errorf("no valid matches found")
}

func findFunctionScope(filePath string, regexToFind *regexp.Regexp) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inFunction := false
	braceCount := 0

	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()

		if inFunction {
			if strings.Contains(line, "{") {
				braceCount++
			}
			if strings.Contains(line, "}") {
				braceCount--
			}
			if braceCount == 0 {
				fmt.Printf("Function ends at line %d\n", lineNumber)
				return
			}
		} else {

			matches := regexToFind.FindAllString(string(line), -1)
			if matches != nil {
				inFunction = true
				braceCount++
				fmt.Printf("Function starts at line %d\n", lineNumber)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
