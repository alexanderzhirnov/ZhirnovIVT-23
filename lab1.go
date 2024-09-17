package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		var choice int
		fmt.Println("Выберите задание:")
		fmt.Println("1. Вывести текущее время и дату")
		fmt.Println("2. Создать переменные различных типов и вывести их на экран")
		fmt.Println("3. Использовать краткую форму объявления переменных")
		fmt.Println("4. Арифметические операции с двумя целыми числами")
		fmt.Println("5. Сумма и разность двух чисел с плавающей запятой")
		fmt.Println("6. Вычислить среднее значение трёх чисел")
		fmt.Println("0. Выйти")
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			task1()
		case 2:
			task2()
		case 3:
			task3()
		case 4:
			task4()
		case 5:
			task5()
		case 6:
			task6()
		case 0:
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func task1() {
	currentTime := time.Now()
	fmt.Println("Текущие время и дата:", currentTime)
}

func task2() {
	var a int = 42
	var b float64 = 3.14
	var c string = "Hello, Go!"
	var d bool = true

	fmt.Println("int:", a)
	fmt.Println("float64:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)
}

func task3() {
	a := 42
	b := 3.14
	c := "Hello, Go!"
	d := true

	fmt.Println("int:", a)
	fmt.Println("float64:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)
}

func task4() {
	var a, b int
	fmt.Print("Введите два целых числа: ")
	fmt.Scan(&a, &b)

	fmt.Println("Сумма:", a+b)
	fmt.Println("Разность:", a-b)
	fmt.Println("Произведение:", a*b)
	if b != 0 {
		fmt.Println("Частное:", a/b)
	} else {
		fmt.Println("Деление на ноль!")
	}
}

func task5() {
	var a, b float64
	fmt.Print("Введите два числа с плавающей запятой: ")
	fmt.Scan(&a, &b)

	fmt.Println("Сумма:", a+b)
	fmt.Println("Разность:", a-b)
}

func task6() {
	var a, b, c float64
	fmt.Print("Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	average := (a + b + c) / 3
	fmt.Println("Среднее значение:", average)
}
