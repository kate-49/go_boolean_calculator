package go_boolean_calculator

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct {
	Input         []string
	Input2        []string
	BoolFirstPass []bool
	BoolInputs    []bool
	Text          string
}

func CreateCalculator(input string) Calculator {
	var stringWithinParams string
	var restOfString string
	cal := Calculator{Text: input}

	if strings.Contains(input, "(") {
		indexOfFirstParenthesis := strings.Index(input, "(")
		indexOfLastParenthesis := strings.Index(input, ")")
		stringWithinParams = input[indexOfFirstParenthesis+1 : indexOfLastParenthesis]
		restOfString = strings.ReplaceAll(input, input[indexOfFirstParenthesis:indexOfLastParenthesis+1], "")
		fmt.Println("ros2")
		fmt.Println(restOfString)
		cal.Input2 = strings.Split(stringWithinParams, " ")
		restOfString = strings.Trim(restOfString, " ")
		cal.Input = strings.Split(restOfString, " ")
	} else {
		cal.Input = strings.Split(input, " ")
	}
	return cal
}

func (c *Calculator) Run() bool {
	fmt.Println("input2")
	fmt.Println(c.Input2)
	fmt.Println("input1")
	fmt.Println(c.Input)

	searchAbleElements := []string{"NOT", "OR", "AND"}
	if len(c.Input2) > 0 {
		c.CalculateSoloValues(c.Input2, true)
		for i := 0; i < len(searchAbleElements); i++ {
			c.CalculateIndividualElements(searchAbleElements[i], c.Input2, true)
		}
		fmt.Println("bool first pass")
		fmt.Println(c.BoolFirstPass)

		//find out where original parenthesis was and convert this back to a string and put there
		indexElement1 := strings.Index(c.Text, c.Input2[0])
		indexElement2 := strings.Index(c.Text, c.Input[0])
		fmt.Println("indexElement1")
		fmt.Println(indexElement1)
		fmt.Println("indexElement2")
		fmt.Println(indexElement2)

		stringToAppend := ""

		for i := 0; i < len(c.BoolFirstPass); i++ {
			stringToAppend = strings.ToUpper(strconv.FormatBool(c.BoolFirstPass[i]))
		}

		fmt.Println("stringToAppend")
		fmt.Println(stringToAppend)

		if indexElement1 > indexElement2 {
			fmt.Println("1 bigger than 2")
			fmt.Println(c.Input)
			fmt.Println(stringToAppend)
			c.Input = append(c.Input, stringToAppend)
		} else {
			fmt.Println("2 bigger than 1")
			fmt.Println(stringToAppend)
			fmt.Println(c.Input)
			str := stringToAppend + " " + strings.Join(c.Input, " ")
			fmt.Println("str")
			fmt.Println(str)
			c.Input = strings.Split(str, " ")
		}
	}
	fmt.Println("c.Input2")
	fmt.Println(c.Input)
	c.CalculateSoloValues(c.Input, false)
	fmt.Println("after solo values")
	fmt.Println(c.BoolInputs)
	for i := 0; i < len(searchAbleElements); i++ {
		c.CalculateIndividualElements(searchAbleElements[i], c.Input, false)
	}
	fmt.Println("bool inputs after second run")
	fmt.Println(c.BoolInputs)
	return c.GetFinalScore()
}

func (c *Calculator) CalculateIndividualElements(searchString string, input []string, parenthesis bool) {
	for i := 0; i < len(input); i++ {
		if input[i] == searchString {
			switch searchString {
			case "NOT":
				fmt.Println("not")
				//if element not null after not
				if input[i+1] != "" {
					value, _ := strconv.ParseBool(input[i+1])
					if parenthesis {
						c.BoolFirstPass = append(c.BoolInputs, !value)
					} else {
						c.BoolInputs = append(c.BoolInputs, !value)
					}
				}
			case "AND":
				value := false
				if i+1 < len(input) {
					fmt.Println(input[i-1] == input[i+1])
					if input[i-1] == input[i+1] {
						value, _ = strconv.ParseBool(input[i+1])
					}
					if parenthesis {
						c.BoolFirstPass = append(c.BoolInputs, value)
					} else {
						c.BoolInputs = append(c.BoolInputs, value)
					}
				}
			case "OR":
				value := false
				fmt.Println("or")
				if i+1 < len(input) {
					fmt.Println("length ok")
					if (input[i-1] == "TRUE") || (input[i+1] == "TRUE") {
						value = true
					}
					fmt.Println(value)
					if parenthesis {
						c.BoolFirstPass = append(c.BoolInputs, value)
					} else {
						c.BoolInputs = append(c.BoolInputs, value)
					}
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
		fmt.Println(i)
		fmt.Println(input[i])
		if i > 0 {
			fmt.Println("a")
			if (input[i-1] != "NOT") && (input[i-1] != "AND") && (input[i-1] != "OR") {
				fmt.Println("b")

				if (input[i] == "TRUE") || (input[i] == "FALSE") {
					fmt.Println("c")

					if i+1 < len(input) {
						fmt.Println("d")

						if (input[i+1] != "AND") && (input[i+1] != "OR") {
							fmt.Println("e c")

							c.ConvertToBoolAndAppend(input[i], parenthesis)
						}
					} else {
						fmt.Println("f c")
						fmt.Println("i - 1")
						fmt.Println(input[i-1])
						if (input[i-1] != "NOT") && (input[i-1] != "AND") && (input[i-1] != "OR") {
							c.ConvertToBoolAndAppend(input[i], parenthesis)
						}
					}
				}
			}
		} else {
			fmt.Println("1b")
			if (input[i] == "TRUE") || (input[i] == "FALSE") {
				fmt.Println("2")

				if i+1 < len(input) {
					fmt.Println("3")

					if (input[i+1] != "AND") && (input[i+1] != "OR") {
						fmt.Println("4 c")

						c.ConvertToBoolAndAppend(input[i], parenthesis)
					}
				} else {
					fmt.Println("5 c")

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
