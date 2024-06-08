package cmd

import (
	"fmt"
	"os"

	"github.com/droppedasbaby/dir-mirror/src"
	"github.com/spf13/cobra"
)

// in is a global variable that holds the input flags
var in src.Input

// rootCmd is the main command
// The command is used to mirror a directory
// It is the only command in this application
var rootCmd = &cobra.Command{
	Use:   "dir-mirror",
	Short: "Mirrors a directory to another location.",
	Long: `Mirrors a directory to another location.
	Has the ability to mirror nested directories and files.
	You can choose to mirror only files or directories only.
	Allows you to filter directories and files that already exist.
	Deletes files and directories that no longer exist at source if required.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := validateInput(in)
		if err != nil {
			cmd.PrintErrln(fmt.Errorf("invalid input, cannot proceed: %v", err))
			os.Exit(1)
		}

		err = src.Mirror(&in)
		if err != nil {
			cmd.PrintErrln(fmt.Errorf("could not mirror directory: %v", err))
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&in.Source, "source", "", "", "Directory to mirror, does not include the directory itself")
	rootCmd.Flags().StringVarP(&in.Destination, "destination", "", "", "Destination directory, the files and directories will be mirrored inside this directory")
	rootCmd.Flags().BoolVarP(&in.IncludeHidden, "include-hidden", "", false, "Include hidden files and directories while mirroring")
	rootCmd.Flags().BoolVarP(&in.CopyExisting, "copy-existing", "", false, "Copy existing files and directories, overwrite if already exists")
	rootCmd.Flags().BoolVarP(&in.DeleteMissing, "delete-missing", "", false, "Delete files and directories that are missing in the source directory")
}
