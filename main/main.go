package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Fonction qui limites les espaces entre les mots
func Spaces(s string) string {
	a := false
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= 33 && s[i] <= 126 {
			result = result + string(s[i])
			a = true
		} else if s[i] == ' ' && i == len(s)-1 {
		} else if a == true && s[i] == 32 {
			result = result + " "
			a = false
		}
	}
	return result
}

// Fonction qui gère les ponctuations
// fonction switch
func SpecialString(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += string(s[i])
		} else if s[i] == ',' || s[i] == '.' || s[i] == ':' || s[i] == ';' || s[i] == '?' || s[i] == '!' {
			if s[i+1] != ',' && s[i+1] != '.' && s[i+1] != ':' && s[i+1] != ';' && s[i+1] != '?' && s[i+1] != '!' && s[i+1] != ' ' && s[i+1] != '\'' {
				result = result + string(s[i]) + " "
			} else {
				result = result + string(s[i])
			}
		} else if s[i] == 32 && s[i+1] != ',' && s[i+1] != '.' && s[i+1] != ':' && s[i+1] != ';' && s[i+1] != '?' && s[i+1] != '!' {
			result += " "
		} else if s[i] >= 33 && s[i] <= 126 {
			result += string(s[i])
		}
	}
	return result
}

// Fonction qui rajoute les an
func AnOrA(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if (s[i] == 'a' || s[i] == 'A') && string(s[i+1]) == " " && (s[i+2] == 'a' || s[i+2] == 'e' || s[i+2] == 'i' || s[i+2] == 'o' || s[i+2] == 'u' || s[i+2] == 'h') {
			result = result + string(s[i]) + "n"
		} else {
			result = result + string(s[i])
		}
	}
	return result
}

// Fonction qui rajoute des majuscules sur la première lettre de chaque mot
func Capitalize(s string) string {
	var result string
	IsNewWord := true
	for _, l := range s {
		alph := (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') || (l >= '0' && l <= '9')
		if alph {
			if IsNewWord {
				if l >= 'a' && l <= 'z' {
					l = l + -32
				}
				IsNewWord = false
			} else {
				if l >= 'A' && l <= 'Z' {
					l = l + 32
				}
			}
		} else {
			IsNewWord = true
		}
		result += string(l)
	}
	return result
}

// Fonction qui converti les valeurs binaires en valeurs décimales
func BinToDec(s string) string {
	Dec, error := strconv.ParseInt(s, 2, 64)
	if error != nil {
		return "Convertion impossible"
	}
	return strconv.Itoa(int(Dec))
}

// Fonction qui converti les valeurs héxadécimales en valeurs décimales
func HexToDec(s string) string {
	Dec, error := strconv.ParseInt(s, 16, 64)
	if error != nil {
		return "Convertion impossible"
	}
	return strconv.Itoa(int(Dec))
}

func main() {
	//Vérification de la longueur des arguments (fichier, txt à modifier, txt final)
	if len(os.Args) != 3 {
		fmt.Println("Arguments non valides")
		os.Exit(0)
	}
	//os.O_CREATE crée le fichier s'il n'existe pas déjà
	file, err := os.OpenFile(os.Args[1], os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Création/Ouverture du fichier sample.txt impossible")
		os.Exit(1)
	}
	//Fermeture du fichier
	defer file.Close()
	//Stocker le contenu du fichier texte dans une variable qui sera en tableau de byte
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Lecture de fichier impossible")
		os.Exit(2)
	}
	//Convertir la variable de byte à string
	text := string(data)
	//Application des fonctions à ma chaine de charactères
	text = Spaces(text)
	text = AnOrA(text)
	//Convertir ma chaine de charactères en tableau de string
	tabString := strings.Split(text, " ")
	//Boucle qui vérifie si la chaine de charactères comprend des balises hex
	for i, word := range tabString {
		if strings.Contains(word, "(hex)") {
			tabString[i] = ""
			tabString[i-1] = HexToDec(tabString[i-1])
		}
	}
	//Boucle qui vérifie si la chaine de charactères comprend des balises bin
	for i, word := range tabString {
		if strings.Contains(word, "(bin)") {
			tabString[i] = ""
			tabString[i-1] = BinToDec(tabString[i-1])
		}
	}
	//Boucle qui vérifie si la chaine de charactères comprend des balises up
	for i, word := range tabString {
		if strings.Contains(word, "(up)") {
			tabString[i] = ""
			tabString[i-1] = strings.ToUpper(tabString[i-1])
		} else if strings.Contains(word, "(up,") {
			//Retirer la parenthèse devant le chiffre de la balise
			tabString[i+1] = tabString[i+1][:len(tabString[i+1])-1]
			//Convertion du chiffre string en int
			a, _ := strconv.Atoi(string(tabString[i+1]))
			for j := 0; j < a; j++ {
				tabString[i-1-j] = strings.ToUpper(tabString[i-1-j])
			}
			//Suppression de la balise
			tabString[i] = ""
			tabString[i+1] = ""
		}
	}
	//Boucle qui vérifie si la chaine de charactères comprend des balises low
	for i, word := range tabString {
		if strings.Contains(word, "(low)") {
			tabString[i] = ""
			tabString[i-1] = strings.ToLower(tabString[i-1])
		} else if strings.Contains(word, "(low,") {
			tabString[i+1] = tabString[i+1][:len(tabString[i+1])-1]
			a, _ := strconv.Atoi(string(tabString[i+1]))
			for j := 0; j < a; j++ {
				tabString[i-1-j] = strings.ToLower(tabString[i-1-j])
			}
			tabString[i] = ""
			tabString[i+1] = ""
		}
	}
	//Boucle qui vérifie si la chaine de charactères comprend des balises cap
	for i, word := range tabString {
		if strings.Contains(word, "(cap)") {
			tabString[i] = ""
			tabString[i-1] = Capitalize(tabString[i-1])
		} else if strings.Contains(word, "(cap,") {
			tabString[i+1] = tabString[i+1][:len(tabString[i+1])-1]
			a, _ := strconv.Atoi(string(tabString[i+1]))
			for j := 0; j < a; j++ {
				tabString[i-1-j] = Capitalize(tabString[i-1-j])
			}
			tabString[i] = ""
			tabString[i+1] = ""
		}
	}
	finalstring := ""
	verif := false
	for i := 0; i < len(tabString); i++ {
		if tabString[i] == "'" {
			if !verif {
				finalstring += " '"
				verif = true
			} else {
				finalstring += "' "
				verif = false
			}
		} else if i == len(tabString)-1 {
			finalstring += tabString[i]
		} else {
			if tabString[i+1] == "'" {
				finalstring += tabString[i]
			} else {
				finalstring += tabString[i] + " "
			}
		}
	}
	finalstring = Spaces(finalstring)
	finalstring = SpecialString(finalstring)
	finalstring = Spaces(finalstring)
	//Ouverture du fichier result pour pouvoir écrire le texte
	//Trunc permet d'écraser le contenue déjà existant dans le fichier result
	file2, err := os.OpenFile(os.Args[2], os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println("Création/Ouverture du fichier result.txt impossible")
		os.Exit(1)
	}
	//file.WriteString permet d'écrire dans le fichier
	_, err = file2.WriteString(finalstring)
	if err != nil {
		fmt.Println("Écriture de fichier impossible")
	}
}
