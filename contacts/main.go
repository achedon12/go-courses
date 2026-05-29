package main

import (
	"fmt"
)

type Personne struct {
	nom    string
	prenom string
	age    int
	email  string
}

type Adresse struct {
	rue        string
	ville      string
	codePostal string
}

type Employe struct {
	personne Personne
	adresse  Adresse
	poste    string
	salaire  float64
}

type Etudiant struct {
	Personne
	promo   string
	moyenne float64
}

func (p Personne) NomComplet() string {
	return p.prenom + " " + p.nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("Bonjour, je m'appelle %s %s, j'ai %d ans et mon adresse email est %s.", p.prenom, p.nom, p.age, p.email)
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.rue, a.codePostal, a.ville)
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf("Employé: %s\nPoste: %s\nSalaire: %.2f\nAdresse: %s", e.personne.NomComplet(), e.poste, e.salaire, e.adresse.Format())
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.salaire += e.salaire * pct / 100
}

func (et Etudiant) MentionObtenue() string {
	switch {
	case et.moyenne >= 16:
		return "TB" // Très Bien
	case et.moyenne >= 14:
		return "B" // Bien
	case et.moyenne >= 12:
		return "AB" // Assez Bien
	case et.moyenne >= 10:
		return "P" // Passable
	default:
		return "Insuffisant"
	}
}

func main() {
	fmt.Println("========== FICHES EMPLOYÉS ==========")

	employe1 := Employe{
		personne: Personne{
			nom:    "Durand",
			prenom: "Marc",
			age:    35,
			email:  "marc.durand@email.com",
		},
		adresse: Adresse{
			rue:        "123 rue de la Paix",
			ville:      "Paris",
			codePostal: "75001",
		},
		poste:   "Développeur Go",
		salaire: 45000,
	}

	employe2 := Employe{
		personne: Personne{
			nom:    "Bernard",
			prenom: "Sophie",
			age:    28,
			email:  "sophie.bernard@email.com",
		},
		adresse: Adresse{
			rue:        "456 avenue des Champs",
			ville:      "Lyon",
			codePostal: "69000",
		},
		poste:   "Chef de Projet",
		salaire: 55000,
	}

	fmt.Println("\n--- Employé 1 ---")
	fmt.Println(employe1.FicheEmploye())

	fmt.Println("\n--- Employé 2 ---")
	fmt.Println(employe2.FicheEmploye())

	fmt.Println("\n========== FICHES ÉTUDIANTS ==========")

	etudiant1 := Etudiant{
		Personne: Personne{
			nom:    "Lefevre",
			prenom: "Pierre",
			age:    21,
			email:  "pierre.lefevre@email.com",
		},
		promo:   "L3",
		moyenne: 13.8,
	}

	etudiant2 := Etudiant{
		Personne: Personne{
			nom:    "Rousseau",
			prenom: "Marie",
			age:    20,
			email:  "marie.rousseau@email.com",
		},
		promo:   "L2",
		moyenne: 17.2,
	}

	fmt.Println("\n--- Étudiant 1 ---")
	fmt.Printf("Nom: %s\nPromo: %s\nMoyenne: %.1f\nMention: %s\n", etudiant1.NomComplet(), etudiant1.promo, etudiant1.moyenne, etudiant1.MentionObtenue())

	fmt.Println("\n--- Étudiant 2 ---")
	fmt.Printf("Nom: %s\nPromo: %s\nMoyenne: %.1f\nMention: %s\n", etudiant2.NomComplet(), etudiant2.promo, etudiant2.moyenne, etudiant2.MentionObtenue())
}
