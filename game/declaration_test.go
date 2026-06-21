package game

import "testing"

func TestNonEqualDeclarationTypes(t *testing.T) {
	for dt := DeclarationNone; dt <= JacksCarre; dt++ {
		for otherDt := dt + 1; otherDt <= JacksCarre; otherDt++ {
			d1 := PreHandDeclaration{Type: dt, HighestCard: Card{Suit: Hearts, Rank: Ace}}
			d2 := PreHandDeclaration{Type: otherDt, HighestCard: Card{Suit: Hearts, Rank: Ace}}
			if CompareDeclarations(d1, d2, Diamonds) >= 0 {
				t.Errorf("Expected declaration type %v to be ranked higher than %v", dt, otherDt)
			}
			if CompareDeclarations(d2, d1, Diamonds) <= 0 {
				t.Errorf("Expected declaration type %v to be ranked lower than %v", otherDt, dt)
			}
		}
	}
}

func TestNonEqualDeclarationTypesIgnoresTrump(t *testing.T) {
	d1 := PreHandDeclaration{Type: Tierce, HighestCard: Card{Suit: Diamonds, Rank: Ace}}
	d2 := PreHandDeclaration{Type: Quarte, HighestCard: Card{Suit: Hearts, Rank: Ace}}

	if CompareDeclarations(d1, d2, Diamonds) >= 0 || CompareDeclarations(d2, d1, Diamonds) <= 0 {
		t.Errorf("Expected declaration type %v to be ranked higher than %v", d1.Type, d2.Type)
	}
}

func TestEqualDeclarationTypesWithTrump(t *testing.T) {
	d1 := PreHandDeclaration{Type: Tierce, HighestCard: Card{Suit: Diamonds, Rank: Ten}}
	d2 := PreHandDeclaration{Type: Tierce, HighestCard: Card{Suit: Hearts, Rank: Ace}}

	if CompareDeclarations(d1, d2, Diamonds) <= 0 || CompareDeclarations(d2, d1, Diamonds) >= 0 {
		t.Errorf("Expected declaration with trump card to be ranked higher than non-trump card")
	}
}

func TestDeclarationsHonorNaturalOrdering(t *testing.T) {
	d1 := PreHandDeclaration{Type: Tierce, HighestCard: Card{Suit: Diamonds, Rank: Ace}}
	d2 := PreHandDeclaration{Type: Tierce, HighestCard: Card{Suit: Diamonds, Rank: Jack}}
	d3 := PreHandDeclaration{Type: Tierce, HighestCard: Card{Suit: Diamonds, Rank: Ten}}

	if CompareDeclarations(d1, d2, Diamonds) <= 0 || CompareDeclarations(d2, d1, Diamonds) >= 0 {
		t.Errorf("Expected declaration with Ace to be ranked higher than Jack")
	}

	if CompareDeclarations(d2, d3, Diamonds) <= 0 || CompareDeclarations(d3, d2, Diamonds) >= 0 {
		t.Errorf("Expected declaration with Jack to be ranked higher than Ten")
	}
}
