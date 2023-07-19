package main

import (
	"github.com/vfdizon/structureids/fileanalysis"
)

func main() {
	csvSearcher := fileanalysis.CSVSearcher{
		Directory: "C:\\Users\\Vincent Dizon\\Documents\\sampleData\\",
	}

	csvSearcher.Search()

}
