package fileanalysis

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

type CSVSearcher struct {
	Directory string
	StructIDs map[*StructureID]bool

	seenPairs      map[string]bool
	uniquePairs    map[string]string
	duplicatePairs map[string]*int
}

func (csvs *CSVSearcher) Search() {
	csvs.StructIDs = make(map[*StructureID]bool)

	var waitGroup sync.WaitGroup

	dir, openErr := os.Open(csvs.Directory)
	files, readErr := dir.Readdir(0)
	if openErr != nil || readErr != nil {
		panic(errors.New("error opening directory"))
	}

	defer dir.Close()

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ".csv") {
			waitGroup.Add(1)

			structID := StructureID{
				FileName: csvs.Directory + file.Name(),
			}

			csvs.StructIDs[&structID] = true

			go structID.SearchPairs(&waitGroup)
		}
	}

	waitGroup.Wait()
	fmt.Println("all done")

}

func (csvs *CSVSearcher) AnalyzePairs() {
	csvs.seenPairs = make(map[string]bool)
	csvs.uniquePairs = make(map[string]string)
	csvs.duplicatePairs = make(map[string]*int)

	for structID, _ := range csvs.StructIDs {
		for pair, _ := range structID.Pairs {
			_, contains := csvs.seenPairs[pair]

			if contains {
				csvs.handleDuplicatePairs(pair)
			} else {
				csvs.seenPairs[pair] = true
				csvs.uniquePairs[pair] = structID.FileName
			}

		}
	}
}

func (csvs *CSVSearcher) handleDuplicatePairs(duplicatePair string) {
	val, contains := csvs.duplicatePairs[duplicatePair]
	if !contains {
		csvs.duplicatePairs[duplicatePair] = new(int)
		*csvs.duplicatePairs[duplicatePair] = 2
	} else {
		*val++
	}

	delete(csvs.uniquePairs, duplicatePair)

}
