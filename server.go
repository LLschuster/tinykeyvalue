package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"io"

	"github.com/gorilla/mux"
)

func keyvalueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("keyvalueHandler got method -> ", r.Method)
	responseHeaders := w.Header()
	key := mux.Vars(r)["key"]

	if r.Method == "PUT" {
		fmt.Println("content type ", r.Header.Get("Content-Type"))

		err := r.ParseMultipartForm(1024 * 1024 * 100)
		if err != nil {
			fmt.Println("Error parsing form: ", err)
			return
		}

		requestFile, err := r.MultipartForm.File["file"][0].Open()
		if err != nil {
			fmt.Println("Error getting file from request form: ", err)
			return
		}
		defer requestFile.Close()
		bodyContent, err := io.ReadAll(requestFile)

		if err != nil {
			log.Default().Printf("Error reading body: %v", err)
			responseHeaders.Add("Content-Type", "application/text")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error reading body"))
			return
		}

		file, err := os.OpenFile("volume", os.O_APPEND|os.O_CREATE|os.O_RDWR, 7777)
		if err != nil {
			log.Default().Printf("Error while opening volume: %v", err)
			responseHeaders.Add("Content-Type", "application/text")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error while opening volume"))
			return
		}
		defer file.Close()

		n, err := file.Write(bodyContent)
		if err != nil {
			log.Default().Printf("Error while writing new value to file: %v", err)
			responseHeaders.Add("Content-Type", "application/text")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error while writing new value to file"))
			return

		}

		fmt.Println("Wrote", n, "bytes to volume")
	}

	responseHeaders.Add("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("keyvalueHandler got key -> " + key + " got method -> " + r.Method))
}

func StartServer(port uint) {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", keyvalueHandler)

	log.Default().Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
