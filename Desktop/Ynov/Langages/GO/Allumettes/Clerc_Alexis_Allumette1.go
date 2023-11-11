package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Bot bool
var n = rand.Intn(30) + 4

func main() {
	rand.Seed(time.Now().UnixNano())
	Menu()
}

func Menu() {
	var choixnb int
	fmt.Println("Souhaitez-vous choisir le nombre d'allumettes ou jouer avec un nombre aléatoire ? (1 : Choix / 2 : Aléatoire)")
	fmt.Scan(&choixnb)

	var choice int
	if choixnb == 1 {
		fmt.Println("Avec combien d'allumettes souhaitez-vous jouer ? (Entre 4 et 30)")
		fmt.Scan(&n)
	}

	fmt.Println("Combien de joueurs êtes-vous ? (1 / 2)")
	fmt.Scan(&choice)

	Bot = (choice == 1)

	Pvp()
}

func TakeMatches(player string, maxMatches int) {
	var choix int
	fmt.Printf("Tour de %s\n", player)
	fmt.Printf("Il reste %d allumettes\n", n)

	for {
		fmt.Println("Combien d'allumettes souhaitez-vous prendre ? (1 / 2 / 3)")
		fmt.Scan(&choix)
		if choix >= 1 && choix <= maxMatches {
			break
		}
	}

	n -= choix

	if n <= 0 {
		fmt.Printf("%s a gagné\n", player)
	} else {
		fmt.Printf("Il reste %d allumettes\n", n)
	}
}

func Pvp() {
	for i := 1; n > 0; i++ {
		if i%2 != 0 {
			TakeMatches("Joueur 1", 3)
		} else if Bot {
			TakeMatches("IA", 3)
		} else {
			TakeMatches("Joueur 2", 3)
		}
	}
}