package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

var ignoreURIs = []string{"/", "/favicon.ico"}

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
	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	isRootOrFavicon := checkRootOrFavicon(r.RequestURI)
	path, _ := os.Getwd()
	if isRootOrFavicon {
		log.Println("Requesting " + path + " from " + clientIP)
		fs := http.FileServer(http.Dir("."))
		fs.ServeHTTP(w, r)
	} else {
		path := path + r.RequestURI
		isFile := checkIfFile(path)
		if isFile {
			fileName := r.RequestURI[1:len(r.RequestURI)]
			log.Println("Downloading " + fileName + " from " + clientIP)
			w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
			w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
			http.ServeFile(w, r, path)

		} else {
			log.Println("Requesting " + path + " from " + clientIP)
			fs := http.FileServer(http.Dir("."))
			fs.ServeHTTP(w, r)
		}

	}

}
func checkRootOrFavicon(val string) (found bool) {
	for i := 0; i < len(ignoreURIs); i++ {
		if ignoreURIs[i] == val {
			return true
		}
	}
	return false
}

func checkIfFile(path string) (isFile bool) {
	file, _ := os.Open(path)
	fileInfo, _ := file.Stat()
	return !fileInfo.IsDir()
}
