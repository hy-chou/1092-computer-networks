package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleConnection(c net.Conn) {
	cReader := bufio.NewReader(c)
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
		lineRead, err := cReader.ReadString('\n')
		check(err)
		uCnt -= len(lineRead)
		lineCnt += 1
		len, err := fWriter.WriteString(strconv.Itoa(lineCnt) + " " + lineRead)
		check(err)
		fWriter.Flush()
		oCnt += len
	}
	fmt.Printf("Output file size: %d\n", oCnt)
	cWriter := bufio.NewWriter(c)
	_, err = cWriter.WriteString(sizeRead[:len(sizeRead)-1] + " bytes received, " + strconv.Itoa(oCnt) + " bytes file generated")
	check(err)

	cWriter.Flush()
	oFile.Close()
	c.Close()
	// (2)
	time.Sleep(5 * time.Second)
}

func main() {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":12999")
	check(err)
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		check(err)
		// (1)
		go handleConnection(conn)
	}
}
