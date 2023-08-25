package go_boolean_calculator

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_UpdateItems(t *testing.T) {
	tests := []struct {
		Input          string
		ExpectedOutput bool
	}{
		{
			Input:          "TRUE",
			ExpectedOutput: true,
		},
		{
			Input:          "FALSE",
			ExpectedOutput: false,
		},
	}
	for k, tc := range tests {
		t.Run("scenario "+strconv.Itoa(k+1), func(t *testing.T) {
			actualOutput := Calculator{tc.Input}.Run()
			assert.Equal(t, actualOutput, tc.ExpectedOutput)
		})
	}
}
