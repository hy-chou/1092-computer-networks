package main

import (
	"bufio"
	"fmt"
	"net"
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
	fmt.Println("Connected!")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("Server replies: %s\n", scanner.Text())
	}
}
