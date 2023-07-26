package main

import (
	"flag"
	"strings"

	"github.com/vfdizon/Tita_geneE/fileanalysis"
)

var (
	CSVDirectory string
	Verbose      bool
)

func init() {
	flag.StringVar(&CSVDirectory, "i", "", "Directory of the cleaned CSV Files")
	flag.BoolVar(&Verbose, "v", false, "Verbose")
}

func main() {
	flag.Parse()
	directory := strings.TrimSpace(CSVDirectory)

	csvSearcher := fileanalysis.CSVSearcher{
		Directory: directory,
		Verbose:   Verbose,
	}

	csvSearcher.Search()
}
