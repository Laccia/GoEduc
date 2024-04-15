package main

import (
	"fmt"
)

func main() {

	var r string
	var numders string
	flag := true
	println("Укажите вариант расчета - 1 или 2")
	fmt.Scan(&r)

	for flag {
		switch r {
		case "1":
			fmt.Print("Введи трехзначное число - ") // Более крутой способ
			fmt.Scan(&numders)
			for len(numders) != 3 {
				print("Введите трехзначное число")
				fmt.Scan(&numders)
			}
			fmt.Println("Единицы", string(numders[2]))
			fmt.Println("Десятки", string(numders[1]))
			fmt.Println("Сотни", string(numders[0]))
			flag = false
		case "2":
			var numbers2 uint16
			fmt.Print("Введи трехзначное число - ") // Оптимальное решение
			fmt.Scan(&numbers2)
			hundred := numbers2 / 100
			fmt.Println("Сотни", hundred)
			dozen := (numbers2 - (hundred * 100)) / 10
			fmt.Println("Десятки", dozen)
			unit := (numbers2 - (hundred * 100) - (dozen * 10))
			fmt.Println("Единицы", unit)
			flag = false
		default:
			println("Введите 1 или 2")
			fmt.Scan(&r)
		}
	}
}
