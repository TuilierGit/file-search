package readfiles

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FolderReader struct {
	jsonPath string
	files    []File
}

type File struct {
	Id      int    `json:"id"`
	Path    string `json:"path"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (fr FolderReader) printFiles() {
	fmt.Println("\nLecture des fichiers")
	for _, file := range fr.files {
		fmt.Println(file.Name)
	}
}

func (fr *FolderReader) getFiles() {
	jsonFile, err := os.Open(fr.jsonPath)
	if err != nil {
		fmt.Println("ERROR: Le fichier de configuration n'existe pas !!")
		return
	}
	defer jsonFile.Close()

	bytes, readErr := io.ReadAll(jsonFile)
	if readErr != nil {
		fmt.Println("ERROR: Impossibilité de lire le fichier")
		return
	}

	var files struct {
		Files []File `json:"files"`
	}
	parsingErr := json.Unmarshal(bytes, &files)

	if parsingErr != nil {
		fmt.Println("ERROR: Le fichier JSON n'est pas de la bonne forme")
		return
	}

	fr.files = files.Files

}

func (fr *FolderReader) init() {
	fr.getFiles()
}

func createFolderReader(folderPath string) FolderReader {
	var fr FolderReader
	fr.files = []File{}
	fr.jsonPath = folderPath + "/.folderReaderConfig.json"
	return fr
}
