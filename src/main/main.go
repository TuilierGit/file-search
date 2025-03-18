package main

import (
	"bufio"
	"file-search/src/readfiles"
	"fmt"
	"os"
	"strings"
)

func getPath(opt string) string {

	path := ""

	if opt == "terminal" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Chemin du dossier : ")
		folderPath, err := reader.ReadString('\n')

		if err != nil {
			println("ERROR: échec de la lecture de l'option")
		}
		path = strings.TrimSpace(folderPath)
		fmt.Printf("Le chemin du dossier est : %v\n", path)

	} else {
		lastErr := fmt.Errorf("ERROR: L'option de lecture n'est pas géré par l'outils")
		panic(lastErr)
	}

	return path
}

func main() {
	readfiles.GetFileList(getPath("terminal"))
}
