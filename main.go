package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nexidian/gocliselect"
)

func main() {
	menu1 := gocliselect.NewMenu("Quel format voulez-vous générer ?")
	menu1.AddItem("PRENOM.NOM (ex : JEAN.DUPONT)", "prenom_nom")
	menu1.AddItem("NOM.PRENOM (ex : DUPONT.JEAN)", "nom_prenom")
	menu1.AddItem("P.NOM (ex : J.DUPONT)", "p_nom")
	menu1.AddItem("PNOM (ex : JDUPONT)", "pnom")
	menu1.AddItem("NOMP (ex : DUPONTJ)", "nomp")
	choice := menu1.Display()

	switch choice {
	case "prenom_nom":
		menu2 := gocliselect.NewMenu("Quel fichier voulez-vous utiliser ?")
		menu2.AddItem("noms.csv", "noms.csv")
		menu2.AddItem("noms_big.csv", "noms_big.csv")
		menu2.AddItem("noms_huge.csv", "noms_huge.csv")
		choice2 := menu2.Display()
		switch choice2 {
		case "noms.csv":
			generate_prenom_point_nom("prenoms.csv", "noms.csv")
		case "noms_big.csv":
			generate_prenom_point_nom("prenoms.csv", "noms_big.csv")
		case "noms_huge.csv":
			generate_prenom_point_nom("prenoms.csv", "noms_huge.csv")
		}
	case "nom_prenom":
		menu2 := gocliselect.NewMenu("Quel fichier voulez-vous utiliser ?")
		menu2.AddItem("noms.csv", "noms.csv")
		menu2.AddItem("noms_big.csv", "noms_big.csv")
		menu2.AddItem("noms_huge.csv", "noms_huge.csv")
		choice2 := menu2.Display()
		switch choice2 {
		case "noms.csv":
			generate_nom_point_prenom("prenoms.csv", "noms.csv")
		case "noms_big.csv":
			generate_nom_point_prenom("prenoms.csv", "noms_big.csv")
		case "noms_huge.csv":
			generate_nom_point_prenom("prenoms.csv", "noms_huge.csv")
		}
	case "p_nom":
		menu2 := gocliselect.NewMenu("Quel fichier voulez-vous utiliser ?")
		menu2.AddItem("noms.csv", "noms.csv")
		menu2.AddItem("noms_big.csv", "noms_big.csv")
		menu2.AddItem("noms_huge.csv", "noms_huge.csv")
		choice2 := menu2.Display()
		switch choice2 {
		case "noms.csv":
			generate_1e_lettre_point_nom("prenoms.csv", "noms.csv")
		case "noms_big.csv":
			generate_1e_lettre_point_nom("prenoms.csv", "noms_big.csv")
		case "noms_huge.csv":
			generate_1e_lettre_point_nom("prenoms.csv", "noms_huge.csv")
		}
	case "pnom":
		menu2 := gocliselect.NewMenu("Quel fichier voulez-vous utiliser ?")
		menu2.AddItem("noms.csv", "noms.csv")
		menu2.AddItem("noms_big.csv", "noms_big.csv")
		menu2.AddItem("noms_huge.csv", "noms_huge.csv")
		choice2 := menu2.Display()
		switch choice2 {
		case "noms.csv":
			generate_1e_lettre_nom("prenoms.csv", "noms.csv")
		case "noms_big.csv":
			generate_1e_lettre_nom("prenoms.csv", "noms_big.csv")
		case "noms_huge.csv":
			generate_1e_lettre_nom("prenoms.csv", "noms_huge.csv")
		}
	case "nomp":
		menu2 := gocliselect.NewMenu("Quel fichier voulez-vous utiliser ?")
		menu2.AddItem("noms.csv", "noms.csv")
		menu2.AddItem("noms_big.csv", "noms_big.csv")
		menu2.AddItem("noms_huge.csv", "noms_huge.csv")
		choice2 := menu2.Display()
		switch choice2 {
		case "noms.csv":
			generate_nom_1e_lettre("prenoms.csv", "noms.csv")
		case "noms_big.csv":
			generate_nom_1e_lettre("prenoms.csv", "noms_big.csv")
		case "noms_huge.csv":
			generate_nom_1e_lettre("prenoms.csv", "noms_huge.csv")
		}
	}
	deleteDuplicateLines("usernames.txt")
}

