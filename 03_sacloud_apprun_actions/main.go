package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 環境変数 PORT が指定されていなければ 8080 を使う
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Hello, World from Sakura Cloud AppRun!")
	})

	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	fmt.Println("test")
}
