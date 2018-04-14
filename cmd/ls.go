package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/khoiracle/dido/pkg/carthage"
	"github.com/spf13/cobra"
)

var platform string

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all frameworks specified in .version files",
	Run: func(cmd *cobra.Command, args []string) {
		if platform != "" && !carthage.IsValidPlatform(platform) {
			exit(fmt.Errorf("%s is not a valid platform", platform))
		}

		versionFiles, err := carthage.GetVersionFiles(carthageBuildPath)

		if err != nil {
			exit(err)
		}

		for _, vf := range versionFiles {
			fmt.Fprintf(os.Stdout, "📦  %s:\n", vf.RepoName)

			if platform == "" || platform == carthage.PLATFORM_iOS {
				outputHashes(os.Stdout, carthage.PLATFORM_iOS, vf.Hashes(carthage.PLATFORM_iOS))
			}

			if platform == "" || platform == carthage.PLATFORM_watchOS {
				outputHashes(os.Stdout, carthage.PLATFORM_watchOS, vf.Hashes(carthage.PLATFORM_watchOS))
			}

			if platform == "" || platform == carthage.PLATFORM_macOS {
				outputHashes(os.Stdout, carthage.PLATFORM_macOS, vf.Hashes(carthage.PLATFORM_macOS))
			}

			if platform == "" || platform == carthage.PLATFORM_tvOS {
				outputHashes(os.Stdout, carthage.PLATFORM_tvOS, vf.Hashes(carthage.PLATFORM_tvOS))
			}
		}
	},
}

func outputHashes(w io.Writer, platform string, hashes []carthage.VersionHash) {
	if len(hashes) > 0 {
		fmt.Fprintf(w, "%s\n", platform)
	}
	for _, h := range hashes {
		fmt.Fprintf(w, "  %s %s\n", h.Name, h.Hash)
	}
}

func init() {
	lsCmd.Flags().StringVar(&platform, "platform", "", "specify platform ios|macos|tvos|watchos|all")
	rootCmd.AddCommand(lsCmd)
}
