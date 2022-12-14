package crud

import (
	"errors"

	"github.com/google/uuid"
	"github.com/julesxx/cli-do-ris/models"
)

func addCard(title, desc string, status bool) {

	models.NewCard(title, desc, status)
}

func addCards(title string, cardList []models.Card) {
	models.NewCards(title, cardList)
}

func deleteCard(id uuid.UUID, cards models.Cards) bool {
	_, err := cards.RemoveCard(id)

	return err != nil

}

func deleteCards(id uuid.UUID, cards []models.Cards) ([]models.Cards, error) {
	index := -1
	for i, v := range cards {
		if v.Id == id {
			index = i
		}
	}

	if index != -1 {
		cards = append(cards[:index], cards[index+1:]...)
		return cards, nil
	}
	return cards, errors.New("Cards with " + (id).String() + "not Found")
}
