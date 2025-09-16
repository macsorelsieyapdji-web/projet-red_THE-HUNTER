package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Définition de la structure Character
type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
}

// Méthode pour afficher les infos du personnage
func (c Character) Afficher() {
	fmt.Println("\n=== Fiche du Personnage ===")
	fmt.Println("Nom       :", c.Nom)
	fmt.Println("Classe    :", c.Classe)
	fmt.Println("Niveau    :", c.Niveau)
	fmt.Println("PV        :", c.PVActuels, "/", c.PVMax)
	fmt.Println("Inventaire:", c.Inventaire)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var perso Character

	// Saisie du nom
	fmt.Print("Entrez le nom du personnage : ")
	perso.Nom, _ = reader.ReadString('\n')
	perso.Nom = strings.TrimSpace(perso.Nom)

	// Saisie de la classe
	fmt.Print("Entrez la classe du personnage : ")
	perso.Classe, _ = reader.ReadString('\n')
	perso.Classe = strings.TrimSpace(perso.Classe)

	// Saisie du niveau
	fmt.Print("Entrez le niveau du personnage : ")
	niveauStr, _ := reader.ReadString('\n')
	niveauStr = strings.TrimSpace(niveauStr)
	perso.Niveau, _ = strconv.Atoi(niveauStr)

	// Saisie des PV max
	fmt.Print("Entrez les points de vie maximum : ")
	pvMaxStr, _ := reader.ReadString('\n')
	pvMaxStr = strings.TrimSpace(pvMaxStr)
	perso.PVMax, _ = strconv.Atoi(pvMaxStr)

	// Saisie des PV actuels
	fmt.Print("Entrez les points de vie actuels : ")
	pvActuelsStr, _ := reader.ReadString('\n')
	pvActuelsStr = strings.TrimSpace(pvActuelsStr)
	perso.PVActuels, _ = strconv.Atoi(pvActuelsStr)

	// Saisie de l'inventaire
	fmt.Print("Entrez les objets de l'inventaire (séparés par des virgules) : ")
	inventaireStr, _ := reader.ReadString('\n')
	inventaireStr = strings.TrimSpace(inventaireStr)
	perso.Inventaire = strings.Split(inventaireStr, ",")

	// Nettoyage des espaces autour des objets
	for i, item := range perso.Inventaire {
		perso.Inventaire[i] = strings.TrimSpace(item)
	}

	// Affichage des infos
	perso.Afficher()

	// Affichage final
	fmt.Println("\nLet's go !!!")
}
