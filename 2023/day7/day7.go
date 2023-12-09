package day7

import (
	"fmt"
	"sort"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day7 type
type Day7 struct{}

type HandType int

type Line struct {
	hand string
	bid  int
	typ  HandType
}

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardLabels = map[rune]int{'2': 1, '3': 2, '4': 3, '5': 4,
	'6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}

// Part1 func
func (d Day7) Part1() {
	scanner := utils.NewScanner(7)
	rs := 0
	lines := make([]Line, 0)
	for scanner.Scan() {
		var hand string
		var bid int
		fmt.Sscanf(scanner.Text(), "%s %d\n", &hand, &bid)
		lines = append(lines, Line{hand, bid, getHandType(hand)})
	}
	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].typ < lines[j].typ ||
			(lines[i].typ == lines[j].typ && compareHand(lines[i].hand, lines[j].hand) <= 0)
	})
	for i, line := range lines {
		rs += (i + 1) * line.bid
	}
	fmt.Println(rs)
}

// hand1 < hand2: -1
// hand1 == hand2: 0
// hand1 > hand2: 1
func compareHand(hand1, hand2 string) int {
	for i := 0; i < len(hand1); i++ {
		if cardLabels[rune(hand1[i])] > cardLabels[rune(hand2[i])] {
			return 1
		}
		if cardLabels[rune(hand1[i])] < cardLabels[rune(hand2[i])] {
			return -1
		}
	}
	return 0
}

// (1: 5) Five of a kind, where all five cards have the same label: AAAAA
// (2: 4,1) Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// (2: 3,2) Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// (3: 3,1,1) Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// (3: 2,2,1) Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// (4: 2,1,1,1) One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// (5: 1) High card, where all cards' labels are distinct: 23456
func getHandType(hand string) HandType {
	labelToCount := make(map[rune]int)
	for _, label := range hand {
		labelToCount[label]++
	}
	return fn(labelToCount)
}

var cardLabel2s = map[rune]int{'J': 0, '2': 1, '3': 2, '4': 3, '5': 4,
	'6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'Q': 12, 'K': 13, 'A': 14}

// Part2 func
func (d Day7) Part2() {
	scanner := utils.NewScanner(7)
	rs := 0
	lines := make([]Line, 0)
	for scanner.Scan() {
		var hand string
		var bid int
		fmt.Sscanf(scanner.Text(), "%s %d\n", &hand, &bid)
		lines = append(lines, Line{hand, bid, getHandType2(hand)})
	}
	sort.SliceStable(lines, func(i, j int) bool {
		return lines[i].typ < lines[j].typ ||
			(lines[i].typ == lines[j].typ && compareHand2(lines[i].hand, lines[j].hand) <= 0)
	})
	for i, line := range lines {
		rs += (i + 1) * line.bid
	}
	fmt.Println(rs)
}

func compareHand2(hand1, hand2 string) int {
	for i := 0; i < len(hand1); i++ {
		if cardLabel2s[rune(hand1[i])] > cardLabel2s[rune(hand2[i])] {
			return 1
		}
		if cardLabel2s[rune(hand1[i])] < cardLabel2s[rune(hand2[i])] {
			return -1
		}
	}
	return 0
}

func getHandType2(hand string) HandType {
	labelToCount := make(map[rune]int)
	maxLabel, maxCnt := '0', 0
	for _, label := range hand {
		labelToCount[label]++
		if label != 'J' {
			if maxCnt < labelToCount[label] {
				maxCnt = labelToCount[label]
				maxLabel = label
			}
		}
	}
	if labelToCount['J'] > 0 {
		labelToCount[maxLabel] += labelToCount['J']
	}
	delete(labelToCount, 'J')
	return fn(labelToCount)
}
func fn(labelToCount map[rune]int) HandType {
	switch len(labelToCount) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		for _, cnt := range labelToCount {
			if cnt == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 2:
		for _, cnt := range labelToCount {
			if cnt == 4 || cnt == 1 {
				return FourOfAKind
			}
			return FullHouse
		}
	case 1:
		return FiveOfAKind
	}
	return HighCard
}
