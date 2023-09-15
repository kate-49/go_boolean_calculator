package go_boolean_calculator

import (
	"strconv"
	"strings"
)

type Calculator struct {
	Input             []string
	InputWithinParams []string
	BoolFirstPass     []bool
	BoolInputs        []bool
	Text              string
	IndexElement1     int
	IndexElement2     int
}

func CreateCalculator(input string) Calculator {
	cal := Calculator{Text: input}

	if strings.Contains(input, "(") {
		cal.CalculateValuesForParenthesisSetup(input)
	} else {
		cal.Input = strings.Split(input, " ")
	}
	return cal
}

func (c *Calculator) CalculateValuesForParenthesisSetup(input string) {
	var stringWithinParams string
	var restOfString string
	indexOfFirstParenthesis := strings.Index(input, "(")
	indexOfLastParenthesis := strings.Index(input, ")")
	stringWithinParams = input[indexOfFirstParenthesis+1 : indexOfLastParenthesis]
	restOfString = strings.ReplaceAll(input, input[indexOfFirstParenthesis:indexOfLastParenthesis+1], "")
	c.InputWithinParams = strings.Split(stringWithinParams, " ")
	restOfString = strings.Trim(restOfString, " ")
	c.Input = strings.Split(restOfString, " ")

	c.IndexElement1 = strings.Index(c.Text, c.InputWithinParams[0])
	c.IndexElement2 = strings.Index(c.Text, c.Input[0])
}

func (c *Calculator) CalculateForArray(input []string, parenthesis bool) {
	searchAbleElements := []string{"NOT", "OR", "AND"}

	c.CalculateSoloValues(input, parenthesis)
	for i := 0; i < len(searchAbleElements); i++ {
		c.CalculateIndividualElements(searchAbleElements[i], input, parenthesis)
	}
}

func (c *Calculator) Run() bool {
	if len(c.InputWithinParams) > 0 {
		c.CalculateForArray(c.InputWithinParams, true)
		stringToAppend := ""
		for i := 0; i < len(c.BoolFirstPass); i++ {
			stringToAppend = strings.ToUpper(strconv.FormatBool(c.BoolFirstPass[i]))
		}

		//append the original sub array back into the string based on order
		if c.IndexElement1 > c.IndexElement2 {
			c.Input = append(c.Input, stringToAppend)
		} else {
			str := stringToAppend + " " + strings.Join(c.Input, " ")
			c.Input = strings.Split(str, " ")
		}
	}
	c.CalculateForArray(c.Input, false)
	return c.GetFinalScore()
}

func (c *Calculator) Append(value, parenthesis bool) {
	if parenthesis {
		c.BoolFirstPass = append(c.BoolInputs, value)
	} else {
		c.BoolInputs = append(c.BoolInputs, value)
	}
}

func (c *Calculator) CalculateIndividualElements(searchString string, input []string, parenthesis bool) {
	value := false
	for i := 0; i < len(input); i++ {
		if input[i] == searchString {
			switch searchString {
			case "NOT":
				if input[i+1] != "" {
					value, _ := strconv.ParseBool(input[i+1])
					c.Append(!value, parenthesis)
				}
			case "AND":
				if i+1 < len(input) {
					if input[i-1] == input[i+1] {
						value, _ = strconv.ParseBool(input[i+1])
					}
					c.Append(value, parenthesis)

				}
			case "OR":
				if i+1 < len(input) {
					if (input[i-1] == "TRUE") || (input[i+1] == "TRUE") {
						value = true
					}
					c.Append(value, parenthesis)
				}
			}
		}
	}
}

func (c *Calculator) ConvertToBoolAndAppend(element string, parenthesis bool) {
	value, _ := strconv.ParseBool(element)
	if parenthesis {
		c.BoolFirstPass = append(c.BoolInputs, value)
	} else {
		c.BoolInputs = append(c.BoolInputs, value)
	}
}

func (c *Calculator) CalculateSoloValues(input []string, parenthesis bool) {
	for i := 0; i < len(input); i++ {
		if i > 0 {
			if (input[i-1] != "NOT") && (input[i-1] != "AND") && (input[i-1] != "OR") {
				if (input[i] == "TRUE") || (input[i] == "FALSE") {
					if i+1 < len(input) {
						if (input[i+1] != "AND") && (input[i+1] != "OR") {
							c.ConvertToBoolAndAppend(input[i], parenthesis)
						}
					} else {
						c.ConvertToBoolAndAppend(input[i], parenthesis)
					}
				}
			}
		} else {
			if (input[i] == "TRUE") || (input[i] == "FALSE") {
				if i+1 < len(input) {
					if (input[i+1] != "AND") && (input[i+1] != "OR") {
						c.ConvertToBoolAndAppend(input[i], parenthesis)
					}
				} else {
					c.ConvertToBoolAndAppend(input[i], parenthesis)
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
