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

	for port := 20001; port <= 20100; port++ {
		conn, errc := net.Dial("tcp", "localhost:"+strconv.Itoa(port))
		check(errc)

		writer := bufio.NewWriter(conn)
		_, errw := writer.WriteString("GET / HTTP/1.1\nHost:\n\n")
		check(errw)
		writer.Flush()

		scanner := bufio.NewScanner(conn)
		scanner.Scan()
		fmt.Printf(strconv.Itoa(port)+": %s\n", scanner.Text())
		conn.Close()
	}
}
