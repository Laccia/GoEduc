package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Введите 1-ю сторону")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Введите 2-ю сторону")
	fmt.Scanf("%d\n", &b)
	c := a * b
	fmt.Println("Площадь = ", c)
}
