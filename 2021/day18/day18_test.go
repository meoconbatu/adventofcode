package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		in       string
		expected *Node
	}{
		{"[1,2]", &Node{-1, &Node{1, nil, nil}, &Node{2, nil, nil}}},
		{"[[1,2],[3,4]]", &Node{-1, &Node{-1, &Node{1, nil, nil}, &Node{2, nil, nil}}, &Node{-1, &Node{3, nil, nil}, &Node{4, nil, nil}}}},
		{"[[[1,2],[3,4]],[5,6]]", &Node{-1, &Node{-1, &Node{-1, &Node{1, nil, nil}, &Node{2, nil, nil}}, &Node{-1, &Node{3, nil, nil}, &Node{4, nil, nil}}}, &Node{-1, &Node{5, nil, nil}, &Node{6, nil, nil}}}},
		{"[[1,2],3]", &Node{-1, &Node{-1, &Node{1, nil, nil}, &Node{2, nil, nil}}, &Node{3, nil, nil}}},
		{"[3,[1,2]]", &Node{-1, &Node{3, nil, nil}, &Node{-1, &Node{1, nil, nil}, &Node{2, nil, nil}}}},
		{"[[[1,2],3],[4,4]]", &Node{-1, &Node{-1, &Node{-1, &Node{1, nil, nil}, &Node{2, nil, nil}}, &Node{3, nil, nil}}, &Node{-1, &Node{4, nil, nil}, &Node{4, nil, nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			temp, _ := buildTree(tt.in, 0)
			assert.Equal(t, tt.expected, temp)
		})
	}
}

func TestBFS(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]", "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{"[[[[[3,0],[5,3]],[4,4]],[5,5]],[6,6]]", "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
		{"[[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]],[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]]", "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"},
		{"[[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]],[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]]", "[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tree, _ := buildTree(tt.in, 0)
			reductionTree(tree)
			extree, _ := buildTree(tt.expected, 0)
			assert.Equal(t, extree, tree)
		})
	}
}
