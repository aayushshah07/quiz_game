package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to Quiz Game !!!! \n")

	// Opening CSV file
	f, err := os.Open("problem.csv")

	if err != nil {
		log.Fatal(err)

	}

	defer f.Close()

	// Creating CSV reader
	csvreader := csv.NewReader(f)

	for {
		rec, err := csvreader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(rec[0])
		fmt.Println("Answer the question \n")
		var ans = 0
		fmt.Scanf("%s", &ans)
		fmt.Println("Answer %s ", rec[1])

		if ans == rec[1] {
			fmt.Println("Correct answer \n")
		} else {
			fmt.Println("Wrong answer \n")
			break
		}
	}

}
