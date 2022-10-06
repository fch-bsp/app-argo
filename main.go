package main

import "net/http"

// Expondo menssagem na porta 8080 (Hello BSPCLOUD)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Fluxo de CI/CD Versão-02<h1>"))
	})
	http.ListenAndServe(":8080", nil)

}
