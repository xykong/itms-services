package ipa

import (
	"archive/zip"
	"errors"
	"howett.net/plist"
	"io"
	"log"
	"regexp"
)

//ParseIpa : It parses the given ipa and returns a map from the contents of Info.plist in it
func ParseIpa(name string) (map[string]interface{}, error) {

	r, err := zip.OpenReader(name)
	if err != nil {
		log.Println("Error opening ipa/zip ", err.Error())
		return nil, err
	}
	defer r.Close()

	re := regexp.MustCompile(`Payload/[^/]*/Info\.plist`)

	for _, file := range r.File {

		if file.FileInfo().Name() == "Info.plist" {

			if !re.MatchString(file.Name) {
				continue
			}

			rc, err := file.Open()
			if err != nil {
				log.Println("Error opening Info.plist in zip", err.Error())
				return nil, err
			}
			buf := make([]byte, file.FileInfo().Size())
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				log.Println("Error reading Info.plist", err.Error())
				return nil, err
			}

			var infoMap map[string]interface{}
			_, err = plist.Unmarshal(buf, &infoMap)
			if err != nil {
				log.Println("Error reading Info.plist", err.Error())
				return nil, err
			}

			return infoMap, nil
		}
	}

	//goland:noinspection GoErrorStringFormat
	return nil, errors.New("Info.plist not found")
}
