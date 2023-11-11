package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var fichierParametres, fichierResultats string

	fmt.Print("Entrez le nom du fichier de paramètres : ")
	fmt.Scan(&fichierParametres)
	fmt.Print("Entrez le nom du fichier de résultats : ")
	fmt.Scan(&fichierResultats)

	allumettes, players, takeMax, err := chargeParametres(fichierParametres)
	if err != nil {
		fmt.Println("Erreur lors de la lecture des paramètres :", err)
		return
	}

	winner := playGame(allumettes, players, takeMax)

	if err := saveResultat(fichierResultats, winner); err != nil {
		fmt.Println("Erreur lors de la sauvegarde du résultat de la partie :", err)
		return
	}
}

func chargeParametres(fichier string) (int, int, int, error) {
	file, err := os.Open(fichier)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()

	var allumettes, players, takeMax int
	_, err = fmt.Fscanln(file, &allumettes, &players, &takeMax)
	if err != nil {
		return 0, 0, 0, err
	}

	return allumettes, players, takeMax, nil
}

func saveResultat(fichier, winner string) error {
	file, err := os.Create(fichier)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, winner)
	if err != nil {
		return err
	}

	return nil
}

func playGame(allumettes, players, takeMax int) string {
	var CurrentPlayer int
	rand.Seed(time.Now().UnixNano())

	for allumettes > 0 {
		fmt.Printf("\nIl reste %d allumettes.\n", allumettes)

		var take int
		for {
			fmt.Printf("Joueur %d, choisissez le nombre d'allumettes souhaité (1 à %d) ? ", CurrentPlayer+1, takeMax)
			fmt.Scan(&take)
			if take >= 1 && take <= takeMax && take <= allumettes {
				break
			}
			fmt.Println("Nombre invalide. Veuillez choisir entre 1 et", takeMax)
		}
		allumettes -= take
		CurrentPlayer = (CurrentPlayer + 1) % players
	}

	winner := (CurrentPlayer+players-1)%players + 1
	return fmt.Sprintf("Joueur %d", winner)
}
