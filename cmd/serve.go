package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xykong/itms-services/generator"
)

var serveCmd = &cobra.Command{
	Use:   "serve [ipa]",
	Short: "Build and serve the installation page over HTTP",
	Long: `Build generates the manifest.plist and index.html from the given .ipa file,
then immediately starts an HTTP server so you can install the app by visiting
the printed URL on your iOS device (must be HTTPS in production — use a
reverse proxy such as Caddy or ngrok for TLS).

Example:
  # Serve locally on port 8080, reachable at http://192.168.1.100:8080
  itms-services serve MyApp.ipa \
    --host-url https://192.168.1.100:8080 \
    --display-image https://192.168.1.100:8080/icon.png`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			viper.Set("ipa", args[0])
		}

		outputDir := viper.GetString("output")

		// Copy the ipa into the output directory so the HTTP server can serve it.
		if viper.IsSet("ipa") {
			ipaPath := viper.GetString("ipa")
			ipaName := filepath.Base(ipaPath)
			destPath := filepath.Join(outputDir, ipaName)

			hostURL := viper.GetString("host-url")
			if hostURL == "" {
				log.Fatal("--host-url is required for the serve command")
			}

			// Auto-set manifest-software-package to <host-url>/<ipa-filename>
			// if not already provided.
			if viper.GetString("manifest-software-package") == "" {
				viper.Set("manifest-software-package", hostURL+"/"+ipaName)
			}

			if err := copyFile(ipaPath, destPath); err != nil {
				log.Fatalf("failed to copy ipa to output dir: %v", err)
			}
		}

		generator.Build()

		addr := fmt.Sprintf(":%d", viper.GetInt("port"))
		log.Printf("Serving %s on http://0.0.0.0%s", outputDir, addr)
		log.Printf("Open the following URL on your iOS device to install:")
		log.Printf("  %s/%s", viper.GetString("host-url"), viper.GetString("index-name"))

		http.Handle("/", http.FileServer(http.Dir(outputDir)))
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	bindOptions(serveCmd, false, "output", "o", "./output",
		"Directory where generated files (and the ipa) are stored.")

	bindOptions(serveCmd, false, "host-url", "", "",
		"Public base URL of this server. Example: https://example.com")

	bindOptions(serveCmd, false, "title", "t", "Test app",
		"App title shown in the install page and alert dialog.")

	bindOptions(serveCmd, false, "bundle-identifier", "", "",
		"iOS bundle identifier (CFBundleIdentifier). Auto-read from ipa when provided.")

	bindOptions(serveCmd, false, "bundle-version", "", "1.0",
		"iOS bundle version (CFBundleVersion). Auto-read from ipa when provided.")

	bindOptions(serveCmd, false, "build-number", "", "0",
		"Build number shown on the install page.")

	bindOptions(serveCmd, false, "manifest-name", "", "manifest.plist",
		"Output filename for the manifest.")

	bindOptions(serveCmd, false, "manifest-software-package", "", "",
		"Full URL of the .ipa file. Defaults to <host-url>/<ipa-filename>.")

	bindOptions(serveCmd, false, "display-image", "", "",
		"Full URL of the 57×57 display image. Required.")

	bindOptions(serveCmd, false, "manifest-full-size-image", "", "",
		"Full URL of the 512×512 full-size image. Defaults to display-image.")

	bindOptions(serveCmd, false, "manifest-subtitle", "", "",
		"Subtitle shown in the install alert.")

	bindOptions(serveCmd, false, "index-name", "", "index.html",
		"Output filename for the install page.")

	bindOptions(serveCmd, false, "index-platform", "", "iOS",
		"Platform badge shown on the install page.")

	bindOptions(serveCmd, false, "index-branch", "", "trunk",
		"Branch badge shown on the install page.")

	bindOptions(serveCmd, false, "index-jumbotron-description", "",
		"Scan the QR code to open this page on your iPhone, click the button to install",
		"Description text on the install page.")

	bindOptions(serveCmd, false, "index-install-button-text", "", "Install",
		"Label of the install button.")

	serveCmd.PersistentFlags().IntP("port", "p", 8080, "TCP port to listen on.")
	_ = viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))
}

// copyFile copies src to dst, creating dst if it does not exist.
func copyFile(src, dst string) error {
	if src == dst {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		return err
	}
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o644)
}
