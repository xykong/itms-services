package generator

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

// Build itms-services required index.html and manifest.plist
func Build() {
	BuildManifest()
	BuildIndex()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Build manifest.plist
func BuildManifest() {
	tpl, err := Asset("template/manifest.plist")
	check(err)

	t, err := template.New("manifest").Parse(string(tpl))
	check(err)

	manifest := struct {
		SoftwarePackage  string
		DisplayImage     string
		FullSizeImage    string
		BundleIdentifier string
		BundleVersion    string
		Subtitle         string
		Title            string
	}{
		SoftwarePackage:  viper.GetString("manifest-software-package"),
		DisplayImage:     viper.GetString("manifest-display-image"),
		FullSizeImage:    viper.GetString("manifest-full-size-image"),
		BundleIdentifier: viper.GetString("bundle-identifier"),
		BundleVersion:    viper.GetString("bundle-version"),
		Subtitle:         viper.GetString("manifest-subtitle"),
		Title:            viper.GetString("title"),
	}

	outputFilename := path.Join(viper.GetString("output"), viper.GetString("manifest-name"))
	_ = os.MkdirAll(filepath.Dir(outputFilename), os.ModePerm)
	output, err := os.Create(outputFilename)
	check(err)
	defer output.Close()

	err = t.Execute(output, manifest)
	check(err)
}

// Build index.html
func BuildIndex() {
	tpl, err := Asset("template/index.html")
	check(err)

	// use asset data

	//	const tpl = `
	//<!DOCTYPE html>
	//<html>
	//    <head>
	//        <meta charset="UTF-8">
	//        <title>{{.Title}}</title>
	//    </head>
	//    <body>
	//        {{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	//    </body>
	//</html>`

	t, err := template.New("index").Parse(string(tpl))
	check(err)

	//<span class="badge badge-primary">iOS</span>
	//<span class="badge badge-dark">trunk</span>
	//<span class="badge badge-success">1.86</span>
	//<span class="badge badge-info">12345678</span>

	index := struct {
		Platform             string
		Branch               string
		BundleIdentifier     string
		BundleVersion        string
		BuildNumber          string
		Title                string
		JumbotronDescription string
		InstallButtonText    string
	}{
		Platform:             viper.GetString("index-platform"),
		Branch:               viper.GetString("index-branch"),
		BundleIdentifier:     viper.GetString("bundle-identifier"),
		BundleVersion:        viper.GetString("bundle-version"),
		BuildNumber:          viper.GetString("build-number"),
		Title:                viper.GetString("title"),
		JumbotronDescription: viper.GetString("index-jumbotron-description"),
		InstallButtonText:    viper.GetString("index-install-button-text"),
	}

	outputFilename := path.Join(viper.GetString("output"), viper.GetString("index-name"))
	_ = os.MkdirAll(filepath.Dir(outputFilename), os.ModePerm)
	output, err := os.Create(outputFilename)
	check(err)
	defer output.Close()

	err = t.Execute(output, index)
	check(err)
}
