package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Produit struct {
	id        int
	nom       string
	marque    string
	prix      float64
	stock     int
	categorie string
	actif     bool
}

type Catalogue struct {
	produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, produit := range c.produits {
		if produit.id == p.id {
			return fmt.Errorf("erreur: un produit avec l'ID %d existe déjà", p.id)
		}
	}

	c.produits = append(c.produits, p)
	return nil
}

func (c *Catalogue) TrouverParId(id int) (Produit, error) {
	for _, produit := range c.produits {
		if produit.id == id {
			return produit, nil
		}
	}
	return Produit{}, fmt.Errorf("erreur: produit avec l'ID %d non trouvé", id)
}

func (c *Catalogue) TrouverParCategorie(categorie string) []Produit {
	var resultats []Produit
	for _, produit := range c.produits {
		if produit.categorie == categorie {
			resultats = append(resultats, produit)
		}
	}
	return resultats
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	var modified = 0
	for _, produit := range c.produits {
		if produit.categorie == categorie {
			produit.prix = produit.prix * (1 - pct/100)
			modified++
		}
	}
	return modified
}

func (c *Catalogue) Vendre(id int, quantite int) error {
	for i, produit := range c.produits {
		if produit.id == id {
			if produit.stock < quantite {
				return fmt.Errorf("erreur: stock insuffisant pour le produit ID %d (disponible: %d, demandé: %d)", id, produit.stock, quantite)
			}
			c.produits[i].stock -= quantite
			return nil
		}
	}
	return fmt.Errorf("erreur: produit avec l'ID %d non trouvé", id)
}

func (c *Catalogue) Rapport() string {
	totalValeur := 0.0
	for _, produit := range c.produits {
		totalValeur += produit.prix * float64(produit.stock)
	}
	return fmt.Sprintf("nb produits: %d, valeur totale du stock: %.2f", len(c.produits), totalValeur)
}

func main() {
	catalogue := &Catalogue{}

	// Pré-remplissage du catalogue
	catalogue.AjouterProduit(Produit{1, "iPhone 15", "Apple", 999.99, 50, "téléphones", true})
	catalogue.AjouterProduit(Produit{2, "MacBook Pro", "Apple", 1999.99, 20, "ordinateurs", true})
	catalogue.AjouterProduit(Produit{3, "Galaxy S24", "Samsung", 899.99, 40, "téléphones", true})
	catalogue.AjouterProduit(Produit{4, "AirPods Pro", "Apple", 249.99, 100, "audio", true})
	catalogue.AjouterProduit(Produit{5, "iPad Pro", "Apple", 1299.99, 30, "tablettes", true})

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== MENU TechShop ===")
		fmt.Println("[1] Ajouter")
		fmt.Println("[2] Chercher")
		fmt.Println("[3] Soldes")
		fmt.Println("[4] Vendre")
		fmt.Println("[5] Rapport")
		fmt.Println("[6] Quitter")
		fmt.Print("Choisir une option: ")

		scanner.Scan()
		choix := scanner.Text()

		switch choix {
		case "1":
			fmt.Print("ID: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Nom: ")
			scanner.Scan()
			nom := scanner.Text()

			fmt.Print("Marque: ")
			scanner.Scan()
			marque := scanner.Text()

			fmt.Print("Prix: ")
			scanner.Scan()
			prix, _ := strconv.ParseFloat(scanner.Text(), 64)

			fmt.Print("Stock: ")
			scanner.Scan()
			stock, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Catégorie: ")
			scanner.Scan()
			categorie := scanner.Text()

			p := Produit{id, nom, marque, prix, stock, categorie, true}
			if err := catalogue.AjouterProduit(p); err != nil {
				fmt.Printf("Erreur: %v\n", err)
			} else {
				fmt.Println("Produit ajouté avec succès!")
			}

		case "2":
			fmt.Print("Chercher par [1] ID ou [2] Catégorie? ")
			scanner.Scan()
			sousChoix := scanner.Text()

			if sousChoix == "1" {
				fmt.Print("ID à chercher: ")
				scanner.Scan()
				id, _ := strconv.Atoi(scanner.Text())

				if p, err := catalogue.TrouverParId(id); err != nil {
					fmt.Printf("Erreur: %v\n", err)
				} else {
					fmt.Printf("Trouvé: ID=%d, Nom=%s, Marque=%s, Prix=%.2f, Stock=%d, Catégorie=%s\n",
						p.id, p.nom, p.marque, p.prix, p.stock, p.categorie)
				}
			} else if sousChoix == "2" {
				fmt.Print("Catégorie à chercher: ")
				scanner.Scan()
				categorie := scanner.Text()

				resultats := catalogue.TrouverParCategorie(categorie)
				if len(resultats) == 0 {
					fmt.Println("Aucun produit trouvé dans cette catégorie")
				} else {
					fmt.Printf("Produits trouvés (%d):\n", len(resultats))
					for _, p := range resultats {
						fmt.Printf("  ID=%d, Nom=%s, Marque=%s, Prix=%.2f, Stock=%d\n",
							p.id, p.nom, p.marque, p.prix, p.stock)
					}
				}
			}

		case "3":
			fmt.Print("Catégorie pour appliquer la réduction: ")
			scanner.Scan()
			categorie := scanner.Text()

			fmt.Print("Pourcentage de réduction: ")
			scanner.Scan()
			pct, _ := strconv.ParseFloat(scanner.Text(), 64)

			modified := catalogue.AppliquerReduction(categorie, pct)
			fmt.Printf("Réduction appliquée à %d produit(s)\n", modified)

		case "4":
			fmt.Print("ID du produit à vendre: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			fmt.Print("Quantité à vendre: ")
			scanner.Scan()
			quantite, _ := strconv.Atoi(scanner.Text())

			if err := catalogue.Vendre(id, quantite); err != nil {
				fmt.Printf("Erreur: %v\n", err)
			} else {
				fmt.Println("Vente effectuée avec succès!")
			}

		case "5":
			fmt.Println(catalogue.Rapport())

		case "6":
			fmt.Println("Au revoir!")
			os.Exit(0)

		default:
			fmt.Println("Option invalide")
		}
	}
}
