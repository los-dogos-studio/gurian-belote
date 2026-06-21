package game

type Declaration interface {
	Points() int
}

type PreHandDeclarationType int

type PreHandDeclaration struct {
	Type        PreHandDeclarationType
	HighestCard Card
}

type Belote struct{}

const (
	DeclarationNone PreHandDeclarationType = iota
	Tierce
	Quarte
	Quinte
	Carre
	NinesCarre
	JacksCarre
)

var declarationPoints = map[PreHandDeclarationType]int{
	DeclarationNone: 0,
	Tierce:          20,
	Quarte:          50,
	Quinte:          100,
	Carre:           100,
	NinesCarre:      150,
	JacksCarre:      200,
}

func (b Belote) Points() int {
	return 20
}

func (d PreHandDeclaration) Points() int {
	return declarationPoints[d.Type]
}

func CompareDeclarations(d1, d2 PreHandDeclaration, trump Suit) int {
	if d1.Type != d2.Type {
		return int(d1.Type) - int(d2.Type)
	}

	if d1.HighestCard.Suit == trump && d2.HighestCard.Suit != trump {
		return 1
	} else if d1.HighestCard.Suit != trump && d2.HighestCard.Suit == trump {
		return -1
	}

	return d1.HighestCard.Rank.NaturalOrder() - d2.HighestCard.Rank.NaturalOrder()
}
