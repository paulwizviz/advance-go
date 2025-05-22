package main

import (
	"context"
	"fmt"
	"go-pattern/internal/csvutil"
	"log"
	"os"
	"path"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to get PWD: %v", err)
	}

	f, err := os.Open(path.Join(pwd, "testdata", "csv", "data1.csv"))
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer f.Close()

	data := csvutil.ParseCSV(context.TODO(), f)
	for d := range data {
		fmt.Printf("Line No: %d Record: %v Error: %v\n", d.Line, d.Record, d.Err)
	}

}
