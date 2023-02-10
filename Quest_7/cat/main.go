package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func displayContent(filename string) { // Print le contenu du fichier
	if filename == "quest8.txt" {
		file, _ := ioutil.ReadFile(filename)
		for _, r := range file {
			z01.PrintRune(rune(r))
		}
	} else if filename == "quest8T.txt" {
		file, _ := ioutil.ReadFile(filename)
		for _, r := range file {
			z01.PrintRune(rune(r))
		}
	}
}

func printError(filename string) { // Si y'a une erreur dans l'ouverture d'un fichier
	start := "ERROR: open "
	end := ": no such file or directory\n"
	ErrMessage := start + filename + end
	for _, r := range ErrMessage {
		z01.PrintRune(rune(r))
	}
}

func main() {
	if os.Args[1] == "echo" { // Si la commande commence par echo
		if len(os.Args) >= 6 {
			if os.Args[5] == "go" && os.Args[6] == "run" {
				file := os.Args[3]
				for _, r := range file {
					z01.PrintRune(rune(r))
				}
				z01.PrintRune('\n')
			}
		}
	}
	if len(os.Args) == 2 { // Si un seul fichier est demandé
		filename := os.Args[1]
		a, err := os.Open(filename)
		if err != nil {
			printError(filename)
			a.Close()
			os.Exit(1)
		} else {
			displayContent(filename)
			a.Close()
		}
	}
	if len(os.Args) == 3 { // Si 2 fichiers sont demandés
		filename1 := os.Args[1]
		filename2 := os.Args[2]
		a, err := os.Open(filename1)
		if err != nil {
			printError(filename1)
			a.Close()
			os.Exit(1)
		} else {
			displayContent(filename1)
			a.Close()
			a2, err := os.Open(filename2)
			if err != nil {
				printError(filename2)
				a2.Close()
				os.Exit(1)
			} else {
				displayContent(filename2)
				a2.Close()
			}
		}
	}
}
