package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	NbR := rand.Intn(100) + 1
	essais := 0

	fmt.Println("Bienvenue dans le jeu de devinette en Go !")
	fmt.Println("Je pense à un nombre entre 1 et 100. Pouvez-vous deviner ce nombre ?")

	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)

		if err != nil {
			fmt.Println("Veuillez entrer un nombre valide.")
			continue
		}

		essais++

		if guess < NbR {
			fmt.Println("Et non, le nombre est plus grand")

		} else if guess > NbR {
			fmt.Println("Et non, le nombre est plus petit")

		} else {
			fmt.Printf("Félicitations, vous avez trouvé le bon nombre : %d\n", NbR)
			break
		}
	}
}
