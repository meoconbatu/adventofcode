package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in       string
		expected *List
	}{
		{"[1,2]", &List{-1, []*List{{1, nil}, {2, nil}}}},
		{"[[1,2],[3,4]]", &List{-1, []*List{{-1, []*List{{1, nil}, {2, nil}}}, {-1, []*List{{3, nil}, {4, nil}}}}}},
		{"[3,[1,2]]", &List{-1, []*List{{3, nil}, {-1, []*List{{1, nil}, {2, nil}}}}}},
		{"[3,[4,[1,2]]]", &List{-1, []*List{{3, nil}, {-1, []*List{{4, nil}, {-1, []*List{{1, nil}, {2, nil}}}}}}}},
		{"[3,[4,[1,2]],[5,6]]", &List{-1, []*List{{3, nil}, {-1, []*List{{4, nil}, {-1, []*List{{1, nil}, {2, nil}}}}}, {-1, []*List{{5, nil}, {6, nil}}}}}},
		{"[3,[4,[1,2]],[5,6],7]", &List{-1, []*List{{3, nil}, {-1, []*List{{4, nil}, {-1, []*List{{1, nil}, {2, nil}}}}}, {-1, []*List{{5, nil}, {6, nil}}}, {7, nil}}}},
		{"[3,4,[1,2]]", &List{-1, []*List{{3, nil}, {4, nil}, {-1, []*List{{1, nil}, {2, nil}}}}}},
		{"[[[1,2],[3,4]],[3,[1,2]]]", &List{-1, []*List{{-1, []*List{{-1, []*List{{1, nil}, {2, nil}}}, {-1, []*List{{3, nil}, {4, nil}}}}}, {-1, []*List{{3, nil}, {-1, []*List{{1, nil}, {2, nil}}}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			index := 1
			temp := parse(tt.in, &index)
			assert.Equal(t, tt.expected, temp)
		})
	}
}
