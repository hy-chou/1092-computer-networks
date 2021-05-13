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
	portMin := 20001
	portMax := 21000
	port := (portMin + portMax) / 2
	bingo := false
	for !bingo {
		port = (portMin + portMax) / 2
		conn, errc := net.Dial("tcp", ip+":"+strconv.Itoa(port))
		check(errc)
		defer conn.Close()
		fmt.Println("Connected to port " + strconv.Itoa(port))
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("Server replies: %s\n", scanner.Text())
			if scanner.Text() == "bingo" {
				bingo = true
			} else if scanner.Text() == "low" {
				portMin = port + 1
			} else if scanner.Text() == "high" {
				portMax = port - 1
			}
		}
	}

	fmt.Println(">>> Bingo at port " + strconv.Itoa(port) + " <<<")
}
