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
	// (1)
	conn, err1 := net.Dial("tcp", "127.0.0.1:12999")
	check(err1)
	defer conn.Close()
	// (2)
	uName := ""
	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &uName)
	uFile, err2 := os.Open(uName)
	check(err2)
	defer uFile.Close()
	// (3)
	uStat, err3 := uFile.Stat()
	check(err3)
	uSize := strconv.FormatInt(uStat.Size(), 10)
	writer := bufio.NewWriter(conn)
	_, err4 := writer.WriteString(uSize + "\n")
	check(err4)
	fmt.Printf("Send the file size first: %s\n", uSize)
	writer.Flush()
	// (4)
	uFileScanner := bufio.NewScanner(uFile)
	for uFileScanner.Scan() {
		_, _ = writer.WriteString(uFileScanner.Text() + "\n")
		writer.Flush()
	}
	// (5)
	scanner := bufio.NewScanner(conn)
	// (6)
	for scanner.Scan() {
		fmt.Printf("Server says: %s\n", scanner.Text())
	}
	// (7)
}
