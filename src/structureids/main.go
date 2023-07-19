package main

import (
	"github.com/vfdizon/structureids/fileanalysis"
)

func main() {
	csvSearcher := fileanalysis.CSVSearcher{
		Directory: "C:\\Users\\Vincent Dizon\\Documents\\sampleData\\out\\",
	}

	csvSearcher.Search()
	csvSearcher.AnalyzePairs()

}
