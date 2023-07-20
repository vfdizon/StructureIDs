package fileanalysis

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type CSVSearcher struct {
	Directory string
	StructIDs map[*StructureID]bool

	seenPairs      map[string]bool
	uniquePairs    map[string]string
	duplicatePairs map[string]*int

	outDirectory      string
	sharedPairsWriter *CSVWriter
	uniquePairsWriter *CSVWriter
}

func (csvs *CSVSearcher) Search() {
	csvs.StructIDs = make(map[*StructureID]bool)
	csvs.outDirectory = filepath.Join(csvs.Directory, "out")
	os.Mkdir(csvs.outDirectory, os.ModePerm)

	csvs.sharedPairsWriter = &CSVWriter{
		FileName:     "shared_pairs.csv",
		OutDirectory: csvs.outDirectory,
	}

	csvs.uniquePairsWriter = &CSVWriter{
		FileName:     "unique_pairs.csv",
		OutDirectory: csvs.outDirectory,
	}

	csvs.sharedPairsWriter.CreateCSV("gene1,gene2,frequency")
	csvs.uniquePairsWriter.CreateCSV("gene1,gene2,structure_id")

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

func (csvs *CSVSearcher) WritePairs() {
	for duplicatedPair, frequency := range csvs.duplicatePairs {
		csvs.sharedPairsWriter.WriteCSV(duplicatedPair + "," + fmt.Sprintf("%d", *frequency))
	}

	for uniquePair, structID := range csvs.uniquePairs {
		csvs.uniquePairsWriter.WriteCSV(uniquePair + "," + structID)
	}

	csvs.sharedPairsWriter.CloseCSV()
	csvs.uniquePairsWriter.CloseCSV()
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
