package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	dir, e := os.Getwd()
	if e != nil {
		fmt.Println(e)
		return
	}
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	fmt.Println("dir = ", dir)
	fmt.Println("listened on http://localhost:8080")
	e = http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println(e)
		return
	}

}
