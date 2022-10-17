package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	//

	//
	choix_personnage()
	//menu()

}

//_______________________________________________________________________________________________________________________________________

func menu() {
	fmt.Println("Bienvenue dans notre jeu hangman.\nQue souhaitez vous faire ?")
	fmt.Println("1 : lancer une nouvelle partie")
	fmt.Println("2 : voir les règles")
	fmt.Println("3 : arrêter le terminal")
	var reponse int
	fmt.Scan(&reponse)
	switch reponse {
	case 1:
		choix_personnage()
	case 2:
		affichage_regle()
	}
}
func affichage_regle() {
	fmt.Println("Le jeu du Hangman consiste à trouver un mot choisit aléatoirement par l'ordinateur.\nVous avez 10 tentatives pour trouver le mot.\nSi vous trouvez le mot avant d'avoir utilisé toutes vos tentatives, vous gagnez.\nSi vous n'avez plus de tentatives, vous perdez.\nBonne chance !!!")
}

// _______________________________________________________________________________________________________________________________________
func transforme_en_liste(fichier *os.File) []string {
	var liste []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		liste = append(liste, scanner.Text())
	}
	return liste
}

func liste_position(fichier *os.File) []string {
	var liste []string
	scanner := bufio.NewScanner(fichier)
	var stockage string
	var compteur int
	for scanner.Scan() {
		compteur++
		stockage += string(scanner.Text()) + "\n"
		if compteur == 8 {
			liste = append(liste, stockage)
			compteur = 0
			stockage = ""
		}

	}
	return liste
}
func choix_personnage() {
	file1, err1 := os.Open("words_1.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	file1_liste := transforme_en_liste(file1)
	file1.Close()
	//
	file2, err2 := os.Open("pos_hangman.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	liste_des_positions := liste_position(file2)
	file2.Close()
	//
	fmt.Println("Tapez le nom de votre personnage : ")
	var nom string
	fmt.Scan(&nom)
	var personnage HangManData
	personnage.Init(nom, nouveau_mot(file1_liste), 10, liste_des_positions)
	lancement_jeu(personnage)
}

func lancement_jeu(h HangManData) {

	for h.Attempts > -1 {
		if h.Attempts == 0 {
			h.perdu()
			break
		} else {
			h.jouer_tour()

		}
	}

}

func (h *HangManData) jouer_tour() {
	h.Attempts -= 1
	fmt.Println("Quelle lettre voulez vous essayer ?")
	var lettre string
	fmt.Scan(&lettre)
	h.AjoutLettre(lettre)
	if h.verifletter(lettre) && !h.DejaDansNom(lettre) {
		h.remplace(lettre)
	} else {
		fmt.Println("vous vous êtes trompez")
		// func affichez hangman +1
	}
	h.Victoire()
}

func (h HangManData) perdu() {
	fmt.Println("\nVous avez perdu !!!")
	time.Sleep(3 * time.Second)
}

//_________________________________________________________________________________________________________________________________________

type HangManData struct {
	nom              string   // Name of the Hangman
	Word             string   // Word composed of '_', ex: H_ll_
	ToFind           string   // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int      // Number of attempts left
	HangmanPositions []string // It can be the array where the positions parsed in "pos_hangman.txt" are stored
	ActualPosition   string
	UsedLetter       []string
}

func (h *HangManData) Init(nom string, a_trouver string, tentatives int, liste_pose []string) {
	h.nom = nom
	h.ToFind = a_trouver
	h.Attempts = tentatives
	h.HangmanPositions = liste_pose
	h.ActualPosition = liste_pose[0]
}

func nouveau_mot(fichier []string) string {
	random := random(len(fichier))
	mot := fichier[random]
	return mot
}
func random(i int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(i)
	return random
}

//_________________________________________________________________________________________________________________________________________

func (h *HangManData) verifletter(letter string) bool {
	h.DejaDansNom(letter)
	for _, i := range h.ToFind {
		if letter == string(i) {
			return true
		}
	}
	return false
}

func (h *HangManData) DejaDansNom(letter string) bool {
	retour := false
	for _, i := range h.Word {
		if string(i) == letter {
			retour = true
			break
		}
	}
	return retour
}

func (h *HangManData) remplace(lettre string) {
	var nouveau_mot string
	for i := 0; i < len(h.ToFind); i++ {
		if string(h.ToFind[i]) == lettre {
			nouveau_mot += lettre
		} else {
			nouveau_mot += h.Word
		}
	}
	h.Word = nouveau_mot
}

func (h *HangManData) LettreUtilise(letter string) bool {

	for _, i := range h.UsedLetter {
		if i == letter {

			return true

		}
		break
	}
	return false
}

func (h *HangManData) AjoutLettre(letter string) {
	if h.LettreUtilise(letter) == true {
		fmt.Println("cette lettre à déja été utilisée")

	} else if h.LettreUtilise(letter) == false {
		h.UsedLetter = append(h.UsedLetter, letter)
	}
}

func (h *HangManData) Victoire() {
	if h.ToFind == h.Word {
		fmt.Println("Vous avez trouvé le mot, félicitations.")
	}
}
