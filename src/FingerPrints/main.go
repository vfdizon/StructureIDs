package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/vfdizon/fingerprints/fileanalysis"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter directory for all of the .csv files: ")
	directory, _ := inputReader.ReadString('\n')
	directory = strings.TrimSpace(directory)

	csvSearcher := fileanalysis.CSVSearcher{
		Directory:       directory,
		OutputDirectory: filepath.Join(directory, "fingerprints"),
	}
	csvSearcher.Search()

	//test
	//comment test
}
