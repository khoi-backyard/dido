package cmd

import (
	"io/ioutil"

	"strings"

	"path"

	"fmt"

	"github.com/khoiracle/dido/pkg"
	"github.com/spf13/cobra"
)

const carthageBuildPath = "./Carthage/Build/"
const iOSPath = carthageBuildPath + "/iOS"
const macOSPath = carthageBuildPath + "/Mac"
const tvOSPath = carthageBuildPath + "/tvOS"
const watchOSPath = carthageBuildPath + "/watchOS"

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all built frameworks in /Carthage",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(carthageBuildPath)

		if err != nil {
			exit(err)
		}

		var frameworks []*pkg.Framework

		for _, f := range files {
			if !strings.HasSuffix(f.Name(), pkg.VersionFileExtension) {
				continue
			}

			versionFilePath := path.Join(carthageBuildPath, f.Name())
			framework, err := pkg.NewFramework(versionFilePath)

			if err != nil {
				exit(err)
			}

			frameworks = append(frameworks, framework)
		}

		for _, v := range frameworks {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
