package structtag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type firstName struct {
	FirstName string `json:"first_name" sqlite:"first_name,TEXT"`
}

type fullName struct {
	FirstName string `json:"first_name" sqlite:"first_name,TEXT"`
	Surname   string `json:"surname" sqlite:"surname,TEXT"`
}

type namedIdentifier struct {
	ID   int `json:"id" sqlite:"id,INTEGER,PRIMARY_KEY"`
	Name fullName
}

func TestExtractTags(t *testing.T) {
	testcases := []struct {
		name     string
		input    any
		expected []Element
	}{
		{
			name:  "firstname struct",
			input: firstName{},
			expected: []Element{
				{FieldName: "FirstName", Tag: "first_name,TEXT"},
			},
		},
		{
			name:  "full name",
			input: fullName{},
			expected: []Element{
				{FieldName: "FirstName", Tag: "first_name,TEXT"},
				{FieldName: "Surname", Tag: "surname,TEXT"},
			},
		},
		{
			name:  "named identifier",
			input: namedIdentifier{},
			expected: []Element{
				{FieldName: "ID", Tag: "id,INTEGER,PRIMARY_KEY"},
				{FieldName: "Name", Tag: ""},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d-%s", i, tc.name), func(t *testing.T) {
			actual := ExtractPromoted("sqlite", tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
