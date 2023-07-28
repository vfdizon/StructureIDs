package fileanalysis

import (
	"bufio"
	"os"
	"strings"
	"time"
)

type CSVReader struct {
	Directory string
	OutFile   string
	Verbose   bool

	GenePairs map[string]*string
}

func (c *CSVReader) Read() {
	c.GenePairs = make(map[string]*string)
	file, openerr := os.Open(c.Directory)

	if openerr != nil {
		panic(openerr)
	}

	fileReader := bufio.NewScanner(file)

	if c.Verbose {
		println("Reading file...")
	}

	startTime := time.Now

	fileReader.Scan()
	for fileReader.Scan() {
		line := fileReader.Text()

		parsedLine := strings.Split(line, ",")

		firstStruct := strings.Split(parsedLine[2], "_")[0]
		secondStruct := strings.Split(parsedLine[3], "_")[0]

		structPair := "\"" + firstStruct + ":" + secondStruct + "\""
		genePair := parsedLine[0] + "," + parsedLine[1]

		val, contains := c.GenePairs[genePair]

		if contains {
			*val = *val + "," + structPair
		} else {
			c.GenePairs[genePair] = &structPair
		}
	}

	if c.Verbose {
		println("Done reading file in", time.Since(startTime()).Milliseconds(), "ms")
	}
}

func (c *CSVReader) WriteToDotFile() {
	file, createErr := os.Create(c.OutFile)

	startTime := time.Now

	if c.Verbose {
		println("Writing to file", c.OutFile)
	}

	if createErr != nil {
		panic(createErr)
	}

	fileWriter := bufio.NewWriter(file)

	fileWriter.WriteString("digraph { \n")
	for key, val := range c.GenePairs {
		keyOut := strings.Replace(key, "\"", "", -1)
		fileWriter.WriteString("\"" + keyOut + "\" -> " + *val + " [arrowhead=none]; \n")
	}

	fileWriter.WriteString("}")

	fileWriter.Flush()
	file.Close()

	if c.Verbose {
		println("Done writing to file", c.OutFile, "in", time.Since(startTime()).Milliseconds(), "ms")
	}
}
