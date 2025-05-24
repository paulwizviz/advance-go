package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCSVToData(t *testing.T) {

	testcases := []struct {
		name        string
		input       []string
		expected    Data
		expectedErr error
	}{
		{
			name:  "properly formatted input",
			input: []string{"1", "hello", "21-Mar-2021"},
			expected: Data{
				Index: 1,
				Value: "hello",
				Date:  time.Date(2021, time.March, 21, 0, 0, 0, 0, time.UTC),
			},
			expectedErr: nil,
		},
		{
			name:  "properly formatted day",
			input: []string{"1", "hello", "1-Mar-2021"},
			expected: Data{
				Index: 1,
				Value: "hello",
				Date:  time.Date(2021, time.March, 1, 0, 0, 0, 0, time.UTC),
			},
			expectedErr: nil,
		},
		{
			name:        "empty record",
			input:       []string{},
			expected:    Data{},
			expectedErr: ErrFldNum,
		},
		{
			name:        "invalid index field",
			input:       []string{"a", "hello", "21-Mar-2021"},
			expected:    Data{},
			expectedErr: ErrIndexFld,
		},
		{
			name:        "empty value field",
			input:       []string{"1", "", "21-Mar-2021"},
			expected:    Data{},
			expectedErr: ErrValueFld,
		},
		{
			name:        "invalid date month",
			input:       []string{"1", "hello", "21-March-2021"},
			expected:    Data{},
			expectedErr: ErrDateFld,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d-%s", i, tc.name), func(t *testing.T) {
			actual, err := csvRecToData(tc.input)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, err, tc.expectedErr)
				assert.Equal(t, tc.expected, actual)
				return
			}
			if assert.NoError(t, err) {
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
