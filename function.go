package goreload

//Fonctions utiles: ToUpperSpecial et ToLowerSpecial et regex FindAllString

// Fonction qui limites les espaces entre les mots
func Spaces(s string) string {
	a := false
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= 33 && s[i] <= 126 {
			result = result + string(s[i])
			a = true
		} else if a == true && s[i] == 32 {
			result = result + " "
			a = false
		}
	}
	return result
}

// Fonction qui gère les ponctuations
func SpecialString(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += string(s[i])
		} else if s[i] == ',' || s[i] == '.' || s[i] == ':' || s[i] == ';' || s[i] == '?' || s[i] == '!' {
			if s[i+1] != ',' && s[i+1] != '.' && s[i+1] != ':' && s[i+1] != ';' && s[i+1] != '?' && s[i+1] != '!' {
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
		if s[i] == 'a' && string(s[i+1]) == " " && (s[i+2] == 'a' || s[i+2] == 'e' || s[i+2] == 'i' || s[i+2] == 'o' || s[i+2] == 'u' || s[i+2] == 'y' || s[i+2] == 'h') {
			result = result + string(s[i]) + "n"
		} else {
			result = result + string(s[i])
		}
	}
	return result
}

func Reverse(s string) string {
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
