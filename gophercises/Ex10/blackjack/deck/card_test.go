package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Rank: Nine, Suit: Spade})
	fmt.Println(Card{Rank: Six, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Nine of Spades
	// Six of Diamonds
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	aceOfSpades := Card{Rank: Ace, Suit: Spade}

	if cards[0] != aceOfSpades {
		t.Errorf("expected Ace of Spades as first card, received: %s", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))

	aceOfSpades := Card{Rank: Ace, Suit: Spade}

	if cards[0] != aceOfSpades {
		t.Errorf("expected Ace of Spades as first card, received: %s", cards[0])
	}
}

func TestJokers(t *testing.T) {
	nJokers := 3
	cards := New(Jokers(nJokers))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != nJokers {
		t.Errorf("Expected %d jokers, received %d", nJokers, count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(c Card) bool {
		return c.Rank == Two || c.Rank == 3
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == 3 {
			t.Error("Expected all twos and threes to be removed from the deck")
		}
	}
}

func TestDeck(t *testing.T) {
	nDecks := 3
	cards := New(Deck(nDecks))
	if len(cards) != 13*4*nDecks {
		t.Errorf("Expected %d cards, received %d", 13*4*nDecks, len(cards))
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0)) // a deterministic rand
	original := New()
	first := original[40]
	second := original[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Errorf("Expected first card to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected second card to be %s, received %s", second, cards[1])
	}
}
