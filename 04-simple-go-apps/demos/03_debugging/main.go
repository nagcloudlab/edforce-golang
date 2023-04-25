package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("menu.txt")

	io.Copy(w, f)
}
