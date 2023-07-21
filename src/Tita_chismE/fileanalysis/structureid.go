package fileanalysis

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type StructureID struct {
	FileName string
	Pairs    map[string]bool

	OutDirectory string
	outWriter    *CSVWriter
}

func (sid *StructureID) SearchPairs(waitGroup *sync.WaitGroup) {
	startTime := time.Now()

	fmt.Println("[goroutine] parsing pairs for Structure ID", sid.FileName)
	defer fmt.Println("[goroutine] done parsing pairs for Structure ID", sid.FileName, "in", time.Since(startTime))

	defer waitGroup.Done()
	file, openError := os.Open(sid.FileName)
	sid.Pairs = make(map[string]bool)
	defer file.Close()
	if openError != nil {
		panic(openError)
	}

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		parsedString := strings.Split(line, ",")
		if len(parsedString) == 0 {
			continue
		}

		var pair string
		gene1 := parsedString[0]
		gene2 := parsedString[1]

		if strings.Compare(gene1, gene2) < 0 {
			pair = gene1 + "," + gene2
		} else {
			pair = gene2 + "," + gene1
		}

		sid.Pairs[pair] = true
	}

}

func (sid *StructureID) exportPairs() {
	fileName := strings.Split(filepath.Base(sid.FileName), ".csv")[0] + "_UNIQUE.csv"
	sid.outWriter = &CSVWriter{
		FileName:     fileName,
		OutDirectory: sid.OutDirectory,
	}

	sid.outWriter.CreateCSV("gene1,gene2")
	for pair := range sid.Pairs {
		sid.outWriter.WriteCSV(pair)
	}

	sid.outWriter.CloseCSV()
}

func (sid *StructureID) RemoveSharedPairs(sharedPairs map[string]bool, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for pair := range sid.Pairs {
		_, contains := sharedPairs[pair]
		if contains {
			delete(sid.Pairs, pair)
		}
	}

	sid.exportPairs()
}
