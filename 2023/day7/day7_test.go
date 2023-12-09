package day7

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareHand(t *testing.T) {
	tests := []struct {
		hand1, hand2 string
		expected     int
	}{
		{"33332", "2AAAA", 1},
		{"2AAAA", "33332", -1},
		{"2AAAA", "2AAAA", 0},
		{"77888", "77788", 1},
		{"KK677", "KTJJT", 1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s, %s", tc.hand1, tc.hand2), func(t *testing.T) {
			actual := compareHand(tc.hand1, tc.hand2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetHandType(t *testing.T) {
	tests := []struct {
		hand     string
		expected HandType
	}{
		{"33332", FourOfAKind},
		{"2AAAA", FourOfAKind},
		{"77888", FullHouse},
		{"77788", FullHouse},
		{"32T3K", OnePair},
		{"KK677", TwoPair},
		{"KTJJT", TwoPair},
		{"T55J5", ThreeOfAKind},
		{"QQQJA", ThreeOfAKind},
	}
	for _, tc := range tests {
		t.Run(tc.hand, func(t *testing.T) {
			actual := getHandType(tc.hand)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestCompareHand2(t *testing.T) {
	tests := []struct {
		hand1, hand2 string
		expected     int
	}{
		{"T55J5", "QQQJA", -1},
		{"QQQJA", "KTJJT", -1},
		{"JKKK2", "QQQQ2", -1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s, %s", tc.hand1, tc.hand2), func(t *testing.T) {
			actual := compareHand2(tc.hand1, tc.hand2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
func TestGetHandType2(t *testing.T) {
	tests := []struct {
		hand     string
		expected HandType
	}{
		{"QJJQ2", FourOfAKind},
		{"JKKK2", FourOfAKind},
		{"77888", FullHouse},
		{"77788", FullHouse},
		{"32T3K", OnePair},
		{"KK677", TwoPair},
		{"KTJJT", FourOfAKind},
		{"T55J5", FourOfAKind},
		{"QQQJA", FourOfAKind},
		{"QQQJJ", FiveOfAKind},
	}
	for _, tc := range tests {
		t.Run(tc.hand, func(t *testing.T) {
			actual := getHandType2(tc.hand)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
