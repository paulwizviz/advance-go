package csvutil

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCSV(t *testing.T) {
	testcases := []struct {
		name     string
		input    []byte
		expected struct {
			CSVRec
			Err error
		}
	}{
		{
			name: "correctly formatted",
			input: []byte(`index,value,date
1,"abc",1-Jan-2022`),
			expected: struct {
				CSVRec
				Err error
			}{
				CSVRec: CSVRec{
					Record: []string{
						"1",
						"abc",
						"1-Jan-2022",
					},
				},
				Err: nil,
			},
		},
		{
			name:  "empty content",
			input: []byte{},
			expected: struct {
				CSVRec
				Err error
			}{
				CSVRec: CSVRec{
					Record: nil,
				},
				Err: ErrCSV,
			},
		},
		{
			name: "mismatched field",
			input: []byte(`index,date
1,1.1,1-Jan-22`),
			expected: struct {
				CSVRec
				Err error
			}{
				CSVRec: CSVRec{
					Record: []string{"1", "1.1", "1-Jan-22"},
				},
				Err: ErrCSVRec,
			},
		},
		{
			name: "incompete data row",
			input: []byte(`index,value,date
1,1-Jan-2022`),
			expected: struct {
				CSVRec
				Err error
			}{
				CSVRec: CSVRec{
					Record: []string{
						"1",
						"1-Jan-2022",
					},
				},
				Err: ErrCSVRec,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d %s", i, tc.name), func(t *testing.T) {
			reader := bytes.NewReader(tc.input)
			records := ParseCSV(context.TODO(), reader)
			for rec := range records {
				if assert.True(t, errors.Is(rec.Err, tc.expected.Err), fmt.Sprintf("Case: %d Error: %v", i, rec.Err)) {
					assert.Equal(t, tc.expected.CSVRec.Record, rec.Record, fmt.Sprintf("Case: %d Value", i))
				}
			}
		})
	}
}
