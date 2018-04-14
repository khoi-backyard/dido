package cmd

import (
	"os"

	"github.com/khoiracle/dido/pkg"
	"github.com/spf13/cobra"
)

const defaultCacheFolderPath = "~/Library/Caches/dido"

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Store the frameworks",
	Run: func(cmd *cobra.Command, args []string) {
		fullPath, err := pkg.Expand(defaultCacheFolderPath)

		if err != nil {
			exit(err)
		}

		if err = os.MkdirAll(fullPath, os.ModePerm); err != nil {
			exit(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
