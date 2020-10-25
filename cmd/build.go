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
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
		generator.Build()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	bindOptions(buildCmd, "output", "o", "./output",
		"Optional name of the output path to be generated.")

	bindOptions(buildCmd, "host-url", "", "https://example.com",
		"itms-services host address in server.")

	bindOptions(buildCmd, "title", "t", "Test app",
		"title shows in install alert dialog and all pages.")

	bindOptions(buildCmd, "bundle-identifier", "", "com.example.sample",
		"iOS app bundle identifier, same as bundleID in ipa.")

	bindOptions(buildCmd, "bundle-version", "", "1.0",
		"iOS app bundle version, same as CFBundleVersion in ipa.")

	bindOptions(buildCmd, "build-number", "", "123456",
		"iOS app context of the build number.")

	// options for manifest.plist

	bindOptions(buildCmd, "manifest-name", "", "manifest.plist",
		"Optional name of the output manifest file name to be generated.")

	bindOptions(buildCmd, "manifest-software-package", "",
		"https://example.com/install/sample.ipa",
		"manifest.plist template context of software-package.")

	bindOptions(buildCmd, "manifest-display-image", "",
		"https://example.com/assets/display-image.png",
		"manifest.plist template context of display-image.")

	bindOptions(buildCmd, "manifest-full-size-image", "",
		"https://example.com/assets/full-size-image.png",
		"manifest.plist template context of full-size-image.")

	bindOptions(buildCmd, "manifest-subtitle", "", "a display subtitle",
		"manifest.plist template context of subtitle.")

	// options for index.html

	bindOptions(buildCmd, "index-name", "", "index.html",
		"Optional name of the output html file name to be generated.")

	bindOptions(buildCmd, "index-platform", "", "iOS",
		"index.html template context of platform.")

	bindOptions(buildCmd, "index-branch", "", "trunk",
		"index.html template context of build code branch.")

	bindOptions(buildCmd, "index-jumbotron-description", "",
		"Scan the QR code to open this page on your iPhone, click the button to install",
		"index.html template context of jumbotron description.")

	bindOptions(buildCmd, "index-install-button-text", "", "Install",
		"index.html template context of install button text.")
}

func bindOptions(cmd *cobra.Command, name, shorthand string, value string, usage string) {
	cmd.PersistentFlags().StringP(name, shorthand, value, usage)
	_ = viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
}
