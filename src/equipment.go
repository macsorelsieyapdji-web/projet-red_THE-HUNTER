package main

import "fmt"

// Structure d'équipement
type Equipment struct {
	Nom      string
	Type     string // "arme", "armure", "tete", "torse", "pieds"
	BonusDmg int
	BonusPV  int
}

// Équipement actuel du personnage
type Equipped struct {
	Tete  *Equipment
	Torse *Equipment
	Pieds *Equipment
	Arme  *Equipment
}

// Liste globale des équipements
var Equipements = []Equipment{
	{"Poings Renforcés", "arme", 5, 0},
	{"Dague de Kurapika", "arme", 10, 0},
	{"Katana de Zushi", "arme", 15, 0},
	{"Gilet pare-chocs", "armure", 0, 10},
	{"Veste en Nen", "armure", 0, 20},
	{"Cape de l’Ombre", "armure", 0, 30},
	{"Chapeau de l’aventurier", "tete", 0, 10},
	{"Tunique de l’aventurier", "torse", 0, 25},
	{"Bottes de l’aventurier", "pieds", 0, 15},
}

// Renvoie un pointeur vers un équipement par nom
func GetEquipmentByName(nom string) *Equipment {
	for i := range Equipements {
		if Equipements[i].Nom == nom {
			return &Equipements[i]
		}
	}
	return nil
}

// Équipe un objet et le retire de l'inventaire
func EquipItem(c *Character, eq *Equipment) {
	if c.Équipement == nil {
		c.Équipement = &Equipped{}
	}

	// Retirer de l’inventaire
	for i, item := range c.Inventaire {
		if item == eq.Nom {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			break
		}
	}

	// Appliquer l’équipement
	switch eq.Type {
	case "arme":
		c.Équipement.Arme = eq
		fmt.Println("⚔️ Arme équipée :", eq.Nom)
	case "tete":
		c.Équipement.Tete = eq
		fmt.Println("🎩 Équipement de tête :", eq.Nom)
	case "torse", "armure":
		c.Équipement.Torse = eq
		fmt.Println("🛡️ Torse équipé :", eq.Nom)
	case "pieds":
		c.Équipement.Pieds = eq
		fmt.Println("👢 Pieds équipés :", eq.Nom)
	}

	// Appliquer les bonus
	c.Force += eq.BonusDmg
	c.PVMax += eq.BonusPV
	c.PVActuels += eq.BonusPV
}

// Déséquipe un objet et le remet dans l’inventaire
func UnequipItem(c *Character, eq *Equipment) {
	// Retirer les bonus
	c.Force -= eq.BonusDmg
	c.PVMax -= eq.BonusPV
	if c.PVActuels > c.PVMax {
		c.PVActuels = c.PVMax
	}

	// Retirer de l’équipement
	switch eq.Type {
	case "arme":
		if c.Équipement.Arme != nil && c.Équipement.Arme.Nom == eq.Nom {
			c.Équipement.Arme = nil
		}
	case "tete":
		if c.Équipement.Tete != nil && c.Équipement.Tete.Nom == eq.Nom {
			c.Équipement.Tete = nil
		}
	case "torse", "armure":
		if c.Équipement.Torse != nil && c.Équipement.Torse.Nom == eq.Nom {
			c.Équipement.Torse = nil
		}
	case "pieds":
		if c.Équipement.Pieds != nil && c.Équipement.Pieds.Nom == eq.Nom {
			c.Équipement.Pieds = nil
		}
	}

	// Remettre dans l’inventaire
	c.Inventaire = append(c.Inventaire, eq.Nom)
	fmt.Println("❌ Déséquipé :", eq.Nom, "-> retourné à l’inventaire")
}
