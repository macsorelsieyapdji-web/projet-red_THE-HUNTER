package main

import "math/rand"

// Structure d'un monstre
type Monster struct {
    Nom        string
    PVMax      int
    PVActuels  int
    Degats     int
    Initiative int
}

// ----------------- Monstres normaux -----------------
func InitGoblin() Monster {
    return Monster{"Gobelin", 40, 40, 5, 5}
}

func InitChimeraAnt() Monster {
    return Monster{"Fourmilier", 80, 80, 12, 8}
}

func InitAntGuard() Monster {
    return Monster{"Ant de garde", 120, 120, 20, 12}
}

func InitNenIllusion() Monster {
    return Monster{"Illusion de Nen", 60, 60, 15, 15}
}

func InitPhantomTroupeMember() Monster {
    return Monster{"Membre de la Brigade Fantôme", 70, 70, 18, 12}
}

func InitHunter() Monster {
    return Monster{"Hunter novice", 50, 50, 10, 10}
}

// ----------------- Boss (tous les 10 étages) -----------------
func InitMiniBoss(name string) Monster {
    switch name {
    case "Meruem":
        return Monster{"Meruem", 250, 250, 35, 20}
    case "Chrollo Lucilfer":
        return Monster{"Chrollo Lucilfer", 180, 180, 30, 18}
    case "Neferpitou":
        return Monster{"Neferpitou", 200, 200, 32, 22}
    case "Hisoka":
        return Monster{"Hisoka", 190, 190, 28, 18}
    case "Menthuthuyoupi":
        return Monster{"Menthuthuyoupi", 220, 220, 33, 20}
    default:
        return Monster{"Boss inconnu", 150, 150, 25, 10}
    }
}

// Renvoie un monstre aléatoire parmi une liste
func RandomMonster(monsters []Monster) Monster {
    if len(monsters) == 0 {
        return InitGoblin()
    }
    return monsters[rand.Intn(len(monsters))]
}

// Génère un monstre selon l'étage
func GenerateMonster(floorNumber int) Monster {
    if floorNumber%10 == 0 {
        // Liste des bosses possibles
        bosses := []string{"Meruem", "Chrollo Lucilfer", "Neferpitou", "Hisoka", "Menthuthuyoupi"}
        name := bosses[rand.Intn(len(bosses))]
        return InitMiniBoss(name)
    }

    switch {
    case floorNumber <= 3:
        return RandomMonster([]Monster{InitGoblin(), InitHunter()})
    case floorNumber <= 7:
        return RandomMonster([]Monster{InitChimeraAnt(), InitNenIllusion()})
    case floorNumber <= 15:
        return RandomMonster([]Monster{InitAntGuard(), InitNenIllusion(), InitPhantomTroupeMember()})
    default:
        return RandomMonster([]Monster{InitChimeraAnt(), InitAntGuard(), InitNenIllusion(), InitPhantomTroupeMember()})
    }
}
