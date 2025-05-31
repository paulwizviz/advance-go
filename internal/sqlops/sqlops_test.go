package sqlops

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
		expected []StructTag
	}{
		{
			name:  "firstname struct",
			input: firstName{},
			expected: []StructTag{
				{FieldName: "FirstName", Tag: "first_name,TEXT"},
			},
		},
		{
			name:  "full name",
			input: fullName{},
			expected: []StructTag{
				{FieldName: "FirstName", Tag: "first_name,TEXT"},
				{FieldName: "Surname", Tag: "surname,TEXT"},
			},
		},
		{
			name:  "named identifier",
			input: namedIdentifier{},
			expected: []StructTag{
				{FieldName: "ID", Tag: "id,INTEGER,PRIMARY_KEY"},
				{FieldName: "Name", Tag: ""},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d-%s", i, tc.name), func(t *testing.T) {
			actual := ExtractTags("sqlite", tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
