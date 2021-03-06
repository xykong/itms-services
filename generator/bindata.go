// Code generated for package generator by go-bindata DO NOT EDIT. (@generated)
// sources:
// template/index.html
// template/manifest.plist
package generator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x59\x77\xe2\x3a\x12\x7e\xcf\xaf\x50\xfb\x9e\x33\xe9\x3e\x89\x6d\xb6\x40\x6e\x1a\x67\x6e\x13\xd2\x09\x59\xc8\x42\xf6\x37\xd9\x2a\x6c\x19\x59\x72\x24\x99\x25\x1c\xfe\xfb\x1c\xdb\x38\x10\x1a\xce\xa4\xef\xcc\xd5\x03\x16\x55\x9f\x3e\xa9\x16\x55\xd9\xcd\x2f\xed\xab\xa3\xbb\xe7\xeb\x63\x14\xe8\x88\x1d\x6e\x35\xbf\x98\xa6\x4a\xe2\x58\x82\x52\xe8\xac\x77\xcf\x25\x28\xc1\x86\x40\x2e\xa8\x2b\xb1\x9c\xdc\xdf\x5e\x20\xd3\x3c\xdc\x6a\xa6\x70\xc4\x30\xf7\x1d\x03\xb8\x91\x0a\x00\x93\xc3\x2d\x84\x10\x4a\x39\xd0\x2d\xbc\x26\x54\x02\x41\x11\x68\x8c\x34\xf6\x55\xb6\x2e\xd3\x67\x22\x2f\xc0\x52\x81\x76\x8c\x44\xf7\xcd\x7d\xc3\x5e\xd6\x71\x1c\x81\x63\x0c\x29\x8c\x62\x21\xb5\x81\x3c\xc1\x35\x70\xed\x18\x23\x4a\x74\xe0\x10\x18\x52\x0f\xcc\xec\xcf\x2e\xa2\x9c\x6a\x8a\x99\xa9\x3c\xcc\xc0\x29\xef\x22\x15\x48\xca\x07\xa6\x16\x66\x9f\x6a\x87\x8b\x94\x7b\x71\xb0\x96\x10\x5a\x69\x89\x63\x74\xd4\xeb\x2d\xce\xc4\x28\x1f\x20\x09\xcc\x31\x94\x9e\x30\x50\x01\x80\x36\x32\x55\x3e\x02\x09\x7d\xc7\x08\xb4\x8e\xd5\x81\x6d\x7b\x84\x5b\xa1\x22\xc0\xe8\x50\x5a\x1c\xb4\xcd\xe3\xc8\x76\x0b\xe6\xbf\x6a\xd6\x9e\x55\xb5\x09\x55\xda\xf6\x94\x5a\x28\xac\x88\x72\xcb\x53\x6a\x99\x98\x72\x0d\xbe\xa4\x7a\xe2\x18\x2a\xc0\xd5\xfd\x9a\x79\xf7\xb4\xaf\x2b\x8d\x63\xef\xf6\xb8\x0a\x36\x0d\xee\x1b\x6f\xd1\xcd\xf8\x81\x7b\xed\x1f\x93\xbd\xa4\x73\xfe\x56\x93\xc7\x03\xbf\xf3\x04\x97\x40\x6a\x97\xa5\x90\xf5\x3b\xed\xeb\xa1\x5f\x4f\x5e\xcf\x3b\x95\xf1\x93\xac\x2c\xb3\x7b\x52\x28\x25\x24\xf5\x29\x77\x0c\xcc\x05\x9f\x44\x22\x51\x1f\x3d\xd2\x54\x9e\xa4\xb1\x46\x4a\x7a\x1f\x2d\x94\x78\xe4\x53\x6d\x79\x22\xb2\x09\x1e\x52\xa2\x02\x1a\x85\xca\x7e\x95\x9e\x20\x10\x2a\xdb\x0f\xcc\x18\xfb\x50\x48\x32\xeb\x42\x65\x1c\x36\xed\x9c\xf1\x70\xe1\xde\x0d\x5b\x84\xca\xf2\x98\x48\x48\x9f\x61\x09\xd9\x46\x38\xc4\x63\x9b\x51\x77\x69\x9b\xb2\x55\xb2\x4a\x1b\xf7\x98\x1b\xa2\xa9\x66\x70\x38\x9d\x5a\x77\xe9\x64\x36\x6b\xda\xb9\x64\xab\x69\xe7\x99\xd9\x74\x05\x99\x1c\x6e\x6d\x35\x39\x1e\x22\x8f\x61\xa5\x1c\x83\xe3\xa1\x8b\x25\xca\x1f\x26\xc1\x72\x80\x5c\x3f\x7b\x1a\x4b\xa9\xdc\xcd\x51\xf3\x24\xcc\xd3\xdf\xe6\x78\x98\x92\x11\xfa\x4e\x96\xea\x31\xe5\x20\x8b\xb5\x4b\x3a\x29\x46\x28\x4c\x94\xa6\xfd\x89\x39\xe7\x31\x3d\xe0\xfa\x1d\xbc\xba\xc0\x13\xcc\x8c\x88\x89\x13\x2d\x96\x10\xab\xa8\x30\x89\x5c\xa1\xa5\xe0\x2b\x98\x0c\x17\x94\x0b\x18\xa1\x2a\x66\x78\x62\xd6\xd6\xc0\x32\x28\x8d\xfc\x3c\x34\xd3\xa9\xd5\xce\xc1\x9d\x08\xfb\x30\x9b\x19\x28\x00\xea\x07\xda\x31\xea\x35\x03\x61\xa6\x1d\xc3\x28\x68\xfb\x4c\x60\x6d\x32\xe8\x6b\x03\x65\x97\xc6\x31\x5c\x21\x09\x48\x53\x62\x42\x13\x75\x80\xca\xf5\x78\xfc\xbd\xb8\xda\xab\x63\x11\xab\x5f\xcf\x6e\x07\xe5\x79\x60\x3f\x9a\x54\xdb\x60\x80\x8a\x31\x2f\x8e\xe5\x62\xe2\x03\xca\x7e\xcd\x58\xd2\x08\xcb\x89\x91\x26\xc6\x35\xc3\xba\x2f\x64\x94\xe6\x46\x8a\xff\x4d\xaa\x3c\x2b\xa6\x53\xab\x25\x31\xf7\x82\xbf\xc9\xa2\x12\xcf\x03\xa5\x72\xa2\x84\x13\x06\x0f\x20\x15\x15\xfc\x6f\xf2\x51\xde\x17\x73\x32\xca\x48\x37\x89\x5c\x90\x9b\xa9\x9a\x76\xea\xc1\x5f\xc5\x71\x41\xcd\x00\x93\x0d\x49\x32\x9d\x5a\x67\x45\xba\xb5\x21\xbf\x7c\xd9\xb9\xd7\xec\x12\xaf\x0d\x9e\x2c\x76\x89\xd2\x54\xb4\xd7\x61\x3e\x77\x65\x10\x25\x8e\x91\xd7\x83\xb4\x10\x10\x3a\x5c\x6b\xd4\xe1\x86\x83\xe0\x6c\x3d\xd5\x91\x32\x15\xc8\xb4\x95\x7c\x28\xc8\xef\xa3\x70\xb7\xe6\xc8\xd5\xbc\xc8\xa5\x6c\xce\xfc\xec\xe1\x32\xe1\x0d\xd6\xae\xcd\xdb\xc5\x87\x3d\x0e\x6c\xfb\xdf\xd8\x4b\x7d\xe6\x10\x31\xe2\x4c\x60\x62\x46\x98\xd3\x3e\x28\xfd\xaf\x44\x32\xa7\xa8\x8b\x30\xc6\x51\xcc\xf2\x72\x48\xb9\xd2\x98\x31\xbb\x00\x5a\x31\xa3\x4a\xaf\xdd\x52\x8a\xec\x06\x26\x5a\xaf\xad\x07\x28\x8f\x61\x27\x27\x6c\x65\xb0\x3b\x18\xeb\xb5\x01\xc4\x2b\x7e\x9b\x3b\x79\xcd\xdf\xf9\xb4\x88\x42\x56\x30\xaf\xb2\xcc\xc0\x0c\x9d\xe1\x21\xee\x65\x99\xf2\x1d\x79\x81\x10\x0a\x90\xe0\x80\x44\x1f\xe9\x00\x90\x1e\x89\x2f\x59\x35\x5d\x5e\x85\xca\x07\x28\xbc\x49\x40\x4e\x10\xe6\x64\xa9\x5b\xe7\xb7\x05\x7d\xa5\xdc\x63\x09\x01\x85\xae\x45\x1c\x83\xfc\x96\xd7\xe3\xb5\xfd\x25\xed\x17\xe1\x6b\xca\x95\xf9\x32\x9f\x9a\x55\x6b\xcf\x2a\x5b\x8a\xd1\xa8\x68\x26\xef\x76\xfd\xda\x8b\xdb\xfd\x27\xf2\x56\x09\xf4\xf5\x69\x89\xa9\x5e\x4f\xed\xf1\xa3\xbb\x38\x09\xed\xb7\x49\xed\x68\xe7\xea\x24\xc6\x91\xf8\xf9\x30\xa9\xee\x5f\x3e\xb4\xf8\xf1\x4e\xc7\x75\x1f\x9e\xef\x61\xb4\x73\x25\x8f\x9e\xf0\xed\xa0\x1f\x2e\xb8\x37\x74\xe2\xd4\x75\xef\x9d\x6c\x63\x27\xfe\xdc\xbb\x46\xb8\xfc\xaa\xe1\x66\xfe\xfa\x84\x89\x81\xd8\x09\x1b\xe1\xe4\xf1\x7c\xbf\xdf\xbd\x81\x9d\x1f\xe5\xca\xa9\xbb\xff\x23\xb8\x7d\xad\xd4\x2f\xe4\x8b\x7d\x16\x7b\xf7\x27\x27\x57\x7c\xe7\xb9\x71\xab\x46\xd0\x95\xba\x6b\xeb\xe3\xea\xa5\x38\x6f\xbc\xc0\x4b\x7b\x32\xfe\x4d\x13\x97\x42\x5d\x29\x42\xbd\x3b\x0f\xa6\x15\xaa\xdd\x95\xb0\x9f\xf5\xfe\xc7\xe0\xfe\x13\x41\xdd\x68\xe9\xc2\xd0\xcf\x87\x32\x2e\x4c\xff\xab\x6c\x95\xeb\x56\x39\x8f\x65\x12\x91\x42\xb3\xd9\x90\x3f\x6d\x09\x3f\xef\x4e\x7e\x3c\xee\x57\x8f\x1f\x2b\xb7\xed\xa4\xd2\x2b\x3d\x9c\xe3\xce\x1b\x8e\xab\xa7\xf5\x3a\x7b\x39\xdd\x2f\x5f\x8b\x67\xf6\x33\x70\x4f\xee\x77\xea\xad\x97\xb8\x7e\xd2\xe0\x34\x69\x54\xf7\x7a\x83\x06\xeb\xfe\x7f\x0d\xf9\xef\x39\xb9\xd9\x92\x51\xf9\xa6\x26\xe4\x73\xd8\xba\xe1\xc4\x1b\x88\xfa\x25\x8d\x1e\xdc\xb7\xe7\x92\xf6\xe3\x5a\xfc\xd8\xaa\xb1\x97\x06\x93\xd5\xd2\xe3\xf9\x5b\x69\x28\x6d\xfc\x78\x1e\x3c\x91\x56\x37\xea\xba\x7b\xed\x3f\x2b\xc3\x86\xfa\x8c\x25\x79\xa9\x99\x9b\xa3\x27\x31\x38\x86\x86\xb1\xb6\x43\x3c\xc4\xb9\x74\x5e\x36\x3d\xc1\x95\x46\x79\x83\x41\x0e\xe2\x30\x42\x37\xb7\x47\x82\xc0\x57\x22\xbc\x24\x02\xae\x2d\x1f\xf4\x31\x83\x74\xda\x9a\x74\xc8\xd7\xed\x1c\xbc\xfd\x6d\x17\x4d\xdf\xaf\x83\x6d\xa3\x94\xff\x00\x8d\x28\x27\x62\x64\x31\xe1\xe1\x34\xef\xad\xb4\x39\xec\xbe\xc3\x72\x4c\xfa\xea\x75\x8d\x7d\xb8\x97\x6c\x36\x33\x16\xda\xec\x23\xe7\x00\x55\xf6\xea\x0b\x59\xfe\x4a\xb6\x22\xf4\x04\x13\xb2\x8d\xe5\xe0\x00\x6d\xff\x51\xca\xc6\xf6\x8a\xf6\x22\x5f\xb6\xfd\x47\x3f\x1b\x1f\xd4\x52\x82\xa7\x2f\x60\x08\xec\x60\x6e\xac\x75\xb4\x24\xb4\x4e\x73\xf0\xec\xdb\xf7\xbc\x33\x30\xd0\x88\xe7\xfe\xd9\xe8\x94\x0f\xed\x6f\xfb\xdb\xf7\x6c\x61\xba\x28\xf3\x00\x72\xd0\x6f\x36\x48\x03\xed\xac\xf5\xa5\x25\x21\x66\xd8\x83\xaf\x76\xd6\xe0\x08\x8c\xbb\x38\x82\xd9\xcc\xde\x45\xdb\xd3\xa9\x75\x39\xe7\xc8\x85\xe9\x39\x96\xeb\x91\x9d\x7f\x17\x34\xed\xfc\x03\xf8\x3f\x01\x00\x00\xff\xff\x72\x71\xae\x37\x11\x0f\x00\x00")

func templateIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateIndexHtml,
		"template/index.html",
	)
}

func templateIndexHtml() (*asset, error) {
	bytes, err := templateIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateManifestPlist = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\x51\x6b\xdb\x30\x10\xc7\x9f\xed\x4f\xa1\xe9\x5d\x56\xf2\x36\x86\xea\xb2\x36\x2d\x04\xc2\x66\x88\x3b\xd8\xa3\x1a\x5d\xd2\x23\xb2\x6c\xa4\xf3\xb2\xb4\xf4\xbb\x0f\xc7\x6e\xa2\x74\x5e\xd7\x87\xed\x4d\x48\xbf\xff\xdd\xcf\xc7\x59\x5d\xfe\xac\x2c\xfb\x01\x3e\x60\xed\x2e\xf8\x34\x9b\x70\x06\x6e\x55\x1b\x74\x9b\x0b\x7e\x57\xde\x8a\x8f\xfc\x32\x4f\xd5\x87\xd9\xd7\xeb\xf2\x7b\x71\xc3\x1a\x8b\x81\x58\x71\x77\xb5\x98\x5f\x33\x2e\xa4\xfc\xdc\x34\x16\xa4\x9c\x95\x33\x56\x2c\xe6\xcb\x92\x4d\xb3\x89\x94\x37\x5f\x38\xe3\x0f\x44\xcd\x27\x29\x77\xbb\x5d\xa6\x3b\x2a\x5b\xd5\x55\x07\x06\x59\xf8\xba\x01\x4f\xfb\x05\x06\x12\xd3\x6c\x92\x19\x32\x3c\x4f\x55\x5f\xfd\x4c\x27\x4f\x95\xc1\x15\xe5\x69\xa2\xb6\xb0\xcf\x91\xa0\x0a\x4a\x76\xc7\x34\x51\xda\x7b\xdd\x1d\x92\x17\x26\xe9\x29\x1d\x02\xd0\x11\x4b\x22\x30\x46\x07\x78\x8b\xce\x9c\xd0\x24\x51\x81\x3c\xba\x4d\x1e\xea\x35\xed\xb4\x07\xd1\xe8\xd5\x56\x6f\x40\xc9\xe1\x21\xca\xb6\xde\x8e\x45\x9f\x9e\xb2\xe5\x90\x2e\xfa\xf0\xf3\xf3\x79\x5c\xc9\x93\xc6\xbb\x8d\x0c\x86\xc6\xea\xbd\xc0\xea\x0f\x3a\x0e\xc0\x04\x11\x1e\xd0\xc1\x59\x9e\x7c\x0b\xf2\x7d\xde\xb3\xbe\xc7\xbc\xfa\x57\xd2\xeb\xd6\x5a\x11\xf0\x11\xfe\xab\xf6\x6d\x6b\xed\x12\x1f\xe1\x6f\xde\x4a\x9e\x56\xe1\x50\xb2\x02\xd2\x46\x93\x8e\x96\x25\xfa\xc8\xee\xee\xbe\x75\xc6\x82\x40\x03\x8e\x70\x8d\xe0\x23\x85\xc8\xe0\xea\x80\xcd\x8f\xd4\x6f\x12\x51\xad\x61\xc3\xdf\x2a\xf4\xad\x47\x46\xab\xbc\x1a\xf5\xeb\x85\x1d\x49\x84\xf6\x9e\x90\x2c\x8c\x77\x5c\x0e\xaf\xa3\xcd\xde\xc8\x95\x23\xa1\xd3\xac\x8f\xa7\xe3\xcc\x5f\x6e\x94\x3c\xfc\xe8\x79\xfa\x2b\x00\x00\xff\xff\x83\xed\xd1\x1c\x7f\x04\x00\x00")

func templateManifestPlistBytes() ([]byte, error) {
	return bindataRead(
		_templateManifestPlist,
		"template/manifest.plist",
	)
}

func templateManifestPlist() (*asset, error) {
	bytes, err := templateManifestPlistBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/manifest.plist", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/index.html":     templateIndexHtml,
	"template/manifest.plist": templateManifestPlist,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"template": &bintree{nil, map[string]*bintree{
		"index.html":     &bintree{templateIndexHtml, map[string]*bintree{}},
		"manifest.plist": &bintree{templateManifestPlist, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
