# THE HUNTER - Tour du Ciel

🎮 Jeu terminal en Go inspiré de l’univers **Hunter x Hunter**. Le joueur incarne un Hunter qui gravit la **Tour du Ciel**, combat des monstres, améliore ses compétences et son équipement pour atteindre le sommet.

---

## Table des matières

1. [Présentation](#présentation)  
2. [Installation](#installation)  
3. [Lancement du jeu](#lancement-du-jeu)  
4. [Gameplay](#gameplay)  
5. [Mécaniques de jeu](#mécaniques-de-jeu)  
6. [Fichiers du projet](#fichiers-du-projet)  
7. [Contributions](#contributions)

---

## Présentation

Le joueur :  
- Crée son **Hunter** avec un nom et une **classe de Nen** (Renforcement, Émission, Transformation, etc.)  
- Progresse étage par étage dans la **Tour du Ciel**  
- Combat des monstres et des **boss tous les 10 étages**  
- Gagne de l’XP pour monter de niveau et améliorer ses statistiques  
- Peut acheter et équiper des **armes et armures** dans le **shop**  
- Gère son **inventaire** et ses **compétences Nen**  

---

## Installation

1. Assurez-vous que [Go](https://golang.org/dl/) est installé sur votre machine (Go 1.20+ recommandé).  
2. Clonez le projet :  

```bash
git clone https://github.com/ton-utilisateur/the-hunter.git
cd the-hunter

Gameplay

Inventaire : contient potions, objets spéciaux et équipements.

Équipement : armes et armures avec bonus. Lorsqu’un équipement est utilisé, il disparaît de l’inventaire et s’applique sur le personnage.

Combat tour par tour :

Choix de compétences pour attaquer

Utilisation d’items

Dégâts calculés selon stats + compétences

Monstres :

Normaux : Gobelins, Fourmiliers, Ants de garde, Illusions de Nen

Boss : tous les 10 étages, nom et stats modifiables

Récompenses : XP et argent après chaque combat

Option quitter la tour : permet d’interrompre la progression et retourner au menu principal

Mécaniques de jeu

Stats du personnage : PV, Force, Intelligence, Initiative, Argent, XP, Niveau

XP et leveling :

XP par combat : ajustable dans StartFloorCombat

Progression par niveau : stats augmentées automatiquement

Compétences :

Chaque compétence inflige des dégâts différents

Bonus selon classe Nen et équipement

Équipement :

Arme : augmente la Force

Armure / Torse / Tête / Pieds : augmentent PV ou stats

Les bonus sont appliqués immédiatement

Fichiers du projet
Fichier	Description
main.go	Point d’entrée du jeu, menu principal
character.go	Création et gestion du personnage, affichage des stats
inventory.go	Gestion de l’inventaire et utilisation d’objets
items.go	Définition des objets et potions
equipment.go	Gestion de l’équipement et des bonus
shop.go	Boutique, achat et équiper les items
monster.go	Définition des monstres et boss
combat.go	Logique des combats tour par tour, compétences
floor.go	Progression dans la Tour, génération des étages et monstres
Contribution

Vous pouvez ajouter de nouveaux monstres, équipements, compétences ou mécaniques.

Respectez la structure des fichiers et la cohérence des stats.

Toutes les contributions sont les bienvenues via pull requests.

Exemple de progression

Créer un Hunter “Gon” avec classe Renforcement

Monter 3 étages : affronter Gobelins et Fourmiliers

Acheter un équipement au shop : Poings Renforcés

Équiper l’objet et observer l’augmentation de Force

Continuer la tour jusqu’au boss du 10ᵉ étage