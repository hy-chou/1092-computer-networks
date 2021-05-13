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

	ip := "localhost"
	port := 20000
	bingo := false
	for !bingo {
		port += 1
		conn, errc := net.Dial("tcp", ip+":"+strconv.Itoa(port))
		check(errc)
		defer conn.Close()
		fmt.Println("Connected to port " + strconv.Itoa(port))
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("Server replies: %s\n", scanner.Text())
			if scanner.Text() == "bingo" {
				bingo = true
			}
		}
	}

	fmt.Println(">>> Bingo at port " + strconv.Itoa(port))
}
