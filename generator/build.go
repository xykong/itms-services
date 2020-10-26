package generator

import (
	"github.com/spf13/viper"
	"github.com/xykong/itms-services/ipa"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

// Build itms-services required index.html and manifest.plist
func Build() {
	BuildOptions()
	ParseIPA()
	BuildManifest()
	BuildIndex()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func BuildOptions() {

	hostUrl := viper.GetString("host-url")

	_, err := url.ParseRequestURI(hostUrl)
	check(err)

	// display-image
	v := viper.GetString("display-image")
	if len(v) == 0 {
		log.Fatal("display-image is empty")
	}

	if _, err := url.ParseRequestURI(v); err != nil {
		host, _ := url.Parse(hostUrl)
		host.Path = path.Join(host.Path, v)
		viper.Set("display-image", host.String())
	}

	// manifest-full-size-image
	v = viper.GetString("manifest-full-size-image")
	if len(v) == 0 {
		log.Print("manifest-full-size-image is empty, use display-image value.")
		v = viper.GetString("display-image")
		viper.Set("manifest-full-size-image", v)
	}

	if _, err := url.ParseRequestURI(v); err != nil {
		host, _ := url.Parse(hostUrl)
		host.Path = path.Join(host.Path, v)
		viper.Set("manifest-full-size-image", host.String())
	}

	// page-url
	host, _ := url.Parse(hostUrl)
	host.Path = path.Join(host.Path, viper.GetString("index-name"))
	viper.Set("page-url", host.String())

	// manifest-software-package
	v = viper.GetString("manifest-software-package")
	if len(v) == 0 {
		log.Fatal("manifest-software-package is empty")
	}

	if _, err := url.ParseRequestURI(v); err != nil {
		host, _ := url.Parse(hostUrl)
		host.Path = path.Join(host.Path, v)
		viper.Set("manifest-software-package", host.String())
	}
}

func ParseIPA() {
	if !viper.IsSet("ipa") {
		return
	}

	ipaPath := viper.GetString("ipa")

	info, err := ipa.ParseIpa(ipaPath)
	check(err)

	//infoString, _ := json.MarshalIndent(info, "", "\t")
	//log.Print(string(infoString))

	viper.Set("title", info["CFBundleDisplayName"])
	viper.Set("bundle-identifier", info["CFBundleIdentifier"])
	viper.Set("bundle-version", info["CFBundleVersion"])
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
		DisplayImage:     viper.GetString("display-image"),
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

	t, err := template.New("index").Parse(string(tpl))
	check(err)

	index := struct {
		Platform             string
		Branch               string
		BundleIdentifier     string
		BundleVersion        string
		BuildNumber          string
		Title                string
		JumbotronDescription string
		InstallButtonText    string
		DisplayImage         string
		IndexName            string
		ManifestName         string
		PageUrl              string
	}{
		DisplayImage:         viper.GetString("display-image"),
		Platform:             viper.GetString("index-platform"),
		Branch:               viper.GetString("index-branch"),
		BundleIdentifier:     viper.GetString("bundle-identifier"),
		BundleVersion:        viper.GetString("bundle-version"),
		BuildNumber:          viper.GetString("build-number"),
		Title:                viper.GetString("title"),
		JumbotronDescription: viper.GetString("index-jumbotron-description"),
		InstallButtonText:    viper.GetString("index-install-button-text"),
		IndexName:            viper.GetString("index-name"),
		ManifestName:         viper.GetString("manifest-name"),
		PageUrl:              viper.GetString("page-url"),
	}

	outputFilename := path.Join(viper.GetString("output"), viper.GetString("index-name"))
	_ = os.MkdirAll(filepath.Dir(outputFilename), os.ModePerm)
	output, err := os.Create(outputFilename)
	check(err)
	defer output.Close()

	err = t.Execute(output, index)
	check(err)
}
