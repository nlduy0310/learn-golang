package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var ranks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

type NewOpt func([]Card) []Card

func New(opts ...NewOpt) []Card {
	var cards []Card

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(fn func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(c []Card) []Card {
		sort.Slice(c, fn(c))
		return c
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].absRank() < cards[j].absRank()
	}
}

func (c Card) absRank() int {
	return int(c.Suit)*int(King) + int(c.Rank)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func Shuffle(cards []Card) []Card {
	var res = make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		res[i] = cards[j]
	}
	return res
}

func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

func Filter(fn func(Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var res []Card
		for _, card := range cards {
			if !fn(card) {
				res = append(res, card)
			}
		}
		return res
	}
}

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var res = make([]Card, 0, len(cards)*n)
		for i := 0; i < n; i++ {
			res = append(res, cards...)
		}
		return res
	}
}
