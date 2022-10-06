package main

import "net/http"

// Expondo menssagem na porta 8080 (Hello BSPCLOUD)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Flow de CI/CD V3<h1>"))
	})
	http.ListenAndServe(":8080", nil)

}
