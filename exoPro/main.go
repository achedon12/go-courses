package main

import "fmt"

/*
EXERCICE : Système de bonus pour commerciaux
Contexte : Une entreprise de vente souhaite calculer les bonus des commerciaux
en fonction de leur montant de ventes.

- Montant >= 50000€ : Bonus 10% + Cadeau + Voyage
- Montant >= 30000€ : Bonus 7% + Cadeau
- Montant >= 10000€ : Bonus 3%
- Montant < 10000€  : Pas de bonus
*/

const (
	SansBonus = iota
	BonusSeul
	BonusEtCadeau
	BonusEtCadeauVoyage
)

type Commercial struct {
	nom     string
	montant float64
	bonuses []string
}

func calculerBonus(commercial *Commercial) float64 {
	bonus := 0.0
	commercial.bonuses = []string{}

	niveau := SansBonus
	switch {
	case commercial.montant >= 50000:
		niveau = BonusEtCadeauVoyage
	case commercial.montant >= 30000:
		niveau = BonusEtCadeau
	case commercial.montant >= 10000:
		niveau = BonusSeul
	}

	switch niveau {
	case BonusEtCadeauVoyage:
		bonus += commercial.montant * 0.10
		commercial.bonuses = append(commercial.bonuses, "Voyage international")
		fallthrough
	case BonusEtCadeau:
		commercial.bonuses = append(commercial.bonuses, "Cadeau executive")
		fallthrough
	case BonusSeul:
		commercial.bonuses = append(commercial.bonuses, "Bonus en cash")
	}

	return bonus
}

func afficherRapport(commerciaux []Commercial) {
	fmt.Println("=== RAPPORT DE BONUS DES COMMERCIAUX ===\n")

	totalBonus := 0.0

	for i, c := range commerciaux {
		bonus := calculerBonus(&c)
		totalBonus += bonus

		fmt.Printf("%d. %s\n", i+1, c.nom)
		fmt.Printf("   Montant ventes: %.2f€\n", c.montant)
		fmt.Printf("   Bonus en cash: %.2f€\n", bonus)

		fmt.Print("   Avantages: ")
		for j, avantage := range c.bonuses {
			fmt.Print(avantage)
			if j < len(c.bonuses)-1 {
				fmt.Print(" + ")
			}
		}
		fmt.Println("\n")
	}

	fmt.Printf("Total bonus distribués: %.2f€\n", totalBonus)
}

func main() {
	commerciaux := []Commercial{
		{nom: "Alice Martin", montant: 65000},
		{nom: "Bob Dupont", montant: 35000},
		{nom: "Charlie Durand", montant: 12000},
		{nom: "Diana Lefevre", montant: 8000},
	}

	afficherRapport(commerciaux)
}
