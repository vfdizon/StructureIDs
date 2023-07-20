package fileanalysis

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type CSVSearcher struct {
	Directory string
	StructIDs map[*StructureID]bool

	structurePairs []*StructurePair
	outDirectory   string
}

func (csvs *CSVSearcher) Search() {
	csvs.StructIDs = make(map[*StructureID]bool)
	csvs.outDirectory = filepath.Join(csvs.Directory, "geneE_out")
	os.Mkdir(csvs.outDirectory, os.ModePerm)

	var parseWaitGroup sync.WaitGroup

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
			parseWaitGroup.Add(1)

			structID := StructureID{
				FileName: csvs.Directory + file.Name(),
			}

			csvs.StructIDs[&structID] = true

			go structID.SearchPairs(&parseWaitGroup)
		}
	}

	parseWaitGroup.Wait()

	csvs.createPairs()

	var searchWaitGroup sync.WaitGroup
	searchWaitGroup.Add(len(csvs.structurePairs))

	for _, structIDPair := range csvs.structurePairs {
		go structIDPair.SearchForDuplicates(&searchWaitGroup)
	}

	searchWaitGroup.Wait()

}

func (csvs *CSVSearcher) createPairs() {
	csvs.structurePairs = []*StructurePair{}
	structIDsCopy := csvs.StructIDs

	for structID1, _ := range structIDsCopy {
		for structID2, _ := range structIDsCopy {
			if structID1 == structID2 {
				continue
			}

			structPair := StructurePair{
				StructureID1:   structID1,
				StructureID2:   structID2,
				ExportFilePath: csvs.outDirectory,
			}

			csvs.structurePairs = append(csvs.structurePairs, &structPair)

		}
		delete(csvs.StructIDs, structID1)
	}
}
