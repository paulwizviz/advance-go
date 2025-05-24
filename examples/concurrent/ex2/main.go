package main

import (
	"context"
	"errors"
	"fmt"
	"go-pattern/internal/csvutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var (
	ErrInvalidFormat = errors.New("invalid month format")
)

var (
	ErrFldNum   = errors.New("insufficient number of fields")
	ErrIndexFld = errors.New("index field error")
	ErrValueFld = errors.New("value field error")
	ErrDateFld  = errors.New("date field error")
)

type Data struct {
	Index int
	Value string
	Date  time.Time
}

func csvRecToData(recs []string) (Data, error) {

	if len(recs) != 3 {
		return Data{}, ErrFldNum
	}

	index, err := strconv.Atoi(recs[0])
	if err != nil {
		return Data{}, fmt.Errorf("%w:%v", ErrIndexFld, err)
	}

	if recs[1] == "" {
		return Data{}, fmt.Errorf("%w:%v", ErrValueFld, errors.New("empty value"))
	}

	dateLayout := "2-Jan-2006"
	tm, err := time.Parse(dateLayout, recs[2])
	if err != nil {
		return Data{}, fmt.Errorf("%w:%v", ErrDateFld, err)
	}
	return Data{
		Index: index,
		Value: recs[1],
		Date:  tm,
	}, nil
}

// processCSV is a worker to process csv record and pipe to result
func processCSV(jobs <-chan csvutil.CSVRec, result chan<- Data) {
	for j := range jobs {
		if j.Err != nil {
			continue
		}
		data, err := csvRecToData(j.Record)
		if err != nil {
			continue
		}
		result <- data
	}
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(filepath.Join(pwd, "testdata", "csv", "data1.csv"))
	if err != nil {
		log.Fatal(err)
	}

	jobs := csvutil.ParseCSVC(context.TODO(), f)
	result := make(chan Data)

	// This spins up 3 workers
	var wg sync.WaitGroup
	for range 3 {
		wg.Add(1)
		go func() {
			processCSV(jobs, result)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for r := range result {
		fmt.Println(r)
	}

}
