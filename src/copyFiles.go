package src

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// copyFile copies a file from source to dest
// It returns an error if it fails
func copyFile(source, dest string) error {
	sFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("could not open source %v: %v", source, err)
	}
	defer sFile.Close()

	dFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("could not create destination %v: %v", dest, err)
	}
	defer dFile.Close()

	_, err = io.Copy(dFile, sFile)
	if err != nil {
		return fmt.Errorf("could not copy %v to %v: %v", source, dest, err)
	}

	return nil
}

// copyFiles copies files/directories from source to dest
// It returns an error if it fails
// It copies only if the file/directory does not exist at destination
// If CopyExisting is true, it copies even if the file/directory exists
func copyFiles(in *Input, sources []pathInfo, dest string) error {
	for _, file := range sources {
		if file.performOperation {
			_, err := os.Stat(file.fullPath)
			if err != nil {
				return fmt.Errorf("could not get %v info: %v", file.fullPath, err)
			}
			dPath := filepath.Join(dest, file.path)
			_, err = os.Stat(dPath)
			if err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("could not get %v info: %v", dPath, err)
			}

			if (in.CopyExisting && file.fileType != dir) || os.IsNotExist(err) {
				if file.fileType == dir {
					err = os.Mkdir(dPath, 0755)
					if err != nil {
						return fmt.Errorf("could not create %v: %v", dPath, err)
					}
				} else {
					err = copyFile(file.fullPath, dPath)
					if err != nil {
						return fmt.Errorf("could not copy %v: %v", dPath, err)
					}
				}
				fmt.Printf("Copied %s to %s\n", file.fullPath, dPath)
			}
		}
	}
	return nil
}
