package service

import (
	"encoding/json"
	"log"
	"os"

	"github.com/julesxx/cli-do-ris/models"
)

const (
	FilePath = "./"
	FileName = "dump.json"
)

func DumpJson(cards []models.Cards) (b []byte, err error) {
	b, err = json.Marshal(cards)
	return
}
func DecodeJson(content []byte) (cards []models.Cards, err error) {
	cards = []models.Cards{}
	err = json.Unmarshal(content, &cards)
	return
}

func WriteFile(filePath, fileName string, content []byte) (isWritten bool, err error) {

	isWritten = false

	err = os.WriteFile(setFileString(filePath, fileName), content, 0644)
	if err != nil {
		log.Fatal(err)

		return
	}

	isWritten = true

	return

}

func setFileString(filePath, fileName string) string {
	if filePath == "" && fileName == "" {
		filePath = FilePath
		fileName = FileName
	}

	return filePath + fileName
}
func LoadFile(filePath, fileName string) (b []byte, err error) {
	b, err = os.ReadFile(setFileString(filePath, fileName))
	return
}
