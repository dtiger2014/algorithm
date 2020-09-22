package main

// https://www.jianshu.com/p/be3d9cdc680b

import(
	"net/http"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.0:8000", nil)
}