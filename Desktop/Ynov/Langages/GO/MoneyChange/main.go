package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

}
func Payement() float64 {
	var NbrRand float64
	NbrRand = 0.5 + rand.Float64()*(200.0)

	fmt.Println("Comment souhaitez-vous régler ?")
	fmt.Println("1. Régler par carte")
	fmt.Println("2. Régler en espèce")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Vous avez choisi de régler par carte")
		fmt.Println("Vous pouvez sortir")
		return NbrRand
	case 2:
		fmt.Println("Vous avez avez choisi de régler en espèce")
		fmt.Println("Veuillez insérer un billet ou une somme")
		var input float64
		fmt.Scanln(&input)
		return input

	default:
		fmt.Println("Option non reconnue")
		return 0.0
	}
}

func calculerMontant(duree int) float64 {
	var montant float64

	switch {
	case duree < 1:
		montant = 0.0 // Stationnement gratuit pour moins d'une heure
	case duree >= 1 && duree < 2:
		montant = 5.0 // Prix pour une heure
	case duree >= 2 && duree < 4:
		montant = 10.0 // Prix pour deux heures
	case duree >= 4 && duree < 12:
		montant = 20.0 // Prix pour quatre heures
	case duree >= 12 && duree < 24:
		montant = 40.0 // Prix pour douze heures
	default:
		montant = 50.0 // Prix pour plus d'une journée (24 heures)
	}

	return montant
}

func Sortie(Prix float64) {
	rand.Seed(time.Now().UnixNano())
	Prix = NbrRand
	for {
		if Prix == 0.0 {
			fmt.Println("Vous êtes resté moins d'une heure, vous n'avez pas à payer.")
			break
		}
		if Prix > 0.0 && Prix < 5.0 {
			fmt.Printf("Vous êtes restez plus d'une heure, voici le montant à payer : %f", Prix)
			Payement()
		}
		if Prix > 5.0 && Prix < 20.0 {
			fmt.Printf("Vous êtes restez plus de deux heures, voici le montant à payer : %f", Prix)
			Payement()
		}
		if Prix > 20.0 && Prix < 75.0 {
			fmt.Printf("Vous êtes restez plus de quatres heures, voici le montant à payer : %f", Prix)
			Payement()
		}
		if Prix > 75.0 && Prix < 150.0 {
			fmt.Printf("Vous êtes restez plus de douzes heures, voici le montant à payer : %f", Prix)
			Payement()
		}
		if Prix > 150.0 && Prix < 200.0 {
			fmt.Printf("Vous êtes restez au moins une journée, voici le montant à payer : %f", Prix)
			Payement()
		}
	}
}
