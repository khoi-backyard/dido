package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

// Set by ldflags for production build
var version = "development"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of dido",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "%s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
