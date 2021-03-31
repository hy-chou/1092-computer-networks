package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	// (1)
	ln, err := net.Listen("tcp", ":12999")
	check(err)
	defer ln.Close()
	conn, err := ln.Accept()
	check(err)
	defer conn.Close()
	// (2)
	cReader := bufio.NewReader(conn)
	sizeRead, err := cReader.ReadString('\n')
	check(err)
	fmt.Printf("Upload file size: %s", sizeRead)
	uCnt, err := strconv.Atoi(sizeRead[:len(sizeRead)-1])
	check(err)
	oFile, err := os.Create("output.txt")
	check(err)
	defer oFile.Close()
	fWriter := bufio.NewWriter(oFile)
	// (3)
	oCnt, lineCnt := 0, 0
	for uCnt > 0 {
		lineRead, err := cReader.ReadString('\n')
		check(err)
		uCnt -= len(lineRead)
		lineCnt += 1
		// (4)
		len, err := fWriter.WriteString(strconv.Itoa(lineCnt) + " " + lineRead)
		check(err)
		fWriter.Flush()
		oCnt += len
		// (5)
	}
	// (6)
	fmt.Printf("Output file size: %d\n", oCnt)
	cWriter := bufio.NewWriter(conn)
	_, errr := cWriter.WriteString(sizeRead[:len(sizeRead)-1] + " bytes received, " + strconv.Itoa(oCnt) + " bytes file generated")
	check(errr)
	cWriter.Flush()
	// (7)
}
