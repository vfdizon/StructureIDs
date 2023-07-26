package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/vfdizon/Tita_chismE/fileanalysis"
)

var (
	StructIDDirectory   string
	MasterFileDirectory string
	Verbose             bool
)

func init() {
	flag.StringVar(&StructIDDirectory, "i", "", "Directory of where the claned Structure ID files are located")
	flag.StringVar(&MasterFileDirectory, "m", "", "Directory of the master file containing the shared pairs")
	flag.BoolVar(&Verbose, "v", false, "Verbose")
}

func main() {
	flag.Parse()
	structIDDirectory := strings.TrimSpace(StructIDDirectory)

	masterFileDirectory := strings.TrimSpace(MasterFileDirectory)

	startTime := time.Now()
	csvSearcher := fileanalysis.CSVSearcher{
		StructIDDirectory:   structIDDirectory,
		MasterFileDirectory: masterFileDirectory,
		Verbose:             Verbose,
	}

	csvSearcher.Search()
	if Verbose {
		fmt.Println("Done in", time.Since(startTime))
	}

}

func ProgramIsVerbose() bool {
	return Verbose
}
