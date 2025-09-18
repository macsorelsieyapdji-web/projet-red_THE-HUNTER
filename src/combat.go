package main

import (
	"bufio"
	"fmt"
)

func StartFloorCombat(player *Character, floorNumber int, scanner *bufio.Scanner) {
	monstre := GenerateMonster(floorNumber)
	fmt.Printf("\n⚔️ Vous rencontrez : %s (%d PV)\n", monstre.Nom, monstre.PVActuels)

	tour := 1
	for player.PVActuels > 0 && monstre.PVActuels > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		PlayerTurn(player, &monstre, scanner)
		if monstre.PVActuels <= 0 {
			fmt.Println("🎉 Vous avez vaincu", monstre.Nom, "!")
			gagnerXP(player, 10+floorNumber*2)
			break
		}

		MonsterTurn(player, &monstre, tour)
		if player.PVActuels <= 0 {
			fmt.Println("💀 Vous êtes mort.")
			return
		}
		tour++
	}
}

func PlayerTurn(player *Character, monstre *Monster, scanner *bufio.Scanner) {
	fmt.Println("\n💡 Choisis ta compétence :")
	for i, skill := range player.Skills {
		fmt.Printf("%d. %s\n", i+1, skill)
	}
	fmt.Println("0. Utiliser un item")
	fmt.Print("> ")

	scanner.Scan()
	choix := scanner.Text()
	if choix == "0" {
		openInventory(player, scanner)
		return
	}

	index := parseInt(choix) - 1
	if index < 0 || index >= len(player.Skills) {
		fmt.Println("❌ Choix invalide ! Tu perds ton tour.")
		return
	}

	skill := player.Skills[index]
	degats := CalculateSkillDamage(skill, player)
	monstre.PVActuels -= degats
	if monstre.PVActuels < 0 {
		monstre.PVActuels = 0
	}

	fmt.Printf("👊 %s utilise %s et inflige %d dégâts à %s !\n", player.Nom, skill, degats, monstre.Nom)
	fmt.Printf("%s PV : %d/%d\n", monstre.Nom, monstre.PVActuels, monstre.PVMax)
}

func CalculateSkillDamage(skill string, player *Character) int {
	switch skill {
	case "Coup de poing":
		return player.Force + 5
	case "Boule de feu":
		return player.Intelligence*3 + 10
	case "Uppercut renforcé":
		return player.Force*2 + 5
	case "Projectile Nen":
		return player.Intelligence*2 + 8
	case "Contrôle mental":
		return player.Intelligence*3 + 5
	case "Poing électrique":
		return player.Force*2 + 8
	case "Épée invoquée":
		return player.Force + 15
	default:
		return 5
	}
}

func MonsterTurn(player *Character, monstre *Monster, turn int) {
	degats := monstre.Degats
	if turn%3 == 0 {
		degats *= 2
		fmt.Printf("⚡ %s utilise un coup spécial !\n", monstre.Nom)
	}
	player.PVActuels -= degats
	if player.PVActuels < 0 {
		player.PVActuels = 0
	}
	fmt.Printf("👹 %s inflige %d dégâts à %s\n", monstre.Nom, degats, player.Nom)
	fmt.Printf("%s PV : %d/%d\n", player.Nom, player.PVActuels, player.PVMax)
}

func gagnerXP(player *Character, xp int) {
	player.Experience += xp
	fmt.Printf("⭐ Tu gagnes %d XP ! (%d/%d)\n", xp, player.Experience, player.ExpToLevel)
	for player.Experience >= player.ExpToLevel {
		player.Experience -= player.ExpToLevel
		player.Niveau++
		player.PVMax += 10
		player.Force += 2
		player.Intelligence += 2
		player.PVActuels = player.PVMax
		player.ExpToLevel += 10
		fmt.Println("🎉 Niveau supérieur ! Tu es maintenant niveau", player.Niveau)
	}
}
