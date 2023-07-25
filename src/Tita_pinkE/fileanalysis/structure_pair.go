package fileanalysis

import (
	"path/filepath"
	"strings"
	"sync"
)

type StructurePair struct {
	StructureID1   *StructureID
	StructureID2   *StructureID
	ExportFilePath string

	exportCSV   *CSVWriter
	SharedPairs map[string]bool
	seenPairs   map[string]bool
	Name        string
}

func (sp *StructurePair) SearchForDuplicates(waitGroup *sync.WaitGroup) {
	sp.Name = strings.Split(filepath.Base(sp.StructureID1.FileName), ".csv")[0] + "," + strings.Split(filepath.Base(sp.StructureID2.FileName), ".csv")[0]
	defer waitGroup.Done()
	sp.SharedPairs = make(map[string]bool)
	sp.seenPairs = make(map[string]bool)

	for pair1, _ := range sp.StructureID1.Pairs {
		sp.seenPairs[pair1] = true
	}

	for pair2, _ := range sp.StructureID2.Pairs {
		sp.seenPairs[pair2] = true

		_, contains := sp.StructureID1.Pairs[pair2]
		if contains {
			sp.SharedPairs[pair2] = true
		}
	}

	sp.exportPairs()

}

func (sp *StructurePair) exportPairs() {
	fileName := strings.Split(filepath.Base(sp.StructureID1.FileName), ".csv")[0] + "_SHARED_" + filepath.Base(sp.StructureID2.FileName)

	sp.exportCSV = &CSVWriter{
		FileName:     fileName,
		OutDirectory: sp.ExportFilePath,
	}

	sp.exportCSV.CreateCSV("Gene1,Gene2")

	for pair, _ := range sp.SharedPairs {
		sp.exportCSV.WriteCSV(pair, sp.Name)
	}

	sp.exportCSV.CloseCSV()

}
