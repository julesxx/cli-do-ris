package models

import "github.com/google/uuid"

type Card struct {
	Id          uuid.UUID
	Title       string
	Description string
}

func NewCard(title, desc string) *Card {
	return &Card{
		Id:          uuid.New(),
		Title:       title,
		Description: desc,
	}
}

type Cards struct {
	Id       uuid.UUID
	Title    string
	CardList []Card
}

func NewCards(title string, cardList []Card) *Cards {
	return &Cards{
		Id:       uuid.New(),
		Title:    title,
		CardList: cardList,
	}
}
