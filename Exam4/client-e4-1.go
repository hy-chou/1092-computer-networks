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
	fmt.Println("Launching client...")

	conn, errc := net.Dial("tcp", "localhost:20001")
	check(errc)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString("GET / HTTP/1.1\nHost:\n\n")
	check(errw)
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	fmt.Printf(strconv.Itoa(20001)+": %s\n", scanner.Text())
}
