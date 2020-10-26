/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xykong/itms-services/generator"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build [ipa]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")

		if len(args) > 0 {
			viper.Set("ipa", args[0])
		}

		generator.Build()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	bindOptions(buildCmd, false, "output", "o", "./output",
		"Optional name of the output path to be generated.")

	bindOptions(buildCmd, true, "host-url", "", "",
		"itms-services host address in server. example: https://example.com")

	bindOptions(buildCmd, false, "title", "t", "Test app",
		"title shows in install alert dialog and all pages.")

	bindOptions(buildCmd, false, "bundle-identifier", "", "",
		"iOS app bundle identifier, same as CFBundleIdentifier in ipa.")

	bindOptions(buildCmd, false, "bundle-version", "", "1.0",
		"iOS app bundle version, same as CFBundleVersion in ipa.")

	bindOptions(buildCmd, false, "build-number", "", "0",
		"iOS app context of the build number.")

	// options for manifest.plist

	bindOptions(buildCmd, false, "manifest-name", "", "manifest.plist",
		"Optional name of the output manifest file name to be generated.")

	bindOptions(buildCmd, true, "manifest-software-package", "", "",
		"manifest.plist template context of software-package. example: https://example.com/install/sample.ipa")

	bindOptions(buildCmd, true, "display-image", "", "",
		"manifest.plist and index.html template context of display-image. example: https://example.com/assets/display-image.png")

	bindOptions(buildCmd, false, "manifest-full-size-image", "", "",
		"manifest.plist template context of full-size-image. example: https://example.com/assets/full-size-image.png")

	bindOptions(buildCmd, false, "manifest-subtitle", "", "",
		"manifest.plist template context of subtitle.")

	// options for index.html

	bindOptions(buildCmd, false, "index-name", "", "index.html",
		"Optional name of the output html file name to be generated.")

	bindOptions(buildCmd, false, "index-platform", "", "iOS",
		"index.html template context of platform.")

	bindOptions(buildCmd, false, "index-branch", "", "trunk",
		"index.html template context of build code branch.")

	bindOptions(buildCmd, false, "index-jumbotron-description", "",
		"Scan the QR code to open this page on your iPhone, click the button to install",
		"index.html template context of jumbotron description.")

	bindOptions(buildCmd, false, "index-install-button-text", "", "Install",
		"index.html template context of install button text.")
}

func bindOptions(cmd *cobra.Command, required bool, name, shorthand string, value string, usage string) {
	cmd.PersistentFlags().StringP(name, shorthand, value, usage)
	_ = viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))

	if required {
		_ = cmd.MarkPersistentFlagRequired(name)
	}
}
