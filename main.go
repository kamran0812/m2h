package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	m2h "github.com/kamran0812/m2h/internal"
)

func main() {
	fName := flag.String("fname", "", "File name of Markdown file")
	preview := flag.Bool("p", false, "Preview converted file")
	save := flag.Bool("s", false, "Save converted file")

	flag.Parse()

	if *fName == "" {
		flag.Usage()
		return
	}

	err := run(*fName, *preview, *save)
	if err != nil {
		log.Panic(err)
	}
}

func run(fName string, preview bool, save bool) error {

	outPutFileName, err := m2h.Process(fName)
	if err != nil {
		return err
	}
	if preview {
		return m2h.Preview(outPutFileName)
	}
	if save {
		fmt.Println("File saved:", outPutFileName)
		return nil
	}

	err = os.Remove(outPutFileName)
	if err != nil {
		return err
	}

	return nil
}
