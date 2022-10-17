package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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

func choix_personnage() {
	fichier, _ := ioutil.ReadFile("hangman/words_1.txt") // peut varier selon le chemin du fichier. Récupère le fichier Hangman.txt à l'aide de ioutil et l'assigne à la variable fichier, "_" permet de ne pas récupérer l'erreur
	str := string(fichier)                               // transforme la variable fichier de type []byte en chaine de caractère et l'assigne à str
	//
	fichier2, _ := ioutil.ReadFile("hangman/pos_hangman.txt")
	liste_des_positions := string(fichier2)
	fmt.Println(str)
	fmt.Println(liste_des_positions)
	//
	fmt.Println("Tapez le nom de votre personnage : ")
	var nom string
	fmt.Scan(&nom)
	//var personnage HangManData
	//personnage.Init(nom,nouveau_mot())
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

}

func (h HangManData) perdu() {
	fmt.Println("Vous avez perdu !!!")
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

func chaque_mot(text string) []string {
	// transforme la chaine de caractère obtenue en lisant le fichier en tableau de string
	liste := []string{}
	mot := ""
	for _, element := range text { // parcours de la string text
		if element == 13 { // vérifie si l'élément est un retour à la ligne
			liste = append(liste, mot) // si c'est le cas ajoute le mot au tableau
			mot = ""                   // et réinitialise la variable mot
		} else { // si l'élément n'est pas un retour à la ligne
			mot += string(element) // l'élément est ajouté au mot en attendant de rencontrer un retour à la ligne
		}
	}
	if mot != "" { // vérifie si il reste quelque chose dans la variable mot
		liste = append(liste, mot) // si mot n'est pas vide, ajout de ce qui reste dans mot au tableau
	}
	return liste
}
