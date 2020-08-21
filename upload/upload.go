package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
)

const maxUploadSize = 2048 * 1024 // 2048 MB

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("./index.html", os.O_RDONLY, 0755)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		reader.WriteTo(w)
	})
	http.HandleFunc("/upload/traininglist", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			fmt.Printf("Could not parse multipart form: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, fileheader, err := r.FormFile("traininglist")
		if err != nil {
			http.Error(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()

		if fileheader.Size > maxUploadSize {
			fmt.Printf("File size (bytes): %v\n", fileheader.Size)
			http.Error(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for _, v := range records {
			fmt.Fprintln(w, v)
		}
	})
	fmt.Println("Listen on port :5000")
	http.ListenAndServe(":5000", nil)
}
