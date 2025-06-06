package csvutil

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCSVC(t *testing.T) {
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
			records := ParseCSVC(context.TODO(), reader)
			for rec := range records {
				if assert.True(t, errors.Is(rec.Err, tc.expected.Err), fmt.Sprintf("Case: %d Error: %v", i, rec.Err)) {
					assert.Equal(t, tc.expected.CSVRec.Record, rec.Record, fmt.Sprintf("Case: %d Value", i))
				}
			}
		})
	}
}

func TestParseCSV(t *testing.T) {
	testcases := []struct {
		name        string
		input       []byte
		expected    []CSVRec
		expectedErr error
	}{
		{
			name: "properly formatted file",
			input: []byte(`field1,field2,field3
1,2,3
4,5,6`),
			expected: []CSVRec{
				{
					Line:   2,
					Record: []string{"1", "2", "3"},
					Err:    nil,
				},
				{
					Line:   3,
					Record: []string{"4", "5", "6"},
					Err:    nil,
				},
			},
		},
		{
			name: "properly formatted file",
			input: []byte(`field1,field2
1,2,3`),
			expected: []CSVRec{
				{
					Line:   2,
					Record: []string{"1", "2", "3"},
				},
			},
			expectedErr: ErrCSVRec,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d-%s", i, tc.name), func(t *testing.T) {
			r := bytes.NewReader(tc.input)
			actual := ParseCSV(context.TODO(), r)
			if tc.expectedErr != nil {
				for i, a := range actual {
					assert.ErrorIs(t, a.Err, tc.expectedErr)
					assert.Equal(t, tc.expected[i].Record, a.Record)
				}
				return
			}
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestCountLine(t *testing.T) {
	testcases := []struct {
		name     string
		input    []byte
		expected uint
	}{
		{
			name: "correct count",
			input: []byte(`field1,field2,field3
1,1,1
2,3,1
4,1,1`),
			expected: uint(4),
		},
		{
			name: "missing header",
			input: []byte(`
1,2,3
1,2,3			
`),
			expected: uint(2),
		},
		{
			name: "unbalanced row",
			input: []byte(`field1,field2,field3
1,3
1,2,3`),
			expected: uint(2),
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d-%s", i, tc.name), func(t *testing.T) {
			r := bytes.NewReader(tc.input)
			actual := CountLines(r)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
