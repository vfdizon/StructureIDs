package fileanalysis

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CSVWriter struct {
	FileName     string
	OutDirectory string
	CSVFile      *os.File
	FileWriter   *bufio.Writer
}

func (csvw *CSVWriter) CreateCSV(header string) {
	file, createError := os.Create(filepath.Join(csvw.OutDirectory, csvw.FileName))
	if createError != nil {
		panic(createError)
	}

	csvw.CSVFile = file

	fileWriter := bufio.NewWriter(file)
	csvw.FileWriter = fileWriter
	fileWriter.WriteString(header + "\n")
}

func (csvw *CSVWriter) WriteCSV(line ...string) {
	csvw.FileWriter.WriteString(strings.Join(line, ",") + "\n")
	csvw.FileWriter.Flush()
}

func (csvw *CSVWriter) CloseCSV() {
	csvw.FileWriter.Flush()
	csvw.CSVFile.Close()

	fmt.Println("[CSVWriter] closed CSV file")
}
