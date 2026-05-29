package main

import (
	"fmt"
)

func main() {
	var weight float64
	var height float64

	const IMCMaigreur float64 = 18.5
	const IMCNormal float64 = 25.5
	const IMCSurpoids float64 = 30.0

	// inputs
	fmt.Print("Enter weight in kg: ")
	fmt.Scan(&weight)

	fmt.Print("Enter height in meters: ")
	fmt.Scan(&height)

	fmt.Printf("Your weight is: %.2f kg\n", weight)
	fmt.Printf("Your height is: %.2f m\n", height)

	var imc = weight / (height * height)

	fmt.Printf("Your IMC is: %.2f\n", imc)

	switch {
	case imc < IMCMaigreur:
		fmt.Println("You are underweight.")
	case imc >= IMCMaigreur && imc < IMCNormal:
		fmt.Println("You have a normal weight.")
	case imc >= IMCNormal && imc < IMCSurpoids:
		fmt.Println("You are overweight.")
	default:
		fmt.Println("You are obese.")
	}

}
