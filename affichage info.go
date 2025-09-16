package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Structure Character
type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
}

// Fonction displayInfo : affiche les informations du personnage
func displayInfo(c Character) {
	fmt.Println("\n=== Fiche du Personnage ===")
	fmt.Println("Nom       :", c.Nom)
	fmt.Println("Classe    :", c.Classe)
	fmt.Println("Niveau    :", c.Niveau)
	fmt.Println("PV        :", c.PVActuels, "/", c.PVMax)
	fmt.Println("Inventaire:", c.Inventaire)
}

// Fonction pour initialiser un personnage
func initCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, inventaire []string) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
	}
}

// Fonction utilitaire pour lire un entier avec vérification
func lireEntier(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if n, err := strconv.Atoi(input); err == nil {
			return n
		}
		fmt.Println("Entrée invalide. Veuillez entrer un nombre entier.")
	}
}

// Fonction pour modifier un personnage
func modifierCharacter(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- Modification du personnage ---")

	fmt.Print("Nom : ")
	nom, _ := reader.ReadString('\n')
	c.Nom = strings.TrimSpace(nom)

	fmt.Print("Classe : ")
	classe, _ := reader.ReadString('\n')
	c.Classe = strings.TrimSpace(classe)

	c.Niveau = lireEntier("Niveau : ")
	c.PVMax = lireEntier("PV Max : ")
	c.PVActuels = lireEntier("PV Actuels : ")

	fmt.Print("Inventaire (séparé par des virgules) : ")
	invStr, _ := reader.ReadString('\n')
	invStr = strings.TrimSpace(invStr)
	c.Inventaire = strings.Split(invStr, ",")
	for i := range c.Inventaire {
		c.Inventaire[i] = strings.TrimSpace(c.Inventaire[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Création initiale du personnage
	c1 := initCharacter(
		"TonNom",
		"Elfe",
		1,
		100,
		40,
		[]string{"Potion", "Potion", "Potion"},
	)

	// Affichage initial
	displayInfo(c1)

	// Demander si l'utilisateur veut modifier
	fmt.Print("\nSouhaitez-vous modifier les informations ? (oui/non) : ")
	reponse, _ := reader.ReadString('\n')
	reponse = strings.ToLower(strings.TrimSpace(reponse))

	if reponse == "oui" || reponse == "o" {
		modifierCharacter(&c1)
	}

	// Affichage final
	displayInfo(c1)
	fmt.Println("\nLet's go !!!")
}
