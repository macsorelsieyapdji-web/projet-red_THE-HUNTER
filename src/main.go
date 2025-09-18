package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(`🎮 Bienvenue dans THE HUNTER - Tour du Ciel 🗼
------------ ----    ---- ------------      ----    ---- ----    ---- ----    ---- ------------ ------------ -----------       
************ ****    **** ************      ****    **** ****    **** *****   **** ************ ************ ***********       
------------ ----    ---- ----              ----    ---- ----    ---- ------  ---- ------------ ----         ----    ---       
****         ************ ************      ************ ****    **** ************     ****     ************ *********         
----         ------------ ------------      ------------ ----    ---- ------------     ----     ------------ ---------         
****         ****    **** ****              ****    **** ************ ****  ******     ****     ****         ****  ****        
----         ----    ---- ------------      ----    ---- ------------ ----   -----     ----     ------------ ----   ----       
****         ****    **** ************      ****    **** ************ ****    ****     ****     ************ ****    ****      
`)

	player := createCharacter(scanner)

	for {
		fmt.Println("\n--- MENU PRINCIPAL ---")
		fmt.Println("1. Afficher les infos du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Monter dans la Tour du Ciel")
		fmt.Println("4. Visiter le marchand")
		fmt.Println("5. Quitter le jeu")
		fmt.Print("Choix : ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			displayCharacterInfo(player)
		case "2":
			openInventory(&player, scanner)
		case "3":
			fmt.Print("Combien d'étages veux-tu tenter ? (0 pour quitter la Tour) : ")
			scanner.Scan()
			floorChoice := parseInt(scanner.Text())
			if floorChoice < 0 {
				fmt.Println("❌ Choix invalide.")
				continue
			}
			if floorChoice == 0 {
				fmt.Println("⛔ Tu as choisi de quitter la Tour.")
				continue
			}
			TraverseFloors(&player, floorChoice, scanner)
		case "4":
			OpenShop(&player, scanner)
		case "5":
			fmt.Println("👋 Merci d’avoir joué à The Hunter !")
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}
