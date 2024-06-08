package src

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// convertToPathInfo converts the path to a pathInfo struct with the required information
func convertToPathInfo(source string, path string, info os.FileInfo, performOperation bool) pathInfo {
	fullPath := path
	cleanedPath := strings.Replace(path, source+"/", "", 1)
	cleanedPath = strings.Replace(cleanedPath, source, "", 1)
	fileType := file
	if info.IsDir() {
		fileType = dir
	}
	hidden := info.Name()[0] == '.'

	return pathInfo{
		fullPath:         fullPath,
		path:             cleanedPath,
		fileType:         fileType,
		hidden:           hidden,
		performOperation: performOperation,
	}
}

// discover walks the directory and returns the paths
// It filters out the paths based on the shouldInclude function
// It returns an error if it fails
func discover(
	rootPath string,
	performOperation bool,
	shouldInclude shouldIncludeFilter,
) ([]pathInfo, error) {
	paths := []pathInfo{}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("could not walk %v: %w", path, err)
		}
		include := shouldInclude(info)

		if rootPath == path {
			return nil
		}
		pathInfo := convertToPathInfo(rootPath, path, info, performOperation)

		if include {
			paths = append(paths, pathInfo)
		} else if !include && info.IsDir() {
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not walk %v: %v", rootPath, err)
	}
	return paths, nil
}
