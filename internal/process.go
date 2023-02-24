package internal

import (
	"os"
	"path"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func Process(fName string) (string, error) {

	data, err := os.ReadFile(fName)
	if err != nil {
		return "", err
	}

	convertData := convert(data)

	outFile, err := os.CreateTemp(path.Dir(fName), "coverted-*.md.html")
	if err != nil {
		return "", err
	}
	outFileName := outFile.Name()

	err = outFile.Close()
	if err != nil {
		return "", err
	}

	err = os.WriteFile(outFileName, convertData, os.ModeAppend)
	if err != nil {
		return "", err
	}

	return outFileName, nil
}

func convert(data []byte) []byte {
	unsafe := blackfriday.Run(data)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return html
}
