package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header()
	w.Write([]byte("hello"))
}
func year(w http.ResponseWriter, r *http.Request) {
	w.Header()
	w.Write([]byte("year"))
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/year", year)
	http.ListenAndServe("0.0.0.0:8089", nil)
}
