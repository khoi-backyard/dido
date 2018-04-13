package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const carthageBuildPath = "./Carthage/Build/"
const iOSPath = carthageBuildPath + "/iOS"
const macOSPath = carthageBuildPath + "/Mac"
const tvOSPath = carthageBuildPath + "/tvOS"
const watchOSPath = carthageBuildPath + "/watchOS"

var rootCmd = &cobra.Command{
	Use:   "dido",
	Short: "Carthage Caching",
}

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
