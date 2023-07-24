package fileanalysis

import (
	"os"
	"sync"
)

type CSVSearcher struct {
	Directory       string
	OutputDirectory string
}

func (s *CSVSearcher) Search() {
	directory, openErr := os.Open(s.Directory)
	if openErr != nil {
		panic(openErr)
	}

	files, readErr := directory.Readdir(-1)
	if readErr != nil {
		panic(readErr)
	}

	structureIDs := make([]*StructureID, 0)
	for _, file := range files {
		if !file.IsDir() {
			structureID := &StructureID{
				Directory:       directory.Name(),
				FileName:        file.Name(),
				OutputDirectory: s.OutputDirectory,
			}

			structureIDs = append(structureIDs, structureID)
		}
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(structureIDs))

	for _, structureID := range structureIDs {
		go func(s *StructureID) {
			defer waitGroup.Done()
			s.ReadFile()
		}(structureID)
	}

	waitGroup.Wait()
}
