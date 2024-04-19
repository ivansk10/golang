package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalc(n1, n2 int8) (int8, int8, int8) {

	sum := n1 + n2
	sub := n1 - n2
	div := n1 / n2
	return sum, sub, div
}

func main() {
	sum := sum(10, 20)
	fmt.Println(sum)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Function F")
	fmt.Println(result)

	resultMathCalcSum, resultMathCalcSub, resultMathCalcDiv := mathCalc(10, 2)
	fmt.Println(resultMathCalcSum)
	fmt.Println(resultMathCalcSub)
	fmt.Println(resultMathCalcDiv)

	resultMathCalcSum1, _, _ := mathCalc(10, 2)
	fmt.Println(resultMathCalcSum1)
}
