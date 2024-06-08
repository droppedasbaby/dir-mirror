package cmd

import (
	"fmt"

	"github.com/droppedasbaby/dir-mirror/src"
)

// validateInput validates the input flags.
// It makes sure the source and destination are provided.
// Source and destination cannot be the same.
func validateInput(in src.Input) error {
	if in.Source == "" {
		return fmt.Errorf("source directory is required")
	}
	if in.Destination == "" {
		return fmt.Errorf("destination directory is required")
	}

	if in.Source == in.Destination {
		return fmt.Errorf("source and destination directories cannot be the same")
	}

	return nil
}
