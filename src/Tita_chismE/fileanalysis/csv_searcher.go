package fileanalysis

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type CSVSearcher struct {
	StructIDDirectory   string
	MasterFileDirectory string
	StructIDs           map[*StructureID]bool
	Verbose             bool

	outDirectory     string
	masterFileReader *MasterFile
}

func (csvs *CSVSearcher) Search() {
	csvs.StructIDs = make(map[*StructureID]bool)
	csvs.outDirectory = filepath.Join(csvs.StructIDDirectory, "chismE_out")
	os.Mkdir(csvs.outDirectory, os.ModePerm)

	csvs.masterFileReader = &MasterFile{
		Directory: csvs.MasterFileDirectory,
	}

	var parseWaitGroup sync.WaitGroup

	dir, openErr := os.Open(csvs.StructIDDirectory)
	files, readErr := dir.Readdir(0)
	if openErr != nil || readErr != nil {
		panic(errors.New("error opening directory"))
	}

	defer dir.Close()
	parseWaitGroup.Add(1)
	go csvs.masterFileReader.Read(&parseWaitGroup)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ".csv") {
			parseWaitGroup.Add(1)

			structID := StructureID{
				FileName:     csvs.StructIDDirectory + file.Name(),
				OutDirectory: csvs.outDirectory,
				Verbose:      csvs.Verbose,
			}

			csvs.StructIDs[&structID] = true

			go structID.SearchPairs(&parseWaitGroup)
		}
	}

	parseWaitGroup.Wait()

	var removeWaitGroup sync.WaitGroup
	for structID := range csvs.StructIDs {
		removeWaitGroup.Add(1)
		go structID.RemoveSharedPairs(csvs.masterFileReader.SharedPairs, &removeWaitGroup)
	}

	removeWaitGroup.Wait()
}
