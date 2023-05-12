package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	fmt.Print("Hi")
	expression, err := govaluate.NewEvaluableExpression("foo > 0")

	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = -1

	result, err := expression.Evaluate(parameters)
	fmt.Println(result, err)
}
