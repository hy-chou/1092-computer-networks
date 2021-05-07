package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func handleConnection(c net.Conn) {
	// (6)
	defer c.Close()
	// (2)
	reader := bufio.NewReader(c)
	reqFileName := ""
	for {
		req, err := reader.ReadString('\n')
		check(err)
		if req == "\r\n" || reqFileName != "" {
			break
		}
		// (3)
		tokens := strings.Split(req, " ")
		reqFileName = strings.TrimLeft(tokens[1], "/")
	}
	// (5)
	if _, err := os.Stat(reqFileName); os.IsNotExist(err) {
		fmt.Println("File not found")
	} else {
		// (4)
		reqFile, err := os.Open(reqFileName)
		check(err)
		defer reqFile.Close()
		uStat, err := reqFile.Stat()
		check(err)
		uSize := strconv.FormatInt(uStat.Size(), 10)
		fmt.Println("File size = " + uSize)
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
		go handleConnection(conn)
	}
}
