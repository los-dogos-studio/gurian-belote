package game

import "testing"

func TestFindNoDeclarations(t *testing.T) {
	cards := []Card{
		{Suit: Spades, Rank: Seven},
		{Suit: Hearts, Rank: Nine},
		{Suit: Diamonds, Rank: Jack},
		{Suit: Clubs, Rank: Ace},
		{Suit: Spades, Rank: Queen},
		{Suit: Hearts, Rank: King},
		{Suit: Diamonds, Rank: Eight},
		{Suit: Clubs, Rank: Ten},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 0 {
		t.Errorf("expected no declarations, got %d", len(result))
	}
}

func TestFindTierce(t *testing.T) {
	cards := []Card{
		{Suit: Hearts, Rank: Seven},
		{Suit: Hearts, Rank: Eight},
		{Suit: Hearts, Rank: Nine},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 1 {
		t.Fatalf("expected 1 declaration, got %d", len(result))
	}
	d := result[0]
	if d.Type != Tierce {
		t.Errorf("expected Tierce, got %+v", d)
	}
	if d.HighestCard.Rank != Nine || d.HighestCard.Suit != Hearts {
		t.Errorf("expected highest card Nine of Hearts, got %v", d.HighestCard)
	}
}

func TestFindQuarte(t *testing.T) {
	cards := []Card{
		{Suit: Diamonds, Rank: Jack},
		{Suit: Diamonds, Rank: Queen},
		{Suit: Diamonds, Rank: King},
		{Suit: Diamonds, Rank: Ace},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 1 {
		t.Fatalf("expected 1 declaration, got %d", len(result))
	}
	d := result[0]
	if d.Type != Quarte {
		t.Errorf("expected Quarte, got %v", d.Type)
	}
	if d.HighestCard.Rank != Ace || d.HighestCard.Suit != Diamonds {
		t.Errorf("expected highest card Ace of Diamonds, got %v", d.HighestCard)
	}
}

func TestFindQuinte(t *testing.T) {
	cards := []Card{
		{Suit: Clubs, Rank: Seven},
		{Suit: Clubs, Rank: Eight},
		{Suit: Clubs, Rank: Nine},
		{Suit: Clubs, Rank: Ten},
		{Suit: Clubs, Rank: Jack},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 1 {
		t.Fatalf("expected 1 declaration, got %d", len(result))
	}
	d := result[0]
	if d.Type != Quinte {
		t.Errorf("expected Quinte, got %v", d.Type)
	}
	if d.HighestCard.Rank != Jack || d.HighestCard.Suit != Clubs {
		t.Errorf("expected highest card Jack of Clubs, got %v", d.HighestCard)
	}
}

func TestFindCarre(t *testing.T) {
	cards := []Card{
		{Suit: Spades, Rank: Ace},
		{Suit: Hearts, Rank: Ace},
		{Suit: Diamonds, Rank: Ace},
		{Suit: Clubs, Rank: Ace},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 1 {
		t.Fatalf("expected 1 declaration, got %d", len(result))
	}
	d := result[0]
	if d.Type != Carre {
		t.Errorf("expected Carre, got %v", d.Type)
	}
	if d.HighestCard.Rank != Ace {
		t.Errorf("expected highest card Ace, got %v", d.HighestCard.Rank)
	}
}

func TestNoBlankCarres(t *testing.T) {
	cards := []Card{
		{Suit: Spades, Rank: Seven},
		{Suit: Hearts, Rank: Seven},
		{Suit: Diamonds, Rank: Seven},
		{Suit: Clubs, Rank: Seven},
		{Suit: Spades, Rank: Eight},
		{Suit: Hearts, Rank: Eight},
		{Suit: Diamonds, Rank: Eight},
		{Suit: Clubs, Rank: Eight},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 0 {
		t.Errorf("expected no declarations for 7s and 8s carres, got %v", result)
	}
}

func TestFindBelote(t *testing.T) {
	cards := []Card{
		{Suit: Hearts, Rank: King},
		{Suit: Hearts, Rank: Queen},
		{Suit: Spades, Rank: Seven},
	}
	if !HasBelote(cards, Hearts) {
		t.Errorf("expected HasBelote to be true for K+Q of trump")
	}
}

func TestNoBeloteWithoutTrump(t *testing.T) {
	cards := []Card{
		{Suit: Hearts, Rank: King},
		{Suit: Hearts, Rank: Queen},
	}
	if HasBelote(cards, Spades) {
		t.Errorf("expected HasBelote to be false when K+Q are not in trump suit")
	}
}

func TestDisjointRunsInSameSuit(t *testing.T) {
	// 7-8-9 and Q-K-A are two separate Tierces, 10-J breaks them apart
	cards := []Card{
		{Suit: Spades, Rank: Seven},
		{Suit: Spades, Rank: Eight},
		{Suit: Spades, Rank: Nine},
		{Suit: Spades, Rank: Queen},
		{Suit: Spades, Rank: King},
		{Suit: Spades, Rank: Ace},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 2 {
		t.Fatalf("expected 2 Tierce declarations, got %d", len(result))
	}
	highestRanks := map[Rank]bool{}
	for _, d := range result {
		if d.Type != Tierce {
			t.Errorf("expected Tierce PreHandDeclaration, got %+v", d)
			continue
		}
		highestRanks[d.HighestCard.Rank] = true
	}
	if !highestRanks[Nine] {
		t.Errorf("expected a Tierce with highest card Nine")
	}
	if !highestRanks[Ace] {
		t.Errorf("expected a Tierce with highest card Ace")
	}
}

func TestSixCardSequence(t *testing.T) {
	cards := []Card{
		{Suit: Hearts, Rank: Seven},
		{Suit: Hearts, Rank: Eight},
		{Suit: Hearts, Rank: Nine},
		{Suit: Hearts, Rank: Ten},
		{Suit: Hearts, Rank: Jack},
		{Suit: Hearts, Rank: Queen},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 1 {
		t.Fatalf("expected 1 declaration, got %d", len(result))
	}
	d := result[0]
	if d.Type != Quinte {
		t.Errorf("expected Quinte for 6-card sequence, got %+v", d)
	}
	if d.HighestCard.Rank != Queen || d.HighestCard.Suit != Hearts {
		t.Errorf("expected highest card Queen of Hearts, got %v", d.HighestCard)
	}
}

func TestOverlappingSequenceAndCarre(t *testing.T) {
	cards := []Card{
		{Suit: Spades, Rank: Jack},
		{Suit: Spades, Rank: Queen},
		{Suit: Spades, Rank: King},
		{Suit: Hearts, Rank: Jack},
		{Suit: Diamonds, Rank: Jack},
		{Suit: Clubs, Rank: Jack},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 2 {
		t.Fatalf("expected 2 declarations (Tierce + JacksCarre), got %d", len(result))
	}
	var tierce, jacksCarre *PreHandDeclaration
	for _, d := range result {
		d := d
		switch d.Type {
		case Tierce:
			tierce = &d
		case JacksCarre:
			jacksCarre = &d
		default:
			t.Errorf("unexpected declaration type %v", d.Type)
		}
	}
	if tierce == nil {
		t.Errorf("expected a Tierce declaration")
	} else if tierce.HighestCard.Rank != King || tierce.HighestCard.Suit != Spades {
		t.Errorf("expected Tierce highest card King of Spades, got %+v", tierce.HighestCard)
	}
	if jacksCarre == nil {
		t.Errorf("expected a JacksCarre declaration")
	} else if jacksCarre.HighestCard.Rank != Jack {
		t.Errorf("expected JacksCarre highest card Jack, got %+v", jacksCarre.HighestCard)
	}
}

func TestSequenceAndCarreInSameHand(t *testing.T) {
	cards := []Card{
		{Suit: Hearts, Rank: Seven},
		{Suit: Hearts, Rank: Eight},
		{Suit: Hearts, Rank: Nine},
		{Suit: Spades, Rank: Jack},
		{Suit: Hearts, Rank: Jack},
		{Suit: Diamonds, Rank: Jack},
		{Suit: Clubs, Rank: Jack},
		{Suit: Clubs, Rank: Ace},
	}
	result := FindPreHandDeclarations(cards)
	if len(result) != 2 {
		t.Fatalf("expected 2 declarations (Tierce + JacksCarre), got %d", len(result))
	}
	var tierce, jacksCarre *PreHandDeclaration
	for _, d := range result {
		d := d
		switch d.Type {
		case Tierce:
			tierce = &d
		case JacksCarre:
			jacksCarre = &d
		default:
			t.Errorf("unexpected declaration type %v", d.Type)
		}
	}
	if tierce == nil {
		t.Errorf("expected a Tierce declaration")
	} else if tierce.HighestCard.Rank != Nine || tierce.HighestCard.Suit != Hearts {
		t.Errorf("expected Tierce highest card Nine of Hearts, got %+v", tierce.HighestCard)
	}
	if jacksCarre == nil {
		t.Errorf("expected a JacksCarre declaration")
	} else if jacksCarre.HighestCard.Rank != Jack {
		t.Errorf("expected JacksCarre highest card Jack, got %+v", jacksCarre.HighestCard)
	}
}
