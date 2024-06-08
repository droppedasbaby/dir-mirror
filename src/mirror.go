package src

import "fmt"

// Mirror mirrors the source directory to the destination directory
// It is the only extra function and therefore the entry point from the command
// to the library.
// It returns an error if it fails
func Mirror(in *Input) error {
	dests, err := discover(in.Destination, false, shouldInclude(in))
	if err != nil {
		return fmt.Errorf("could not discover destination directory: %v", err)
	}
	sources, err := discover(in.Source, true, shouldInclude(in))
	if err != nil {
		return fmt.Errorf("could not discover source directory: %v", err)
	}

	forDelete := reverse(filterMissing(in, sources, dests))
	forCopy := filterExisting(in, sources, dests)

	err = deleteFiles(forDelete)
	if err != nil {
		return fmt.Errorf("could not delete missing files: %v", err)
	}
	err = copyFiles(in, forCopy, in.Destination)
	if err != nil {
		return fmt.Errorf("could not copy files: %v", err)
	}

	return nil
}
