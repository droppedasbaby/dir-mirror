package src

import (
	"os"
)

type (
	// shouldIncludeFilter is a type for filtering files based on the input
	shouldIncludeFilter func(os.FileInfo) bool
	// existingFileFilter is a type for filtering existing files based on the input
	existingFileFilter func(os.FileInfo, os.FileInfo) bool
)

// shouldInclude returns a filter function that filters files based on command line input
func shouldInclude(in *Input) shouldIncludeFilter {
	return func(info os.FileInfo) bool {
		if !in.IncludeHidden && info.Name()[0] == '.' {
			return false
		}

		return true
	}
}

// convertToSet converts a list of pathInfo to a set, for faster lookup
func convertToSet(paths []pathInfo) map[string]bool {
	set := make(map[string]bool)
	for _, path := range paths {
		set[path.path] = true
	}
	return set
}

// filterExisting filters the source, it returns the paths that don't exist in the destination based on input
func filterExisting(in *Input, sources []pathInfo, dests []pathInfo) []pathInfo {
	if in.CopyExisting {
		return sources
	}

	destsSet := convertToSet(dests)
	for i, s := range sources {
		if _, ok := (destsSet)[s.path]; ok {
			s.performOperation = false
			(sources)[i] = s
		}
	}
	return sources
}

// filterMissing filters the destination paths to delete missing files based on the input
func filterMissing(in *Input, sources []pathInfo, dests []pathInfo) []pathInfo {
	if !in.DeleteMissing {
		return dests
	}

	sourcesSet := convertToSet(sources)
	for i, d := range dests {
		if _, ok := (sourcesSet)[d.path]; !ok {
			d.performOperation = true
			(dests)[i] = d
		}
	}
	return dests
}

// reverse reverses the elements of the array
func reverse[T any](arr []T) []T {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
