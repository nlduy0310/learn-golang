package main

import (
	"blackjack/deck"
	"fmt"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i, v := range h {
		strs[i] = v.String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, card := range h {
		score += min(int(card.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var gs GameState
	gs.Shuffle()
	gs.Deal()

	var input string
	for gs.State == StatePlayerTurn {
		fmt.Print("\n     ----------\n")
		fmt.Println("Player: ", gs.Player, " --  currently", gs.Player.Score())
		fmt.Println("Dealer: ", gs.Dealer.DealerString())
		fmt.Println("What is your decision? (h)it or (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			gs.Hit()
		case "s":
			gs.Stand()
		default:
			fmt.Println("That's not a valid option")
		}
	}

	// If the dealer has <= 16, hit.
	// If the dealer has a soft 17, hit. (which translate to having MinScore < 17 and Score == 17)
	for gs.State == StateDealerTurn {
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() < 17) {
			gs.Hit()
		} else {
			gs.Stand()
		}
	}

	gs.EndHand()
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("it isn't currently any player's turn")
	}
}

// func (gs GameState) clone() GameState {
// 	res := GameState{
// 		Deck:   make([]deck.Card, len(gs.Deck)),
// 		State:  gs.State,
// 		Player: make([]deck.Card, len(gs.Player)),
// 		Dealer: make([]deck.Card, len(gs.Dealer)),
// 	}
// 	copy(res.Deck, gs.Deck)
// 	copy(res.Player, gs.Player)
// 	copy(res.Dealer, gs.Dealer)
// 	return res
// }

func (gs *GameState) Shuffle() {
	gs.Deck = deck.New(deck.Deck(3), deck.Shuffle)
}

func (gs *GameState) Deal() {
	gs.Player = make(Hand, 0, 5)
	gs.Dealer = make(Hand, 0, 5)

	var card deck.Card
	for i := 0; i < 2; i++ {
		card, gs.Deck = draw(gs.Deck)
		gs.Player = append(gs.Player, card)
		card, gs.Deck = draw(gs.Deck)
		gs.Dealer = append(gs.Dealer, card)
	}
	gs.State = StatePlayerTurn
}

func (gs *GameState) Hit() {
	hand := gs.CurrentPlayer()
	var card deck.Card
	card, gs.Deck = draw(gs.Deck)
	if hand.Score() > 21 {
		gs.Stand()
	}
	*hand = append(*hand, card)
}

func (gs *GameState) Stand() {
	gs.State++
}

func (gs *GameState) EndHand() {
	// game result
	pScore, dScore := gs.Player.Score(), gs.Dealer.Score()
	fmt.Print("\n\n----- FINAL HANDS -----\n\n")
	fmt.Println("Player: ", gs.Player, "\nScore: ", pScore)
	fmt.Println("Dealer: ", gs.Dealer, "\nScore: ", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore < dScore:
		fmt.Println("You lost")
	case pScore > dScore:
		fmt.Println("Dealer lost")
	case pScore == dScore:
		fmt.Println("It's a tie")
	}
	gs.Player = nil
	gs.Dealer = nil
}
