package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vfdizon/Tita_geneE/fileanalysis"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the directory for the .csv file(s):")
	directory, _ := inputReader.ReadString('\n')
	directory = strings.TrimSpace(directory)

	csvSearcher := fileanalysis.CSVSearcher{
		Directory: directory,
	}

	csvSearcher.Search()
	csvSearcher.AnalyzePairs()
	csvSearcher.WritePairs()
}
