package fileanalysis

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type StructureID struct {
	Directory       string
	FileName        string
	OutputDirectory string
	Complexity *int

	NodeConnections map[string]*string
}

func (s *StructureID) ReadFile() {
	s.NodeConnections = make(map[string]*string)

	file, openerr := os.Open(filepath.Join(s.Directory, s.FileName))
	if openerr != nil {
		panic(openerr)
	}

	defer file.Close()

	fileReader := bufio.NewScanner(file)

	fileReader.Scan()

	for fileReader.Scan() {
		line := fileReader.Text()
		parsedString := strings.Split(line, ",")

		val, contains := s.NodeConnections[parsedString[0]]
		if contains {
			*val = *val + "," + parsedString[1]
		} else {
			s.NodeConnections[parsedString[0]] = &parsedString[1]
		}
	}

	s.WriteFile()

}

func (s *StructureID) WriteFile() {
	outputFileName := strings.Split(s.FileName, ".csv")[0] + "_fingerprint.dot"
	os.Mkdir(s.OutputDirectory, os.ModePerm)
	outputFile, createErr := os.Create(filepath.Join(s.OutputDirectory, outputFileName))

	if createErr != nil {
		panic(createErr)
	}

	defer outputFile.Close()
	
	outputFileWriter := bufio.NewWriter(outputFile)

	outputFileWriter.WriteString("digraph {\n")
	for key, value := range s.NodeConnections {
		outputFileWriter.WriteString(key + " -> " + *value + "[arrowhead=none]; \n")
	}

	outputFileWriter.WriteString("}")
	outputFileWriter.Flush()
}

