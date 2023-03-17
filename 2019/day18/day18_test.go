package day18

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMinStep(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{strings.Split(`#########
#b.A.@.a#
#########`, "\n"), 8},
		{strings.Split(`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`, "\n"), 86},
		{strings.Split(`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`, "\n"), 132},
		{strings.Split(`########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`, "\n"), 81},
		{strings.Split(`#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`, "\n"), 136},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.in), func(t *testing.T) {
			actual := getMinStep(tt.in)
			assert.Equal(t, tt.out, actual)
		})
	}
}
