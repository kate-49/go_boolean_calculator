package go_boolean_calculator

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_E2E_Simple_Inputs(t *testing.T) {
	tests := []struct {
		Input          string
		ExpectedOutput bool
	}{
		{Input: "TRUE", ExpectedOutput: true},
		{Input: "FALSE", ExpectedOutput: false},
		{Input: "NOT TRUE", ExpectedOutput: false},
		{Input: "NOT FALSE", ExpectedOutput: true},
		{Input: "TRUE AND TRUE", ExpectedOutput: true},
		{Input: "TRUE AND FALSE", ExpectedOutput: false},
		{Input: "FALSE AND FALSE", ExpectedOutput: false},
		{Input: "TRUE OR TRUE", ExpectedOutput: true},
		{Input: "TRUE OR FALSE", ExpectedOutput: true},
		{Input: "FALSE OR FALSE", ExpectedOutput: false},
	}
	for k, tc := range tests {
		t.Run("scenario "+strconv.Itoa(k+1)+" given:"+tc.Input, func(t *testing.T) {
			calculator := CreateCalculator(tc.Input)
			assert.Equal(t, tc.ExpectedOutput, calculator.Run())
		})
	}
}

func Test_E2E_Longer_Inputs(t *testing.T) {
	tests := []struct {
		Input          string
		ExpectedOutput bool
	}{
		{Input: "TRUE OR TRUE OR TRUE AND FALSE", ExpectedOutput: true},
		{Input: "TRUE OR FALSE AND NOT FALSE", ExpectedOutput: true},
		{Input: "(TRUE TRUE OR FALSE) AND TRUE", ExpectedOutput: true},
		{Input: "(NOT FALSE) OR FALSE", ExpectedOutput: true},
		{Input: "(TRUE AND FALSE) AND TRUE", ExpectedOutput: false},
		{Input: "(TRUE AND TRUE)", ExpectedOutput: true},
		{Input: "(TRUE OR TRUE OR TRUE) AND FALSE", ExpectedOutput: false},
		{Input: "NOT (TRUE AND TRUE)", ExpectedOutput: false},
		{Input: "(NOT TRUE AND NOT FALSE) AND NOT TRUE", ExpectedOutput: false},
		{Input: "(TRUE OR FALSE) AND NOT TRUE", ExpectedOutput: false},
	}
	for k, tc := range tests {
		t.Run("scenario "+strconv.Itoa(k+1)+" given:"+tc.Input, func(t *testing.T) {
			calculator := CreateCalculator(tc.Input)
			assert.Equal(t, calculator.Run(), tc.ExpectedOutput)
		})
	}
}

func Test_Create_Calculator(t *testing.T) {
	tests := []struct {
		Input          string
		ExpectedOutput Calculator
	}{
		{Input: "TRUE OR TRUE OR TRUE AND FALSE", ExpectedOutput: Calculator{
			Input: []string{"TRUE", "OR", "TRUE", "OR", "TRUE", "AND", "FALSE"},
			Text:  "TRUE OR TRUE OR TRUE AND FALSE",
		}},
		{Input: "(NOT FALSE) OR FALSE", ExpectedOutput: Calculator{
			Input:             []string{"OR", "FALSE"},
			InputWithinParams: []string{"NOT", "FALSE"},
			Text:              "(NOT FALSE) OR FALSE",
			IndexElement1:     1,
			IndexElement2:     12,
		}},
		{Input: "NOT (TRUE AND TRUE)", ExpectedOutput: Calculator{
			Input:             []string{"NOT"},
			InputWithinParams: []string{"TRUE", "AND", "TRUE"},
			Text:              "NOT (TRUE AND TRUE)",
			IndexElement1:     5,
			IndexElement2:     0,
		}},
	}
	for k, tc := range tests {
		t.Run("scenario "+strconv.Itoa(k+1)+" given:"+tc.Input, func(t *testing.T) {
			assert.Equal(t, CreateCalculator(tc.Input), tc.ExpectedOutput)
		})
	}
}

func Test_Calculate_For_Array(t *testing.T) {
	tests := []struct {
		Input         []string
		Parenthesis   bool
		BoolFirstPass []bool
		BoolInput     []bool
	}{
		{
			Input:         []string{"TRUE", "OR", "TRUE", "OR", "TRUE", "AND", "FALSE"},
			Parenthesis:   true,
			BoolFirstPass: []bool{false},
			BoolInput:     nil,
		},
		{
			Input:         []string{"TRUE", "OR", "TRUE", "OR", "TRUE", "AND", "FALSE"},
			Parenthesis:   false,
			BoolFirstPass: nil,
			BoolInput:     []bool{false},
		},
		{
			Input:         []string{"NOT", "TRUE", "AND", "TRUE"},
			Parenthesis:   true,
			BoolFirstPass: []bool{true},
			BoolInput:     nil,
		},
	}
	for k, tc := range tests {
		t.Run("scenario "+strconv.Itoa(k+1), func(t *testing.T) {
			calc := CreateCalculator("")
			calc.CalculateForArray(tc.Input, tc.Parenthesis)
			outputCalc := Calculator{
				Input:         []string{""},
				BoolFirstPass: tc.BoolFirstPass,
				BoolInputs:    tc.BoolInput,
			}
			assert.Equal(t, calc, outputCalc)
		})
	}
}
