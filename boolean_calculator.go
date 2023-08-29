package go_boolean_calculator

import (
	"strconv"
	"strings"
)

type Calculator struct {
	Input      []string
	BoolInputs []bool
}

func CreateCalculator(input string) Calculator {
	inputString1 := string
	if strings.Contains(input, "(") {
		indexOfFirstParenthesis := strings.Split(input, "(")
		indexOfLastParenthesis := strings.Split(input, ")")
		inputString1 = (input[indexOfFirstParenthesis+1], [indexOfFirstParenthesis])
	}

	//const inputStringForParenthesis = inputString.slice([inputString.indexOf("(")+1], [inputString.indexOf(")")])
	//const restOfInputString = inputString.replace(inputStringForParenthesis, "");
	//return [inputStringForParenthesis, restOfInputString.replace("()", "")]

	inputAsArray := strings.Split(input, " ")
	cal := Calculator{Input: inputAsArray}
	return cal
}

func (c *Calculator) Run() bool {

	searchAbleElements := []string{"NOT", "OR", "AND"}
	c.CalculateSoloValues()
	for i := 0; i < len(searchAbleElements); i++ {
		c.CalculateIndividualElements(searchAbleElements[i])
	}
	return c.GetFinalScore()
}

func (c *Calculator) CalculateIndividualElements(searchString string) {
	for i := 0; i < len(c.Input); i++ {
		if c.Input[i] == searchString {
			switch searchString {
			case "NOT":
				value, _ := strconv.ParseBool(c.Input[i+1])
				c.BoolInputs = append(c.BoolInputs, !value)
			case "AND":
				value := false
				if c.Input[i-1] == c.Input[i+1] {
					value, _ = strconv.ParseBool(c.Input[i+1])
				}
				c.BoolInputs = append(c.BoolInputs, value)
			case "OR":
				value := false
				if (c.Input[i-1] == "TRUE") || (c.Input[i+1] == "TRUE") {
					value = true
				}
				c.BoolInputs = append(c.BoolInputs, value)
			}
		}
	}
}

func (c *Calculator) ConvertToBoolAndAppend(element string) {
	value, _ := strconv.ParseBool(element)
	c.BoolInputs = append(c.BoolInputs, value)
}

func (c *Calculator) CalculateSoloValues() {
	for i := 0; i < len(c.Input); i++ {
		if i > 0 {
			if c.Input[i-1] != "NOT" {
				if (c.Input[i] == "TRUE") || (c.Input[i] == "FALSE") {
					c.ConvertToBoolAndAppend(c.Input[i])
				}
			}
		} else {
			if (c.Input[i] == "TRUE") || (c.Input[i] == "FALSE") {
				if i+1 < len(c.Input) {
					if (c.Input[i+1] != "AND") && (c.Input[i+1] != "OR") {
						c.ConvertToBoolAndAppend(c.Input[i])
					}
				} else {
					c.ConvertToBoolAndAppend(c.Input[i])
				}
			}
		}
	}
}

func (c *Calculator) GetFinalScore() bool {
	numberOfTrue := 0
	numberOfFalse := 0
	if len(c.BoolInputs) == 1 {
		return c.BoolInputs[0]
	} else {
		for i := 0; i < len(c.BoolInputs); i++ {
			if c.BoolInputs[i] == true {
				numberOfTrue++
			}
			if c.BoolInputs[i] == false {
				numberOfFalse++
			}
		}
		return numberOfTrue >= numberOfFalse
	}
}
