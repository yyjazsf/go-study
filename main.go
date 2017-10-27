package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// https://www.ctolib.com/docs-go-web-programming-c-03-2.html
func app(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Printf("%s %s", k, strings.Join(v, ""))
	}
	fmt.Fprintf(w, r.RemoteAddr) //response
}

func main() {
	http.HandleFunc("/", app)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("error:", err)
	}
}
