package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Character struct {
	Nom          string
	Classe       string
	Niveau       int
	PVMax        int
	PVActuels    int
	Force        int
	Intelligence int
	Initiative   int
	Argent       int
	Inventaire   []string
	Skills       []string
	Experience   int
	ExpToLevel   int
	InventaireMax int
	Équipement  *Equipped
	DegatsBase  int
}

func createCharacter(scanner *bufio.Scanner) Character {
	fmt.Print("🧍 Quel est ton nom, Hunter ? ")
	scanner.Scan()
	nom := scanner.Text()
	classe := choisirClasse(scanner)

	player := Character{
		Nom:          strings.Title(nom),
		Classe:       classe,
		Niveau:       1,
		PVMax:        100,
		PVActuels:    100,
		Force:        10,
		Intelligence: 5,
		Initiative:   10,
		Argent:       20,
		Inventaire:   []string{"Potion de vie"},
		Skills:       []string{"Coup de poing"},
		Experience:   0,
		ExpToLevel:   10,
		InventaireMax: 10,
		DegatsBase:   5,
	}

	switch classe {
	case "Renforcement":
		player.PVMax += 30
		player.Force += 5
		player.Skills = append(player.Skills, "Uppercut renforcé")
	case "Émission":
		player.Intelligence += 5
		player.Skills = append(player.Skills, "Projectile Nen")
	case "Manipulation":
		player.Skills = append(player.Skills, "Contrôle mental")
	case "Spécialisation":
		player.Skills = append(player.Skills, "Effet inconnu…")
	case "Transformation":
		player.Force += 2
		player.Skills = append(player.Skills, "Poing électrique")
	case "Conjuration":
		player.Skills = append(player.Skills, "Épée invoquée")
	}

	player.PVActuels = player.PVMax
	fmt.Printf("\n✅ Bienvenue, %s (%s)\n", player.Nom, player.Classe)
	return player
}

func choisirClasse(scanner *bufio.Scanner) string {
	classes := []string{"Renforcement", "Émission", "Manipulation", "Transformation", "Conjuration", "Spécialisation"}
	for {
		fmt.Println("\n🧠 Choisis ton type de Nen :")
		for i, cls := range classes {
			fmt.Printf("%d. %s\n", i+1, cls)
		}
		fmt.Print("👉 Ton choix : ")
		scanner.Scan()
		choix := strings.ToLower(scanner.Text())
		for _, cls := range classes {
			if strings.ToLower(cls) == choix {
				return cls
			}
		}
		fmt.Println("❌ Classe inconnue, réessaie.")
	}
}

func displayCharacterInfo(c Character) {
	fmt.Println("\n--- INFOS DU PERSONNAGE ---")
	fmt.Println("👤 Nom       :", c.Nom)
	fmt.Println("📘 Classe    :", c.Classe)
	fmt.Println("🔢 Niveau    :", c.Niveau)
	fmt.Printf("❤️ PV        : %d / %d\n", c.PVActuels, c.PVMax)
	fmt.Println("💪 Force     :", c.Force)
	fmt.Println("🧠 Intelligence :", c.Intelligence)
	fmt.Println("⚡ Initiative:", c.Initiative)
	fmt.Println("💰 Argent    :", c.Argent)
	fmt.Println("🎒 Inventaire:", c.Inventaire)
	fmt.Println("🌀 Compétences :", c.Skills)
	fmt.Printf("⭐ Expérience : %d / %d\n", c.Experience, c.ExpToLevel)

	// Affichage des équipements
	fmt.Println("\n--- ÉQUIPEMENTS ÉQUIPÉS ---")
	if c.Équipement == nil {
		fmt.Println("Aucun équipement équipé.")
	} else {
		if c.Équipement.Arme != nil {
			fmt.Printf("⚔️ Arme : %s (+%d Force)\n", c.Équipement.Arme.Nom, c.Équipement.Arme.BonusDmg)
		} else {
			fmt.Println("⚔️ Arme : Aucun")
		}
		if c.Équipement.Tete != nil {
			fmt.Printf("🎩 Tête : %s (+%d PV)\n", c.Équipement.Tete.Nom, c.Équipement.Tete.BonusPV)
		} else {
			fmt.Println("🎩 Tête : Aucun")
		}
		if c.Équipement.Torse != nil {
			fmt.Printf("🛡️ Torse : %s (+%d PV)\n", c.Équipement.Torse.Nom, c.Équipement.Torse.BonusPV)
		} else {
			fmt.Println("🛡️ Torse : Aucun")
		}
		if c.Équipement.Pieds != nil {
			fmt.Printf("👢 Pieds : %s (+%d PV)\n", c.Équipement.Pieds.Nom, c.Équipement.Pieds.BonusPV)
		} else {
			fmt.Println("👢 Pieds : Aucun")
		}
	}
}

