package go_boolean_calculator

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct {
	Input      []string
	BoolInputs []bool
}

func (c Calculator) Run() bool {

	fmt.Println("Input within the run function")
	fmt.Println(c.Input)
	fmt.Println("length")
	fmt.Println(len(c.Input))
	value, _ := strconv.ParseBool(c.Input[0])
	return value
}

func CreateCalculator(input string) Calculator {
	inputAsArray := strings.Split(input, ",")
	cal := Calculator{Input: inputAsArray}
	cal.ConvertSingleValues(cal.Input)
	return cal
}

func (c Calculator) ConvertSingleValues(input []string) {
	value := false

	for i := 0; i < len(input); i++ {
		fmt.Println(i, input[i])
		if strings.Contains(input[i], "NOT") {
			value = !strings.Contains(input[i], "TRUE")
		} else if strings.Contains(input[i], "AND") {
			value = !strings.Contains(input[i], "FALSE")
		} else if strings.Contains(input[i], "OR") {
			value = strings.Contains(input[i], "TRUE")
		} else {
			value = strings.Contains(input[i], "TRUE")
		}
		c.BoolInputs = append(c.BoolInputs, value)
	}
}
