package deck

import (
	"crypto/rand"
	"math/big"
)

type Card struct {
	FaceValue Value
	Suit      string
}

type Value struct {
	Name  string
	Value int
}

type Deck struct {
	Cards []Card
}

func (d *Deck) BuildDeck() {

	suits := []string{
		"Diamonds",
		"Clubs",
		"Hearts",
		"Spades",
	}

	facevalues := []Value{
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"10", 10},
		{"Jack", 10},
		{"Queen", 10},
		{"King", 10},
		{"Ace", 11},
	}

	d.Cards = nil

	for _, suit := range suits {
		for _, facevalue := range facevalues {
			d.Cards = append(d.Cards, Card{FaceValue: facevalue, Suit: suit})
		}
	}
}

func (d *Deck) Shuffle() {
	var old []Card
	old = d.Cards
	var shuffled []Card

	for i := len(old); i > 0; i-- {
		nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(i)))
		j := nBig.Int64()
		shuffled = append(shuffled, old[j])

		old = remove(old, j)
	}
	d.Cards = shuffled
}

func (d *Deck) Draw() Card {
	card := d.Cards[0]
	d.Cards = remove(d.Cards, 0)
	return card
}

func remove(slice []Card, i int64) []Card {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
