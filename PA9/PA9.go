package main

import (
	"fmt"
	"net/http"
	"os"
)

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

	http.ListenAndServeTLS(":12999", "server.cer", "server.key", hc)
}
