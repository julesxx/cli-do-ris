package models

import (
	"errors"

	"github.com/google/uuid"
)

type Card struct {
	Id          uuid.UUID
	Title       string
	Description string
	Done        bool
}

func NewCard(title, desc string, status bool) *Card {
	return &Card{
		Id:          uuid.New(),
		Title:       title,
		Description: desc,
		Done:        false,
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

func (c *Card) setDone(isDone bool) Card {
	c.Done = isDone
	return *c
}
func (c *Card) setTitle(title string) Card {
	c.Title = title
	return *c
}

func (c *Card) setDesc(desc string) Card {
	c.Description = desc
	return *c
}
func (c Cards) getIndexFromCard(id uuid.UUID) int {
	for i, v := range c.CardList {
		if v.Id == id {
			return i
		}
	}
	return -1
}

func (c *Cards) RemoveCard(id uuid.UUID) ([]Card, error) {
	index := c.getIndexFromCard(id)
	if index != -1 {
		c.CardList = append(c.CardList[:index], c.CardList[index+1:]...)
		return c.CardList, nil
	}
	return make([]Card, 0), errors.New("Card with " + (id).String() + "not Found")

}
func (c *Cards) AddCard(card Card) []Card {
	c.CardList = append(c.CardList, card)
	return c.CardList
}
