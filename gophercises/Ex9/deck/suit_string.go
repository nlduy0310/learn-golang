package deck

func (s Suit) String() string {
	switch s {
	case Spade:
		return "Spade"
	case Heart:
		return "Heart"
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	default:
		return "Joker"
	}
}
