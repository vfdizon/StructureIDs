package fileanalysis

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type MasterFile struct {
	Directory   string
	SharedPairs map[string]bool
}

func (csvs *MasterFile) Read(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	file, openErr := os.Open(csvs.Directory)
	if openErr != nil {
		fmt.Println(csvs.Directory)
		panic(openErr)
	}

	defer file.Close()

	csvs.SharedPairs = make(map[string]bool)
	fileReader := bufio.NewScanner(file)

	fileReader.Scan() // skip header
	for fileReader.Scan() {
		line := fileReader.Text()
		csvs.SharedPairs[line] = true
	}
}

func (csvs *MasterFile) ContainsPair(pair string) bool {
	_, contains := csvs.SharedPairs[pair]
	return contains
}
