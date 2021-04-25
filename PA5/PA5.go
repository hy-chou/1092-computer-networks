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

	for {
		conn, err := ln.Accept()
		check(err)
		// (2)
		cReader := bufio.NewReader(conn)
		sizeRead, err := cReader.ReadString('\n')
		check(err)
		fmt.Printf("Upload file size: %s", sizeRead)
		uCnt, err := strconv.Atoi(sizeRead[:len(sizeRead)-1])
		check(err)

		oFile, err := os.Create("whatever.txt")
		check(err)
		fWriter := bufio.NewWriter(oFile)
		oCnt, lineCnt := 0, 0
		for uCnt > 0 {
			// (3)
			lineRead, err := cReader.ReadString('\n')
			check(err)
			uCnt -= len(lineRead)
			lineCnt += 1
			// (4)
			len, err := fWriter.WriteString(strconv.Itoa(lineCnt) + " " + lineRead)
			check(err)
			fWriter.Flush()
			oCnt += len
		}
		// (5)
		fmt.Printf("Output file size: %d\n", oCnt)
		// (6)
		cWriter := bufio.NewWriter(conn)
		_, err = cWriter.WriteString(sizeRead[:len(sizeRead)-1] + " bytes received, " + strconv.Itoa(oCnt) + " bytes file generated")
		check(err)
		cWriter.Flush()

		oFile.Close()
		conn.Close()
	}
	// (7)
}
