package src

// Input is the input struct for the command line utility and the configuration for the library.
type Input struct {
	Source        string
	Destination   string
	IncludeHidden bool
	CopyExisting  bool
	DeleteMissing bool
}

// fileType is a custom type for indicating whether a path is a file or a directory
type fileType string

const (
	file fileType = "file"
	dir  fileType = "dir"
)

// pathInfo is the struct that contains the required information about a path
type pathInfo struct {
	fullPath         string   // fullPath is the full path of the file/directory from the root
	path             string   // path is the partial path of the file/directory from the source/destination passed in
	fileType         fileType // fileType is the type of the path, either the file or directory
	hidden           bool     // hidden is true if the file/directory at the path is hidden
	performOperation bool     // stores whether the operation should be performed on the path, copy or delete
}
