package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/vfdizon/Tita_chismE/fileanalysis"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter directory of Structure ID files:")
	structIDDirectory, _ := inputReader.ReadString('\n')
	structIDDirectory = strings.TrimSpace(structIDDirectory)

	fmt.Println("Enter directory of master file:")
	masterFileDirectory, _ := inputReader.ReadString('\n')
	masterFileDirectory = strings.TrimSpace(masterFileDirectory)

	startTime := time.Now()
	csvSearcher := fileanalysis.CSVSearcher{
		StructIDDirectory:   structIDDirectory,
		MasterFileDirectory: masterFileDirectory,
	}

	csvSearcher.Search()

	fmt.Println("Done in", time.Since(startTime))

}