func generate_prenom_point_nom(filename1 string, filename2 string) {
	firstNames, err := readLines(filename1)
	if err != nil {
		fmt.Println("Error reading first names:", err)
		return
	}

	lastNames, err := readLines(filename2)
	if err != nil {
		fmt.Println("Error reading last names:", err)
		return
	}

	file, err := os.Create("usernames.txt")
	if err != nil {
		fmt.Println("Error creating usernames file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, firstName := range firstNames {
		for _, lastName := range lastNames {
			username := fmt.Sprintf("%s%s%s", firstName, ".", lastName)
			_, err := writer.WriteString(username + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}
	writer.Flush()
	fmt.Println("Usernames successfully generated in usernames.txt")
}

func generate_nom_point_prenom(filename1 string, filename2 string) {
	firstNames, err := readLines(filename1)
	if err != nil {
		fmt.Println("Error reading first names:", err)
		return
	}

	lastNames, err := readLines(filename2)
	if err != nil {
		fmt.Println("Error reading last names:", err)
		return
	}

	file, err := os.Create("usernames.txt")
	if err != nil {
		fmt.Println("Error creating usernames file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, firstName := range firstNames {
		for _, lastName := range lastNames {
			username := fmt.Sprintf("%s%s%s", lastName, ".", firstName)
			_, err := writer.WriteString(username + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}
	writer.Flush()
	fmt.Println("Usernames successfully generated in usernames.txt")
}

func generate_1e_lettre_point_nom(filename1 string, filename2 string) {
	firstNames, err := readLines(filename1)
	if err != nil {
		fmt.Println("Error reading first names:", err)
		return
	}

	lastNames, err := readLines(filename2)
	if err != nil {
		fmt.Println("Error reading last names:", err)
		return
	}

	file, err := os.Create("usernames.txt")
	if err != nil {
		fmt.Println("Error creating usernames file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, firstName := range firstNames {
		for _, lastName := range lastNames {
			username := fmt.Sprintf("%s%s%s", firstName[0:1], ".", lastName)
			_, err := writer.WriteString(username + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}
	writer.Flush()
	fmt.Println("Usernames successfully generated in usernames.txt")
}

func generate_1e_lettre_nom(filename1 string, filename2 string) {
	firstNames, err := readLines(filename1)
	if err != nil {
		fmt.Println("Error reading first names:", err)
		return
	}

	lastNames, err := readLines(filename2)
	if err != nil {
		fmt.Println("Error reading last names:", err)
		return
	}

	file, err := os.Create("usernames.txt")
	if err != nil {
		fmt.Println("Error creating usernames file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, firstName := range firstNames {
		for _, lastName := range lastNames {
			username := fmt.Sprintf("%s%s", firstName[0:1], lastName)
			_, err := writer.WriteString(username + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}
	writer.Flush()
	fmt.Println("Usernames successfully generated in usernames.txt")
}

func generate_nom_1e_lettre(filename1 string, filename2 string) {
	firstNames, err := readLines(filename1)
	if err != nil {
		fmt.Println("Error reading first names:", err)
		return
	}

	lastNames, err := readLines(filename2)
	if err != nil {
		fmt.Println("Error reading last names:", err)
		return
	}

	file, err := os.Create("usernames.txt")
	if err != nil {
		fmt.Println("Error creating usernames file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, firstName := range firstNames {
		for _, lastName := range lastNames {
			username := fmt.Sprintf("%s%s", lastName, firstName[0:1])
			_, err := writer.WriteString(username + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}
	writer.Flush()
	fmt.Println("Usernames successfully generated in usernames.txt")
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func deleteDuplicateLines(filename string) error {
	lines, err := readLines(filename)
	if err != nil {
		return err
	}

	seen := make(map[string]bool)
	var uniqueLines []string

	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range uniqueLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}
