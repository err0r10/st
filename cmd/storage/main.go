package main

import (
	"fmt"
	"log"

	"github.com/err0r10/st/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("text test"))
	if err != nil {
		log.Fatal(err)
	}
	foundFile, err := st.GetByID(file.ID)
	fmt.Println("File >> ", foundFile)
}
