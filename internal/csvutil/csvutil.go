// Package csvutil provides utility types and functions to simplify working with CSV data.
//
// This package includes:
//   - A function to read a CSV file in a goroutine and channel data.
//   - A function to help count number of lines in a CSV file
package csvutil

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
)

var (
	ErrCSV    = errors.New("unspecified")
	ErrCSVRec = errors.New("invalid record")
)

// CSVRec represents a record in a CSV file
type CSVRec struct {
	Record []string
	Line   uint
	Err    error
}

// ParseCSVC parse a CSV file in a Goroutine and returns a
// csv record in a channel
func ParseCSVC(ctx context.Context, r io.Reader) chan CSVRec {
	c := make(chan CSVRec)
	go func(ch chan CSVRec) {
		defer close(ch)
		csvr := csv.NewReader(r)
		header, err := csvr.Read()
		ln := uint(1)
		if err != nil {
			ch <- CSVRec{
				Record: header,
				Line:   ln,
				Err:    fmt.Errorf("%w-%s", ErrCSV, err.Error()),
			}
			return
		}
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
				ln++
				rec, err := csvr.Read()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break loop
					}
					ch <- CSVRec{
						Record: rec,
						Line:   ln,
						Err:    fmt.Errorf("%w-%s", ErrCSVRec, err.Error()),
					}
					continue loop
				}
				ch <- CSVRec{
					Record: rec,
					Line:   ln,
					Err:    nil,
				}
			}
		}
	}(c)
	return c
}

// ParseCSV reads a file extract a record and aggregate them
// in a slice
func ParseCSV(ctx context.Context, r io.Reader) []CSVRec {
	recs := []CSVRec{}
	csvr := csv.NewReader(r)
	ln := uint(1)
	header, err := csvr.Read()
	if err != nil {
		rec := CSVRec{
			Record: header,
			Line:   ln,
			Err:    fmt.Errorf("%w-%v", ErrCSV, err),
		}
		recs = append(recs, rec)
		return recs
	}
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			ln++
			r, err := csvr.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break loop
				}
				rec := CSVRec{
					Line:   ln,
					Record: r,
					Err:    fmt.Errorf("%w-%v", ErrCSVRec, err),
				}
				recs = append(recs, rec)
				continue loop
			}
			rec := CSVRec{
				Line:   ln,
				Record: r,
				Err:    nil,
			}
			recs = append(recs, rec)
		}
	}
	return recs
}

func CountLines(r io.Reader) uint {
	reader := csv.NewReader(r)
	lc := uint(0)
loop:
	for {
		_, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break loop
			}
			continue
		}
		lc++
	}
	return lc
}
