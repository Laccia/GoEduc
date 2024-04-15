package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	fmt.Println("площадь круга = ")
	fmt.Scanln(&a)
	c := 2 * math.Sqrt(math.Pi*a)
	fmt.Println("Длина окружности = ", c)
	b = c / math.Pi
	fmt.Println("Диаметр = ", b)
}
