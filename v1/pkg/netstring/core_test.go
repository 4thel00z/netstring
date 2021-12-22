package netstring

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type testCase struct {
	Title       string
	Reader      *strings.Reader
	Expected    NetString
	ExpectedErr bool
}

func TestFromReader(t *testing.T) {
	testCases := []testCase{
		{
			Title:       "Happy flow",
			Reader:      strings.NewReader("100:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,"),
			Expected:    NetString("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ExpectedErr: false,
		},
	}
	for _, tc := range testCases {
		fmt.Println(tc.Title)
		output, err := FromReader(tc.Reader)
		if err != nil && !tc.ExpectedErr {
			t.Fail()
		}
		assert.Equal(t, tc.Expected, output)

	}

}
