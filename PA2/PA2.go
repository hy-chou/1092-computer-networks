package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	iName, oName := "", ""

	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &iName)
	iFile, err := os.Open(iName)
	defer iFile.Close()
	check(err)

	fmt.Printf("Output filename: ")
	fmt.Scanf("%s", &oName)
	oFile, err := os.Create(oName)
	defer oFile.Close()
	check(err)

	iFileScanner := bufio.NewScanner(iFile)
	oFileWriter := bufio.NewWriter(oFile)
	lineCount := 0
	for iFileScanner.Scan() {
		lineCount++
		_, _ = oFileWriter.WriteString(strconv.Itoa(lineCount) + " " + iFileScanner.Text() + "\n")
		oFileWriter.Flush()
	}
}
