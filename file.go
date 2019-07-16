package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// open the test file
	csvfile, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln("sorry could not open the file", err)
	}

	// pase the file
	r := csv.NewReader(csvfile)
	fmt.Println("Outpuing the Age of all the users in the CSV file")
	for {
		//reading the file
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//		fmt.Printf("this is the file", record[0], record[1], record[3])

		fmt.Printf("age %s\n", record[2])

	}
}
