package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {

	// open browser with
	switch runtime.GOOS {
	case "windows":
		exec.Command("cmd", "/c", "start", "http://localhost:8080").Start()
	case "darwin":
		exec.Command("open", "http://localhost:8080").Start()
	default:
		exec.Command("xdg-open", "http://localhost:8080").Start()
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
