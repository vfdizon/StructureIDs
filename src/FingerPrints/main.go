package main

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/vfdizon/fingerprints/fileanalysis"
)

var (
	CSVFileDirectory string
)

func init() {
	flag.StringVar(&CSVFileDirectory, "i", "", "Directory of .csv files containing unique gene pairs for each Structure ID")
}

func main() {
	flag.Parse()
	directory := strings.TrimSpace(CSVFileDirectory)

	csvSearcher := fileanalysis.CSVSearcher{
		Directory:       directory,
		OutputDirectory: filepath.Join(directory, "fingerprints"),
	}
	csvSearcher.Search()
}
