package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
	_ "github.com/ymotongpoo/go-container-bazel-sample/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatalf("Error on creating statik file system: %v", err)
	}
	fmt.Println("Go container sample")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	if err := http.ListenAndServe("0.0.0.0:8888", nil); err != nil {
		log.Fatalf("Error on serving HTTP: %v", err)
	}
}
