package game

import "sort"

func FindPreHandDeclarations(cards []Card) []PreHandDeclaration {
	var result []PreHandDeclaration
	result = append(result, findSequences(cards)...)
	result = append(result, findCarres(cards)...)
	return result
}

func findSequences(cards []Card) []PreHandDeclaration {
	bySuit := map[Suit][]Card{}
	for _, c := range cards {
		bySuit[c.Suit] = append(bySuit[c.Suit], c)
	}

	var result []PreHandDeclaration
	for _, suitCards := range bySuit {
		sort.Slice(suitCards, func(i, j int) bool {
			return suitCards[i].Rank.NaturalOrder() < suitCards[j].Rank.NaturalOrder()
		})

		runStart := 0
		for i := 1; i <= len(suitCards); i++ {
			if i == len(suitCards) || suitCards[i].Rank.NaturalOrder() != suitCards[i-1].Rank.NaturalOrder()+1 {
				runLen := i - runStart
				if runLen >= 3 {
					var declType PreHandDeclarationType
					switch runLen {
					case 3:
						declType = Tierce
					case 4:
						declType = Quarte
					default:
						declType = Quinte
					}
					result = append(result, PreHandDeclaration{Type: declType, HighestCard: suitCards[i-1]})
				}
				runStart = i
			}
		}
	}
	return result
}

func findCarres(cards []Card) []PreHandDeclaration {
	countByRank := map[Rank]int{}
	firstByRank := map[Rank]Card{}
	for _, c := range cards {
		countByRank[c.Rank]++
		if _, seen := firstByRank[c.Rank]; !seen {
			firstByRank[c.Rank] = c
		}
	}

	var result []PreHandDeclaration
	for rank, count := range countByRank {
		if count == 4 {
			var declType PreHandDeclarationType
			switch rank {
			case Seven, Eight:
				continue
			case Jack:
				declType = JacksCarre
			case Nine:
				declType = NinesCarre
			default:
				declType = Carre
			}
			result = append(result, PreHandDeclaration{Type: declType, HighestCard: firstByRank[rank]})
		}
	}
	return result
}

func HasBelote(cards []Card, trump Suit) bool {
	hasKing, hasQueen := false, false
	for _, c := range cards {
		if c.Suit == trump {
			switch c.Rank {
			case King:
				hasKing = true
			case Queen:
				hasQueen = true
			}
		}
	}
	return hasKing && hasQueen
}
