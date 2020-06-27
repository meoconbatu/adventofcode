package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12Part1GetState(t *testing.T) {
	tests := []struct {
		input       string
		stepth      int
		expectedPos []point3D
		expectedVel []point3D
	}{
		{`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`, 10, []point3D{{2, 1, -3}, {1, -8, 0}, {3, -6, 1}, {2, 0, 4}},
			[]point3D{{-3, -2, 1}, {-1, 1, 3}, {3, 2, -3}, {1, -1, -1}}},
		{`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`, 1, []point3D{{2, -1, 1}, {3, -7, -4}, {1, -7, 5}, {2, 2, 0}},
			[]point3D{{3, -1, -1}, {1, 3, 3}, {-3, 1, -3}, {-1, -3, 1}}},

		{`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`, 10, []point3D{{-9, -10, 1}, {4, 10, 9}, {8, -10, -3}, {5, -10, 3}},
			[]point3D{{-2, -2, -1}, {-3, 7, -2}, {5, -1, -2}, {0, -4, 5}}},
	}
	for _, tt := range tests {
		buf := &bytes.Buffer{}
		buf.WriteString(tt.input)
		pos, vel := day12Part1GetState(buf, tt.stepth)
		assert.Equal(t, tt.expectedPos, pos)
		assert.Equal(t, tt.expectedVel, vel)
	}
}
func TestCalcEnegy(t *testing.T) {
	tests := []struct {
		pos            []point3D
		vel            []point3D
		expectedOutput float64
	}{
		{[]point3D{{8, -12, -9}, {13, 16, -3}, {-29, -11, -1}, {16, -13, 23}},
			[]point3D{{-7, 3, 0}, {3, -11, -5}, {-3, 7, 4}, {7, 1, 1}},
			1940},
	}

	for _, tt := range tests {
		result := calcEnegy(tt.pos, tt.vel)
		assert.Equal(t, tt.expectedOutput, result)
	}
}

func TestDay12Part2GetStepth(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`, 2772},
		{`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`, 4686774924},
	}
	for _, tt := range tests {
		buf := &bytes.Buffer{}
		buf.WriteString(tt.input)
		result := day12Part2GetStepthThread(buf)
		assert.Equal(t, tt.expectedOutput, result)
	}
}
