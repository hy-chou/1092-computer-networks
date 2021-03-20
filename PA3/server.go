package main

import (
	"bufio"
	"fmt"
	"net"
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
	ln, _ := net.Listen("tcp", ":12999")
	defer ln.Close()

	connectionNum := 0
	for {
		conn, _ := ln.Accept()
		connectionNum += 1
		fmt.Printf("\n(%d) Connected!\n", connectionNum)
		// (2)
		// (3)
		scanner := bufio.NewScanner(conn)
		fileSize := ""
		if scanner.Scan() {
			fileSize = scanner.Text()
			fmt.Println("Know the file size first: " + fileSize)
		}
		// (4)
		cnt := 0
		fileContent := ""
		for scanner.Scan() {
			fileContent = scanner.Text()
			// fmt.Println(fileContent)
			cnt += len(fileContent) + 1
			if strconv.Itoa(cnt) == fileSize {
				break
			}
		}
		// (5)
		writer := bufio.NewWriter(conn)
		newline := fmt.Sprintf("%d bytes received\n", cnt)
		_, errw := writer.WriteString(newline)
		check(errw)
		writer.Flush()
		fmt.Printf("Tell client: %d bytes received\n", cnt)
		// (6)
		// (7)
		conn.Close()
		fmt.Printf("(%d) Disconnected!\n", connectionNum)
	}
}
