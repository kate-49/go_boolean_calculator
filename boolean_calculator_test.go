package go_boolean_calculator

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

//func Test_Simple_Inputs(t *testing.T) {
//	tests := []struct {
//		Input          string
//		ExpectedOutput bool
//	}{
//		{Input: "TRUE", ExpectedOutput: true},
//		{Input: "FALSE", ExpectedOutput: false},
//		{Input: "NOT TRUE", ExpectedOutput: false},
//		{Input: "NOT FALSE", ExpectedOutput: true},
//		{Input: "TRUE AND TRUE", ExpectedOutput: true},
//		{Input: "TRUE AND FALSE", ExpectedOutput: false},
//		{Input: "FALSE AND FALSE", ExpectedOutput: false},
//		{Input: "TRUE OR TRUE", ExpectedOutput: true},
//		{Input: "TRUE OR FALSE", ExpectedOutput: true},
//		{Input: "FALSE OR FALSE", ExpectedOutput: false},
//	}
//	for k, tc := range tests {
//		t.Run("scenario "+strconv.Itoa(k+1)+" given:"+tc.Input, func(t *testing.T) {
//			calculator := CreateCalculator(tc.Input)
//			assert.Equal(t, calculator.Run(), tc.ExpectedOutput)
//		})
//	}
//}

func Test_Longer_Inputs(t *testing.T) {
	tests := []struct {
		Input          string
		ExpectedOutput bool
	}{
		{Input: "TRUE OR TRUE OR TRUE AND FALSE", ExpectedOutput: true},
		{Input: "TRUE OR FALSE AND NOT FALSE", ExpectedOutput: true},
		//{Input: "NOT TRUE", ExpectedOutput: false},
		//{Input: "NOT FALSE", ExpectedOutput: true},
		//{Input: "TRUE AND TRUE", ExpectedOutput: true},
		//{Input: "TRUE AND FALSE", ExpectedOutput: false},
		//{Input: "FALSE AND FALSE", ExpectedOutput: false},
		//{Input: "TRUE OR TRUE", ExpectedOutput: true},
		//{Input: "TRUE OR FALSE", ExpectedOutput: true},
		//{Input: "FALSE OR FALSE", ExpectedOutput: false},
	}
	for k, tc := range tests {
		t.Run("scenario "+strconv.Itoa(k+1)+" given:"+tc.Input, func(t *testing.T) {
			calculator := CreateCalculator(tc.Input)
			assert.Equal(t, calculator.Run(), tc.ExpectedOutput)
		})
	}
}
