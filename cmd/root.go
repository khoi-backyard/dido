package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dido",
	Short: "Carthage Caching",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exit(err)
	}
}

func init() {
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "ERR: %s", err)
	os.Exit(1)
}
