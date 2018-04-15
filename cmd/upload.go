package cmd

import (
	"path"

	"fmt"

	"os"

	"github.com/khoiracle/dido/pkg/carthage"
	"github.com/khoiracle/dido/pkg/homepath"
	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
)

var defaultCacheFolderPath, _ = homepath.Expand("~/Library/Caches/dido")

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Store the frameworks",
	Run: func(cmd *cobra.Command, args []string) {
		versionFiles, err := carthage.GetVersionFiles(carthageBuildPath)

		if err != nil {
			exit(err)
		}

		for _, vf := range versionFiles {
			repoSavePath := path.Join(defaultCacheFolderPath, vf.RepoName)
			commitishSavePath := path.Join(repoSavePath, vf.Version.Commitish)

			for _, v := range vf.Version.IOS {
				fmt.Fprintf(os.Stdout, "💾  Saving %s\n", v.Name)
				frameworkSavePath := path.Join(commitishSavePath, carthage.PLATFORM_iOS)
				zipOutput := path.Join(frameworkSavePath, v.Name+".zip")

				if err := os.MkdirAll(frameworkSavePath, os.ModePerm); err != nil {
					exit(err)
				}

				frameworkOnDisk, _ := carthage.FrameworkExist(carthageBuildPath, v.Name, carthage.PLATFORM_iOS)
				dsymOnDisk, _ := carthage.DsymExist(carthageBuildPath, v.Name, carthage.PLATFORM_iOS)

				if err := archiver.Zip.Make(zipOutput, []string{frameworkOnDisk, dsymOnDisk}); err != nil {
					exit(err)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
