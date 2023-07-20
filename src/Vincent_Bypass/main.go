package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/vfdizon/Vincent_Bypass/fileanalysis"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the directory for the .csv file(s):")
	directory, _ := inputReader.ReadString('\n')
	directory = strings.TrimSpace(directory)

	fmt.Println(directory)

	csvSearcher := fileanalysis.CSVSearcher{
		Directory: directory,
	}

	startTime := time.Now()

	csvSearcher.Search()

	fmt.Println("Done in", time.Since(startTime))
}
