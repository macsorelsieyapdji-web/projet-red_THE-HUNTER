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

// Méthode pour afficher le personnage
func (c Character) Afficher() {
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

	fmt.Print("Niveau : ")
	niveauStr, _ := reader.ReadString('\n')
	niveauStr = strings.TrimSpace(niveauStr)
	c.Niveau, _ = strconv.Atoi(niveauStr)

	fmt.Print("PV Max : ")
	pvMaxStr, _ := reader.ReadString('\n')
	pvMaxStr = strings.TrimSpace(pvMaxStr)
	c.PVMax, _ = strconv.Atoi(pvMaxStr)

	fmt.Print("PV Actuels : ")
	pvActuelsStr, _ := reader.ReadString('\n')
	pvActuelsStr = strings.TrimSpace(pvActuelsStr)
	c.PVActuels, _ = strconv.Atoi(pvActuelsStr)

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
	c1.Afficher()

	// Demander si l'utilisateur veut modifier
	fmt.Print("\nSouhaitez-vous modifier les informations ? (oui/non) : ")
	reponse, _ := reader.ReadString('\n')
	reponse = strings.ToLower(strings.TrimSpace(reponse))

	if reponse == "oui" || reponse == "o" {
		modifierCharacter(&c1)
	}

	// Affichage final
	c1.Afficher()
	fmt.Println("\nLet's go !!!")
}
