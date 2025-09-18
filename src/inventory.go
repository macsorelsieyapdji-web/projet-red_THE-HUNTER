package main

import (
	"bufio"
	"fmt"
)

// Ouvre l’inventaire et permet d’utiliser ou d’équiper un objet
func openInventory(c *Character, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n🎒 --- INVENTAIRE ---")
		if len(c.Inventaire) == 0 {
			fmt.Println("Aucun objet dans l’inventaire.")
			return
		}

		for i, item := range c.Inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Println("0. Quitter l’inventaire")
		fmt.Print("Choisis un objet : ")

		scanner.Scan()
		choix := scanner.Text()
		index := parseInt(choix) - 1

		if choix == "0" {
			return
		}
		if index < 0 || index >= len(c.Inventaire) {
			fmt.Println("❌ Choix invalide")
			continue
		}

		item := c.Inventaire[index]

		// Vérifie si c’est un équipement
		equip := GetEquipmentByName(item)
		if equip != nil {
			EquipItem(c, equip)
		} else {
			useItem(c, item)
			// Retire l’objet consommable
			c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
		}
	}
}

// Utilisation d’un item
func useItem(c *Character, item string) {
	switch item {
	case "Potion de vie":
		c.PVActuels += 30
		if c.PVActuels > c.PVMax {
			c.PVActuels = c.PVMax
		}
		fmt.Printf("💊 Potion utilisée : +30 PV (%d/%d)\n", c.PVActuels, c.PVMax)
	case "Potion d’énergie":
		c.Initiative += 3
		fmt.Println("⚡ +3 Initiative")
	case "Cristal Nen":
		c.Intelligence += 5
		fmt.Println("🌀 +5 Intelligence")
	default:
		fmt.Println("❌ Objet inconnu")
	}
}
