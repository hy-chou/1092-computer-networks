package main

import (
	"fmt"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.String()[1:]
	if _, errf := os.Stat(filename); os.IsNotExist(errf) {
		fmt.Fprintln(w, "File not found")
	} else {
		http.ServeFile(w, r, filename)
	}
}

func main() {
	fmt.Println("Launching server...")
	hc := http.HandlerFunc(handleConnection)
	http.ListenAndServe(":12999", hc)
}
