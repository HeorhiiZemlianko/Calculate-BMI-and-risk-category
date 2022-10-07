package main

import (
	"fmt"
	"math"
)

func main() {
	var weight, height, bmi float64
	// Считывание веса и роста
	fmt.Print("Введите свой вес, кг: ")
	fmt.Scan(&weight)
	fmt.Print("Введите свой рост, см: ")
	fmt.Scan(&height)

	// Вывод внесенных данных: веса и роста
	fmt.Println("\nВы ввели вес, кг: ", weight)
	fmt.Println("Вы ввели рост, см: ", height)

	// Расчет ИМТ и перевод см в метры
	heightInMeter := height / 100
	bmi = weight / math.Pow(heightInMeter, 2)

	// Вывод полученного ИМТ и категории
	fmt.Printf("\nВаш ИМТ: %0.2f", bmi)
	fmt.Print("\nВаша категория риска: ")

	// Условия проверки категории
	if weight < ((height - 100) - (heightInMeter * 10) + 10) {
		fmt.Println("Нужно поправиться.")
	} else if weight > ((height - 100) + (heightInMeter * 10)) {
		fmt.Println("Нужно похудеть.")
	} else {
		fmt.Println("Нормальный вес.")
	}

	// Условия проверки веса по ИМТ
	normalWeight := (height - 100) - (height-150)/2
	delta := weight - normalWeight
	fmt.Printf("Идеальный вес для вашего роста: %0.2v кг.\n", normalWeight)

	if (delta > 0) && (bmi > 30) {
		fmt.Printf("Вам нужно похудеть до идеального веса на %0.2v кг.\n", math.Abs(delta))
	} else if (delta > 0) && (bmi > 25) {
		fmt.Printf("Вам нужно похудеть до идеального веса на %0.2v кг.\n", math.Abs(delta))
	} else if (delta < 0) && (bmi < float64(18.5)) {
		fmt.Printf("Вам нужно поправится до идеального веса на %0.2v кг.\n", math.Abs(delta))
	}
}
