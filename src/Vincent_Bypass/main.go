package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/vfdizon/Vincent_Bypass/fileanalysis"
)

var (
	CSVFileDirectory string
	Verbose          bool
)

var Test bool

func init() {
	flag.StringVar(&CSVFileDirectory, "i", "", "Directory of \"dirty\" CSV files to search")
	flag.BoolVar(&Verbose, "v", false, "Verbose")
}

func main() {
	flag.Parse()
	directory := strings.TrimSpace(CSVFileDirectory)

	if Verbose {
		fmt.Println("Reading CSV files from:", directory)
	}

	csvSearcher := fileanalysis.CSVSearcher{
		Directory: directory,
		Verbose:   Verbose,
	}

	startTime := time.Now()

	csvSearcher.Search()

	if Verbose {
		fmt.Println("Done in", time.Since(startTime))
	}
}
