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
}

func (sid *StructureID) SearchPairs(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	startTime := time.Now()

	fmt.Println("[goroutine] parsing pairs for Structure ID" + filepath.Base(sid.FileName))
	defer fmt.Println("[goroutine] done parsing pairs for Structure ID", filepath.Base(sid.FileName), "in", time.Since(startTime))

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

		if strings.Compare(gene1, gene2) < 0 {
			pair = gene1 + "," + gene2
		} else {
			pair = gene2 + "," + gene1
		}

		sid.Pairs[pair] = true

	}

	sid.WritePairs(file.Name())
	fmt.Println("[goroutine] wrote pairs for Structure ID", filepath.Base(sid.FileName))

}

func (sid *StructureID) WritePairs(newFileName string) {
	err := os.Mkdir(sid.OutDirectory, os.ModePerm)

	if err != nil {
		fmt.Println("directory already exists")
	}

	newFileName = filepath.Join(sid.OutDirectory, strings.Split(filepath.Base(newFileName), ".csv")[0]+"_CLEANED.csv")

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

	fmt.Println("[goroutine] wrote pairs for Structure ID", filepath.Base(sid.FileName))

}
