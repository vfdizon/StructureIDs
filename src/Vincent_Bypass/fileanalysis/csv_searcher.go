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
	Directory    string
	StructIDs    map[*StructureID]bool
	outDirectory string
	Verbose      bool
}

func (csvs *CSVSearcher) Search() {
	csvs.StructIDs = make(map[*StructureID]bool)
	csvs.outDirectory = filepath.Join(csvs.Directory, "clean_out")
	os.Mkdir(csvs.outDirectory, os.ModePerm)

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
				FileName:     csvs.Directory + file.Name(),
				OutDirectory: csvs.outDirectory,
				Verbose:      csvs.Verbose,
			}

			csvs.StructIDs[&structID] = true

			go structID.SearchPairs(&waitGroup)
		}

		if csvs.Verbose {
			fmt.Print("\n")
		}
	}

	waitGroup.Wait()
}
