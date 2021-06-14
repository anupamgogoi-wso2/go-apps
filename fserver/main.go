package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fileServer()
}
func fileServer() {
	log.Println("Starting file server...")
	err := http.ListenAndServe(":9000", http.HandlerFunc(fileHandler))
	if err != nil {
		log.Fatal("Server error:", err)
	}

}
func fileHandler(w http.ResponseWriter, r *http.Request) {
	found := find(r.RequestURI)
	path, _ := os.Getwd()
	if found {
		log.Println("Serving from: " + path)
		fs := http.FileServer(http.Dir("."))
		fs.ServeHTTP(w, r)
	} else {
		path := path + r.RequestURI
		isFile := checkIfFile(path)
		if isFile {
			log.Println("Downloading: " + path)
			w.Header().Set("Content-Disposition", "attachment; filename="+r.RequestURI[1:len(r.RequestURI)])
			w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		} else {
			log.Println("Serving from: " + path)
			fs := http.FileServer(http.Dir("."))
			fs.ServeHTTP(w, r)
		}

	}

}
func find(val string) (found bool) {
	uris := getExcludedURIs()
	for i := 0; i < len(uris); i++ {
		if uris[i] == val {
			return true
		}
	}
	return false
}
func getExcludedURIs() []string {
	return []string{"/", "/favicon.ico"}
}
func checkIfFile(path string) (isFile bool) {
	file, _ := os.Open(path)
	fileInfo, _ := file.Stat()
	return !fileInfo.IsDir()
}
