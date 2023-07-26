package fileanalysis

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type StructureID struct {
	FileName string
	Pairs    map[string]bool
	Verbose  bool
	// SearchedPairs map[*StructureID]bool
}

func (sid *StructureID) SearchPairs(waitGroup *sync.WaitGroup) {
	startTime := time.Now()
	if sid.Verbose {
		fmt.Println("[goroutine] parsing pairs for Structure ID", sid.FileName)
		defer fmt.Println("[goroutine] done parsing pairs for Structure ID", sid.FileName, "in", time.Since(startTime))
	}
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

// func (structID *StructureID)
