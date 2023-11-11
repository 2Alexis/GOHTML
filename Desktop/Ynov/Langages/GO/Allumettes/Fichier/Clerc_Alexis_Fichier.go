package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Sélectionner un fichier txt")
		fmt.Println("2. Créer un nouveau fichier txt")
		fmt.Println("3. Supprimer le fichier txt")
		fmt.Println("4. Quitter")

		var choix int
		fmt.Print("Que souhaitez vous faire ? ")
		_, err := fmt.Scan(&choix)
		if err != nil {
			fmt.Println("Erreur de saisie.")
			continue
		}

		switch choix {
		case 1:
			chooseFile()
		case 2:
			makeFichier()
		case 3:
			dltFichier()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func chooseFile() {
	var nameFile string
	fmt.Print("Entrez le nom du fichier : ")
	_, err := fmt.Scan(&nameFile)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		return
	}

	contenu, err := os.ReadFile(nameFile)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier.")
		return
	}

	fmt.Println("Contenu du fichier :")
	fmt.Println(string(contenu))

	fmt.Println("1. Ajouter du texte ")
	fmt.Println("2. Supprimer le contenu")
	fmt.Println("3. Remplacer le contenu")

	var choix int
	fmt.Print("Choix : ")
	_, err = fmt.Scan(&choix)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		return
	}

	switch choix {
	case 1:
		ajoutertexte(nameFile)
	case 2:
		err := os.WriteFile(nameFile, []byte{}, 0666)
		if err != nil {
			fmt.Println("Erreur de la suppression du contenu.")
		} else {
			fmt.Println("Contenu supprimé.")
		}
	case 3:
		changeContenu(nameFile)
	default:
		fmt.Println("Choix invalide.")
	}
}

func ajoutertexte(nameFile string) {
	f, err := os.OpenFile(nameFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erreur d'ouverture du fichier.")
		return
	}
	defer f.Close()

	buffer := make([]byte, 4096)
	n, err := os.Stdin.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Erreur lors de la lecture de l'entrée.")
		return
	}

	txt := fmt.Sprintf("%s", buffer[:n])

	_, err = f.WriteString(txt)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier.")
		return
	}

	fmt.Println("Ecrivez votre texte :")
	n, err = os.Stdin.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Erreur lors de la lecture de l'entrée.")
		return
	}

	txt = fmt.Sprintf("%s", buffer[:n])

	_, err = f.WriteString(txt)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier.")
		return
	}

	fmt.Println("txt ajouté avec succès.")
}

func changeContenu(nameFile string) {
	fmt.Print("Entrez le nouveau txt : ")
	var txt string
	_, err := fmt.Scan(&txt)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		return
	}

	err = os.WriteFile(nameFile, []byte(txt), 0666)
	if err != nil {
		fmt.Println("Erreur lors de la modification du contenu.")
	} else {
		fmt.Println("Contenu remplacé avec succès.")
	}
}

func dltFichier() {
	var nameFile string
	fmt.Print("Entrez le nom du fichier à supprimer : ")
	_, err := fmt.Scan(&nameFile)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		return
	}

	err = os.Remove(nameFile)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier.")
	} else {
		fmt.Println("Fichier supprimé avec succès.")
	}
}

func makeFichier() {
	var nameFile string
	fmt.Print("Entrez le nom du nouveau fichier : ")
	_, err := fmt.Scan(&nameFile)
	if err != nil {
		fmt.Println("Erreur de saisie.")
		return
	}

	_, err = os.Create(nameFile)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier.")
	} else {
		fmt.Println("Fichier créé avec succès.")
	}
}

/*#Accus -N.M*/
