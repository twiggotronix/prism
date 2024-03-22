package utils

import (
	"regexp"
	"strings"
)

type PathUtils interface {
	UrlCorrespondsToPath(url string) bool
}

type pathUtils struct {
	path string
}

func NewPathUtils(path string) PathUtils {
	return &pathUtils{
		path,
	}
}

func replaceAfterPosition(input, replaceStr string, position int, replacement string) string {
	// Find the index of the first occurrence of the string after the specified position
	index := strings.Index(input[position:], replaceStr)
	if index == -1 {
		// String not found after the specified position
		return input
	}
	index += position // Adjust index to account for position offset

	// Replace the first occurrence of the string after the specified position
	replaced := input[:index] + replacement + input[index+len(replaceStr):]
	return replaced
}

func extractVariableValue(input string, fromIndex int) string {
	subUrl := input[fromIndex:]
	nextSlashIndex := strings.Index(subUrl, "/")
	var variableValue string
	if nextSlashIndex >= 0 {
		variableValue = subUrl[:nextSlashIndex]
	} else {
		variableValue = subUrl
	}
	return variableValue
}

func (p *pathUtils) UrlCorrespondsToPath(url string) bool {
	path := p.path
	if path == url {
		return true
	}
	re := regexp.MustCompile(`:[a-zA-Z_][a-zA-Z0-9_]*(?:/|$)`)
	newPath := url
	variables := re.FindAllString(path, -1)
	for _, variable := range variables {
		cleanVariableName := strings.TrimRight(strings.TrimSpace(variable), "/")
		pathIndex := strings.Index(path, cleanVariableName)
		if pathIndex >= 0 {
			if len(newPath) > pathIndex {
				variableValue := extractVariableValue(newPath, pathIndex)
				newPath = replaceAfterPosition(newPath, variableValue, pathIndex, cleanVariableName)
			} else {
				return false
			}
		}
	}

	return path == newPath
}
