package main

import (
	"flag"

	"github.com/vfdizon/PairNodes/fileanalysis"
)

var (
	MasterFileDirectory string
	OutFile             string
	Verbose             bool
)

func init() {
	flag.StringVar(&MasterFileDirectory, "i", "", "Master file directory")
	flag.StringVar(&OutFile, "o", "", "Output file")
	flag.BoolVar(&Verbose, "v", false, "Verbose")
}

func main() {
	flag.Parse()

	csvReader := fileanalysis.CSVReader{
		Directory: MasterFileDirectory,
		OutFile:   OutFile,
		Verbose:   Verbose,
	}

	csvReader.Read()
	csvReader.WriteToDotFile()
}
