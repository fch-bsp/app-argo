package main

import "net/http"

// Expondo menssagem na porta 8080 (Hello BSPCLOUD)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Mateus viado V6<h1>"))
	})
	http.ListenAndServe(":8080", nil)

}
