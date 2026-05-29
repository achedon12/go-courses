package main

import "fmt"

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division par zéro")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("opérateur inconnu: %s", op)
	}
}

func creerOperation(op string) (func(float64, float64) float64, error) {
	switch op {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}, nil
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}, nil
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}, nil
	case "/":
		return func(a, b float64) float64 {
			if b == 0 {
				return 0
			}
			return a / b
		}, nil
	default:
		return nil, fmt.Errorf("opérateur inconnu: %s", op)
	}
}

func main() {
	// exercice 1
	var a, b float64 = 8, 0
	var op = "/"

	resultat, err := operer(a, b, op)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Printf("Résultat: %.2f\n", resultat)
	}

	// exercice 3
	fmt.Println("\n--- Calculatrice Interactive ---")
	for {
		var num1, num2 float64
		var operation string

		fmt.Print("Entrez deux nombres et une opération (+, -, *, /) ou 'quit': ")
		_, err := fmt.Scan(&num1, &num2, &operation)

		if err != nil {
			fmt.Println("Erreur de lecture")
			continue
		}
		if operation == "quit" {
			fmt.Println("Au revoir!")
			break
		}

		resultat, err := operer(num1, num2, operation)
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Printf("Résultat: %.2f\n", resultat)
		}
	}
}
