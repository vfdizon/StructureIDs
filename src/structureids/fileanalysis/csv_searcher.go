package fileanalysis

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type CSVSearcher struct {
	Directory string
	StructIDs map[int]*StructureID

	seenPairs      map[string]bool
	uniquePairs    map[string]int
	duplicatePairs map[string]*int
}

func (csvs *CSVSearcher) Search() {
	csvs.StructIDs = make(map[int]*StructureID)

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

			parsedFileName := strings.Split(file.Name(), "]")
			id, atoiErr := strconv.Atoi(strings.Replace(parsedFileName[0], "[", "", -1))
			if atoiErr != nil {
				panic(atoiErr)
			}

			structID := StructureID{
				FileName: csvs.Directory + file.Name(),
				ID:       id,
			}

			csvs.StructIDs[id] = &structID

			go structID.SearchPairs(&waitGroup)
		}
	}

	waitGroup.Wait()

	fmt.Println("all done")

}

func (csvs *CSVSearcher) AnalyzePairs() {
	csvs.seenPairs = make(map[string]bool)
	csvs.uniquePairs = make(map[string]int)
	csvs.duplicatePairs = make(map[string]*int)

	for id, structID := range csvs.StructIDs {

	}
}
