package main

import (
	"bufio"
	"fmt"
	"time"
	"math/rand"
)

type Floor struct {
	Numero  int
	Monstre Monster
	IsBoss  bool
}

func InitFloor(floorNumber int) Floor {
	monstre := GenerateMonster(floorNumber)
	isBoss := floorNumber%10 == 0

	fmt.Printf("\n--- Étage %d ---\n", floorNumber)
	if isBoss {
		fmt.Println("⚔️ Boss de l'étage ! Préparez-vous !")
	} else {
		fmt.Printf("Monstre rencontré : %s\n", monstre.Nom)
	}

	return Floor{
		Numero:  floorNumber,
		Monstre: monstre,
		IsBoss:  isBoss,
	}
}

func TraverseFloors(player *Character, maxFloor int, scanner *bufio.Scanner) {
	currentFloor := 1
	for currentFloor <= maxFloor {
		floor := InitFloor(currentFloor)
		StartFloorCombat(player, floor.Numero, scanner)

		if player.PVActuels <= 0 {
			fmt.Println("💀 Vous êtes mort. Fin du jeu.")
			return
		}

		// Récompenses après combat
		fmt.Println("🏆 Combat terminé ! Récompenses et repos...")
		AfterCombat(player)

		// Demander si le joueur veut continuer ou quitter la tour
		fmt.Println("Voulez-vous continuer à monter ? (O/N)")
		scanner.Scan()
		choice := scanner.Text()
		if choice == "N" || choice == "n" {
			fmt.Println("🛑 Vous quittez la tour pour l'instant.")
			break
		}

		currentFloor++
	}
	if currentFloor > maxFloor {
		fmt.Println("🎉 Vous avez atteint le dernier étage ! Félicitations !")
	}
}


func AfterCombat(player *Character) {
	player.PVActuels += 20
	if player.PVActuels > player.PVMax {
		player.PVActuels = player.PVMax
	}
	player.Argent += 3
	fmt.Printf("💖 PV restaurés : %d/%d | 💰 Argent : %d\n", player.PVActuels, player.PVMax, player.Argent)
}


func randomMonster(monsters []Monster) Monster {
	return monsters[SeedRand(len(monsters))]
}

func SeedRand(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
