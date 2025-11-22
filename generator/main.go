package main

import "fmt"

type Number interface {
	int | float64 | string
}

func AddNumber[T Number](t T) T {
	return t + t
}

func main() {
	stringNum := AddNumber("1")
	fmt.Printf("%s %T\n", stringNum, stringNum)
	intNum := AddNumber(1)
	floatNum := AddNumber(1.0)
	fmt.Printf("%d %T\n", intNum, intNum)
	fmt.Printf("%f %T\n", floatNum, floatNum)
}
