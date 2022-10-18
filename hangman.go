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
	mot := nouveau_mot(file1_liste)
	personnage.Init(nom, mot, word_with_blank(mot), 10, liste_des_positions)
	lancement_jeu(personnage)
}

func lancement_jeu(h HangManData) {

	for h.Attempts > -1 {
		if h.Attempts == 0 {
			h.perdu()
			break
		} else if h.Word == h.ToFind {
			h.Victoire()
		} else {
			h.jouer_tour()

		}
	}

}

func (h *HangManData) jouer_tour() {

	fmt.Println("Il vous reste", h.Attempts, "tentatives.")
	fmt.Println("Voici la pose actuelle de", h.nom, ":\n ")
	fmt.Println(h.ActualPosition)
	fmt.Println("Voici les lettres que vous avez déja essayées :", h.UsedLetter)
	fmt.Println("Voici ce que vous avez trouvé du mot :", h.Word)
	fmt.Println("\nQuelle lettre voulez vous essayer ?")
	var lettre string
	fmt.Scan(&lettre)
	h.AjoutLettre(lettre)
	if h.verifletter(lettre) && !h.DejaDansNom(lettre) {
		fmt.Println(" \nCette lettre fait bien partie du mot, bravo !")
		h.remplace(lettre)
	} else {
		fmt.Println("Vous vous êtes trompé.")
		h.Attempts -= 1
		h.ActualPosition = h.HangmanPositions[10-h.Attempts]
	}
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

func (h *HangManData) Init(nom string, a_trouver string, mot_actuel string, tentatives int, liste_pose []string) {
	h.nom = nom
	h.ToFind = a_trouver
	h.Word = mot_actuel
	h.Attempts = tentatives
	h.HangmanPositions = liste_pose
	h.ActualPosition = liste_pose[0]
}
func word_with_blank(mot string) string {
	var liste []string
	for _, element := range mot {
		liste = append(liste, string(element))
	}
	n := len(mot)/2 + 1
	var nouveau_mot string
	i := 0
	for i < len(mot)-n {
		index := random(len(mot) - 1)
		if liste[index] != "_" {
			liste[index] = "_"
			i++
		}
	}
	for _, element := range liste {
		nouveau_mot += string(element)
	}
	return nouveau_mot
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
		if letter == string(i) || letter == string(i-32) {
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
		if string(h.ToFind[i]) == lettre && !h.DejaDansNom(lettre) {
			nouveau_mot += lettre
		} else {
			nouveau_mot += string(h.Word[i])
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
		fmt.Println("Cette lettre à déja été utilisée.")

	} else if h.LettreUtilise(letter) == false {
		h.UsedLetter = append(h.UsedLetter, letter)
	}
}

func (h *HangManData) Victoire() {
	if h.ToFind == h.Word {
		fmt.Printf("\nVous avez trouvé le mot %s, félicitations.", h.ToFind)
		fmt.Println("Que souhaitez vous faire ?\n1 : relancer une partie\n2 : Quitter")
		var rep int
		fmt.Scan(&rep)
		switch rep {
		case 1:
			choix_personnage()
		case 2:
			Quit()
		}
	}
}

func Quit() {
	os.Exit(0)
}
