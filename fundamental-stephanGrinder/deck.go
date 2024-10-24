package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newCards() deck {
	var cards deck
	suit := []string{"Spades", "Diamond", "Heats", "Clubs"}
	number := []string{"Ace", "Two", "Three", "Four"}
	for _, suitName := range suit {

		for _, num := range number {
			cards = append(cards, num+" of "+suitName)
		}
	}
	return cards
}

func handSize(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {

	for index, value := range d {
		fmt.Println(index, value)
	}
}
func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveFile(fileName string) error {

	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		return err
	}
	return nil
}

func readFile(fileName string) (deck, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return deck(strings.Split(string(data), ",")), nil
	}
	return nil, err

}

func (d deck) shuffle() {
	// rand.Seed(time.Now().UnixMicro())
	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)
	for index := range d {
		newIndx := r.Intn(len(d) - 1)
		d[index], d[newIndx] = d[newIndx], d[index]
	}

}
