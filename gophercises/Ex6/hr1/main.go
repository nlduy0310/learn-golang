package main

import (
	"fmt"
	"hr1/camel"
	"hr1/ceasar"
)

func main() {
	fmt.Println(camel.CountLettersInCamelCase("thisIsACamelCaseString"))
	fmt.Println(ceasar.Encode("There's-a-starman-waiting-in-the-sky", 3))
}
