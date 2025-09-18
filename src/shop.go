package main

import (
	"bufio"
	"fmt"
	"strconv"
)

// Structure d'un item vendu par le marchand
type ShopItem struct {
	Name  string
	Price int
}

// Liste des items vendus
var ShopItems = []ShopItem{
	{"Potion de vie", 3},
	{"Potion d’énergie", 5},
	{"Cristal Nen", 10},
	{"Poings Renforcés", 15},
	{"Dague de Kurapika", 25},
	{"Katana de Zushi", 35},
	{"Gilet pare-chocs", 10},
	{"Veste en Nen", 20},
	{"Cape de l’Ombre", 30},
	{"Chapeau de l’aventurier", 5},
	{"Tunique de l’aventurier", 5},
	{"Bottes de l’aventurier", 5},
	{"billes de gun", 5},
}

// Ouvre le shop et permet l'achat
func OpenShop(c *Character, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Marchand ---")
		fmt.Println("💰 Argent :", c.Argent)
		for i, item := range ShopItems {
			fmt.Printf("%d. %s (%d pièces)\n", i+1, item.Name, item.Price)
		}
		fmt.Println("0. Quitter le marchand")
		fmt.Print("Choisis un item : ")

		scanner.Scan()
		choix := scanner.Text()
		index := parseInt(choix) - 1

		if choix == "0" {
			return
		}

		if index < 0 || index >= len(ShopItems) {
			fmt.Println("❌ Choix invalide")
			continue
		}

		item := ShopItems[index]
		if c.Argent < item.Price {
			fmt.Println("❌ Pas assez d’argent pour acheter", item.Name)
			continue
		}

		c.Argent -= item.Price

		// Vérifie si c’est un équipement
		isEquip := GetEquipmentByName(item.Name) != nil
		if isEquip {
			equip := GetEquipmentByName(item.Name)
			if equip != nil {
				EquipItem(c, equip)
				fmt.Println("🛡️ Équipé et retiré de l’inventaire :", equip.Nom)
			}
		} else {
			// Ajoute dans l’inventaire normal
			c.Inventaire = append(c.Inventaire, item.Name)
			fmt.Println("📦 Ajouté à l’inventaire :", item.Name)
		}
	}
}

// Conversion string -> int
func parseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
