package main

import (
	"flag"
	"strings"

	"github.com/vfdizon/Tita_pinkE/fileanalysis"
)

var (
	CSVFileDirectory string
	Verbose          bool
)

func init() {
	flag.StringVar(&CSVFileDirectory, "i", "", "Directory of the cleaned CSV files")
	flag.BoolVar(&Verbose, "v", false, "Verbose")
}

func main() {
	flag.Parse()
	directory := strings.TrimSpace(CSVFileDirectory)

	csvSearcher := fileanalysis.CSVSearcher{
		Directory: directory,
		Verbose:   Verbose,
	}

	csvSearcher.Search()
}
