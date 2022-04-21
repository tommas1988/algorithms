package red_black_tree

import "testing"

func TestInsertion(t *testing.T) {
	var tests = []struct {
		key int
		value int
		expected []node
	}{
		{1, 1, []node{
			{1, 1, black},
		}},
		{9, 9, []node{
			{1, 1, black},
			{9, 9, red},
		}},
		{2, 2, []node{
			{2, 2, black},
			{1, 1, red},
			{9, 9, red},
		}},
		{8, 8, []node{
			{2, 2, black},
			{1, 1, black},
			{9, 9, black},
			{8, 8, red},
		}},
		{3, 3, []node{
			{2, 2, black},
			{1, 1, black},
			{8, 8, black},
			{3, 3, red},
			{9, 9, red},
		}},
		{7, 7, []node{
			{2, 2, black},
			{1, 1, black},
			{8, 8, red},
			{3, 3, black},
			{7, 7, red},
			{9, 9, black},
		}},
		{4, 4, []node{
			{2, 2, black},
			{1, 1, black},
			{8, 8, red},
			{4, 4, black},
			{3, 3, red},
			{7, 7, red},
			{9, 9, black},
		}},
		{6, 6, []node{
			{4, 4, black},
			{2, 2, red},
			{1, 1, black},
			{3, 3, black},
			{8, 8, red},
			{7, 7, black},
			{6, 6, red},
			{9, 9, black},
		}},
		{5, 5, []node{
			{4, 4, black},
			{2, 2, red},
			{1, 1, black},
			{3, 3, black},
			{8, 8, red},
			{6, 6, black},
			{5, 5, red},
			{7, 7, red},
			{9, 9, black},
		}},
	}
}

func inorderTreeWalk(t *rbtree) []node {
	
}
