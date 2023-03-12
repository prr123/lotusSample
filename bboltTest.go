package main

import (
	"os"
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func main () {

	path := "boltTest.db"
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
  		fmt.Printf("error opening bolt: %v\n", err)
		os.Exit(-1)
	}
	defer db.Close()
	fmt.Println("success opening boltdb!")
}
