package main

import (
	"fmt"
	"math"
	"strings"
)

/*
EXERCICE : Système de paiement
Contexte : Un système e-commerce doit accepter plusieurs modes de paiement.

- Payeur : interface avec Payer(montant float64) (string, error)
- CarteCredit : déduit du solde, "Transaction CB #XXXX confirmée"
- PayPal      : déduit du solde, "Paiement PayPal de X€ vers email"
- Crypto      : convertit le montant en crypto (1 BTC = 50000€)
- ProcesserPanier : affiche le total, appelle Payer, type switch sur le mode
*/

const TauxBTC = 50000.0

type Payeur interface {
	Payer(montant float64) (string, error)
}

// ErrSoldeInsuffisant est renvoyée quand le moyen de paiement ne peut pas couvrir le montant.
type ErrSoldeInsuffisant struct {
	Montant float64
	Solde   float64
}

func (e *ErrSoldeInsuffisant) Error() string {
	return fmt.Sprintf("solde insuffisant : %.2f€ demandés, %.2f€ disponibles", e.Montant, e.Solde)
}

type CarteCredit struct {
	Numero    string
	Titulaire string
	Solde     float64
}

func (cc *CarteCredit) Payer(montant float64) (string, error) {
	if montant > cc.Solde {
		return "", &ErrSoldeInsuffisant{Montant: montant, Solde: cc.Solde}
	}
	cc.Solde -= montant
	return fmt.Sprintf("Transaction CB #%s confirmée", cc.QuatreDerniers()), nil
}

// QuatreDerniers renvoie les 4 derniers chiffres du numéro, ou le numéro entier s'il est plus court.
func (cc *CarteCredit) QuatreDerniers() string {
	if len(cc.Numero) <= 4 {
		return cc.Numero
	}
	return cc.Numero[len(cc.Numero)-4:]
}

func (cc *CarteCredit) Reseau() string {
	switch {
	case strings.HasPrefix(cc.Numero, "4"):
		return "Visa"
	case strings.HasPrefix(cc.Numero, "5"):
		return "Mastercard"
	default:
		return "Inconnu"
	}
}

type PayPal struct {
	Email string
	Solde float64
}

func (pp *PayPal) Payer(montant float64) (string, error) {
	if montant > pp.Solde {
		return "", &ErrSoldeInsuffisant{Montant: montant, Solde: pp.Solde}
	}
	pp.Solde -= montant
	return fmt.Sprintf("Paiement PayPal de %.2f€ vers %s", montant, pp.Email), nil
}

type Crypto struct {
	Adresse string
	Solde   float64
	Monnaie string
}

func (c *Crypto) Payer(montant float64) (string, error) {
	if montant > c.Solde {
		return "", &ErrSoldeInsuffisant{Montant: montant, Solde: c.Solde}
	}
	c.Solde -= montant
	montantCrypto := math.Round(montant/TauxBTC*1000) / 1000
	return fmt.Sprintf("Paiement de %.3f %s (%.2f€) vers %s", montantCrypto, c.Monnaie, montant, c.Adresse), nil
}

// Vérifications statiques : les trois types satisfont bien Payeur.
var (
	_ Payeur = &CarteCredit{}
	_ Payeur = &PayPal{}
	_ Payeur = &Crypto{}
)

func ProcesserPanier(payeur Payeur, articles []float64) {
	total := 0.0
	for _, prix := range articles {
		total += prix
	}

	fmt.Printf("Panier : %d article(s) — total %.2f€\n", len(articles), total)

	switch p := payeur.(type) {
	case *CarteCredit:
		fmt.Printf("Mode : Carte de crédit %s (%s)\n", p.Reseau(), p.Titulaire)
	case *PayPal:
		fmt.Printf("Mode : PayPal (%s)\n", p.Email)
	case *Crypto:
		fmt.Printf("Mode : Crypto %s\n", p.Monnaie)
	default:
		fmt.Printf("Mode : inconnu (%T)\n", p)
	}

	recu, err := payeur.Payer(total)
	if err != nil {
		fmt.Println("Échec :", err)
		fmt.Println()
		return
	}
	fmt.Println("OK :", recu)
	fmt.Println()
}

func main() {
	carte := &CarteCredit{Numero: "4539123456789012", Titulaire: "Alice Martin", Solde: 500}
	paypal := &PayPal{Email: "bob@example.com", Solde: 120}
	crypto := &Crypto{Adresse: "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", Solde: 10000, Monnaie: "BTC"}

	ProcesserPanier(carte, []float64{29.99, 149.50, 15.00})
	ProcesserPanier(paypal, []float64{89.90, 45.00})
	ProcesserPanier(crypto, []float64{2500, 1200.50})

	fmt.Println("=== Soldes restants ===")
	fmt.Printf("CB #%s   : %.2f€\n", carte.QuatreDerniers(), carte.Solde)
	fmt.Printf("PayPal     : %.2f€\n", paypal.Solde)
	fmt.Printf("Crypto     : %.2f€\n", crypto.Solde)
}
