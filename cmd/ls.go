package cmd

import (
	"io/ioutil"

	"strings"

	"path"

	"fmt"

	"os"

	"io"

	"github.com/khoiracle/dido/pkg"
	"github.com/spf13/cobra"
)

var platform string

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all frameworks specified in .version files",
	Run: func(cmd *cobra.Command, args []string) {
		if platform != "" && !pkg.IsValidPlatform(platform) {
			exit(fmt.Errorf("%s is not a valid platform", platform))
		}

		files, err := ioutil.ReadDir(carthageBuildPath)

		if err != nil {
			exit(err)
		}

		for _, f := range files {
			if !strings.HasSuffix(f.Name(), pkg.VersionFileExtension) {
				continue
			}

			versionFilePath := path.Join(carthageBuildPath, f.Name())
			versionFile, err := pkg.NewVersionFile(versionFilePath)

			if err != nil {
				exit(err)
			}

			fmt.Fprintf(os.Stdout, "📦 %s:\n", versionFile.RepoName)

			if platform == "" || platform == pkg.PLATFORM_iOS {
				outputHashes(os.Stdout, pkg.PLATFORM_iOS, versionFile.Hashes(pkg.PLATFORM_iOS))
			}

			if platform == "" || platform == pkg.PLATFORM_watchOS {
				outputHashes(os.Stdout, pkg.PLATFORM_watchOS, versionFile.Hashes(pkg.PLATFORM_watchOS))
			}

			if platform == "" || platform == pkg.PLATFORM_macOS {
				outputHashes(os.Stdout, pkg.PLATFORM_macOS, versionFile.Hashes(pkg.PLATFORM_macOS))
			}

			if platform == "" || platform == pkg.PLATFORM_tvOS {
				outputHashes(os.Stdout, pkg.PLATFORM_tvOS, versionFile.Hashes(pkg.PLATFORM_tvOS))
			}
		}
	},
}

func outputHashes(w io.Writer, platform string, hashes []pkg.VersionHash) {
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
