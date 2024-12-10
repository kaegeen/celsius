package main

import (
	"fmt"
)

// celsiusToFahrenheit converts Celsius to Fahrenheit.
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// fahrenheitToCelsius converts Fahrenheit to Celsius.
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var choice int
	var temp float64

	fmt.Println("Temperature Converter")
	fmt.Println("======================")
	fmt.Println("1. Convert Celsius to Fahrenheit")
	fmt.Println("2. Convert Fahrenheit to Celsius")
	fmt.Print("Enter your choice (1 or 2): ")

	_, err := fmt.Scan(&choice)
	if err != nil || (choice != 1 && choice != 2) {
		fmt.Println("Invalid choice. Please enter 1 or 2.")
		return
	}

	fmt.Print("Enter the temperature to convert: ")
	_, err = fmt.Scan(&temp)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	if choice == 1 {
		fmt.Printf("%.2f°C is equal to %.2f°F\n", temp, celsiusToFahrenheit(temp))
	} else if choice == 2 {
		fmt.Printf("%.2f°F is equal to %.2f°C\n", temp, fahrenheitToCelsius(temp))
	}
}
