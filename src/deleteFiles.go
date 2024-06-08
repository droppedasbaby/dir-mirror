package src

import (
	"fmt"
	"os"
)

// deleteFiles deletes files/directories from the destination directory that are no longer
// present in the source directory
func deleteFiles(files []pathInfo) error {
	for _, file := range files {
		if file.performOperation {
			err := os.RemoveAll(file.fullPath)
			if err != nil {
				return fmt.Errorf("could not delete %v: %v", file.fullPath, err)
			}
			fmt.Printf("Deleted %s\n", file.fullPath)
		}
	}
	return nil
}
