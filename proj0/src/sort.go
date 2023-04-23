package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"sort"
)

func readFrom(filePath string) [][]byte {
	data, err := ioutil.ReadFile(filePath)
	checkError(err)

	var size int = len(data) / 100
	records := make([][]byte, size)

	for i := range records {
		records[i] = data[i*100 : (i+1)*100]
	}

	return records
}

func writeInto(filepath string, slice [][]byte) {
	outputFile, err := os.Create(filepath)
	checkError(err)
	defer outputFile.Close()

	bufferedWriter := bufio.NewWriter(outputFile)
	defer bufferedWriter.Flush()

	for _, d := range slice {
		_, err = bufferedWriter.Write(d)
		checkError(err)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	inputFile := os.Args[1]
	records := readFrom(inputFile)

	sort.Slice(records, func(i, j int) bool {
		return bytes.Compare(records[i][:10], records[j][:10]) < 0
	})

	outputFile := os.Args[2]
	writeInto(outputFile, records)
	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])
}
