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
	FileName     string
	OutDirectory string
	Pairs        map[string]bool
	Verbose      bool
}

func (sid *StructureID) SearchPairs(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	startTime := time.Now()
	if sid.Verbose {
		fmt.Println("[goroutine] parsing pairs for Structure ID" + filepath.Base(sid.FileName))
		defer fmt.Println("[goroutine] done parsing pairs for Structure ID", filepath.Base(sid.FileName), "in", time.Since(startTime))
	}

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
		gene1 := parsedString[1]
		gene2 := parsedString[2]

		gene1 = strings.TrimSpace(gene1)
		gene2 = strings.TrimSpace(gene2)

		comparison := strings.Compare(gene1, gene2)

		if comparison == -1 {
			pair = gene1 + "," + gene2
		} else if comparison == 1 {
			pair = gene2 + "," + gene1
		} else {
			continue
		}

		sid.Pairs[pair] = true

	}

	sid.WritePairs(file.Name())
	if sid.Verbose {
		fmt.Println("[goroutine] wrote pairs for Structure ID", filepath.Base(sid.FileName))
	}

}

func (sid *StructureID) WritePairs(newFileName string) {
	os.Mkdir(sid.OutDirectory, os.ModePerm)

	csvFilename := strings.Replace(strings.Split(filepath.Base(sid.FileName), ".csv")[0], "dirty", "cleaned", -1) + ".csv"

	newFileName = filepath.Join(sid.OutDirectory, csvFilename)
	file, createErr := os.Create(newFileName)

	if createErr != nil {
		panic(createErr)
	}

	defer file.Close()

	csvWriter := bufio.NewWriter(file)

	for pair := range sid.Pairs {
		csvWriter.WriteString(pair + "\n")
	}

	csvWriter.Flush()

	if sid.Verbose {
		fmt.Println("[goroutine] wrote pairs for Structure ID", filepath.Base(sid.FileName))
	}

}
