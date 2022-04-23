package red_black_tree

import "testing"

func TestInsertion(t *testing.T) {
	var tests = []struct {
		key int
		value int
		expected []node
	}{
		{1, 1, []node{
			{1, 1, nil, nil, black},
		}},
		{9, 9, []node{
			{1, 1, nil, nil, black},
			{9, 9, nil, nil, red},
		}},
		{2, 2, []node{
			{2, 2, nil, nil, black},
			{1, 1, nil, nil, red},
			{9, 9, nil, nil, red},
		}},
		{8, 8, []node{
			{2, 2, nil, nil, black},
			{1, 1, nil, nil, black},
			{9, 9, nil, nil, black},
			{8, 8, nil, nil, red},
		}},
		{3, 3, []node{
			{2, 2, nil, nil, black},
			{1, 1, nil, nil, black},
			{8, 8, nil, nil, black},
			{3, 3, nil, nil, red},
			{9, 9, nil, nil, red},
		}},
		{7, 7, []node{
			{2, 2, nil, nil, black},
			{1, 1, nil, nil, black},
			{8, 8, nil, nil, red},
			{3, 3, nil, nil, black},
			{7, 7, nil, nil, red},
			{9, 9, nil, nil, black},
		}},
		{4, 4, []node{
			{2, 2, nil, nil, black},
			{1, 1, nil, nil, black},
			{8, 8, nil, nil, red},
			{4, 4, nil, nil, black},
			{3, 3, nil, nil, red},
			{7, 7, nil, nil, red},
			{9, 9, nil, nil, black},
		}},
		{6, 6, []node{
			{4, 4, nil, nil, black},
			{2, 2, nil, nil, red},
			{1, 1, nil, nil, black},
			{3, 3, nil, nil, black},
			{8, 8, nil, nil, red},
			{7, 7, nil, nil, black},
			{6, 6, nil, nil, red},
			{9, 9, nil, nil, black},
		}},
		{5, 5, []node{
			{4, 4, nil, nil, black},
			{2, 2, nil, nil, black},
			{1, 1, nil, nil, black},
			{3, 3, nil, nil, black},
			{8, 8, nil, nil, black},
			{6, 6, nil, nil, black},
			{5, 5, nil, nil, red},
			{7, 7, nil, nil, red},
			{9, 9, nil, nil, black},
		}},
	}

	const testFailedPanic = "test_failed_panic"
	defer func() {
		// prevent from printing stack trace for test failed panic
		if err := recover(); err != nil && err != testFailedPanic {
			panic(err)
		}
	}()

	tree := New()
	for _, test := range tests {
		tree.Insert(test.key, test.value)

		keys := make([]int, 0, len(test.expected))
		preorderTreeWalk(tree, func(actual *node, i int) {
			keys = append(keys, actual.key)
			expected := test.expected[i]
			if actual.key != expected.key ||
				actual.value != expected.value ||
				actual.color != expected.color {
				t.Errorf("RedBlackTree.Insert(%d, %d) produce result: %v", test.key, test.value, keys)

				// TODO: go try-catch-finally panic recover defer
				panic(testFailedPanic)
			}
		})
	}
}

func preorderTreeWalk(t *RedBlackTree, cb func(node *node, i int)) {
	if t.root == nil {
		return
	}

	var treeWalk func(*node)
	nodeCnt := 0
	treeWalk = func(node *node) {
		cb(node, nodeCnt)
		nodeCnt += 1

		if node.left != t.nilNode {
			treeWalk(node.left)
		}

		if node.right != t.nilNode {
			treeWalk(node.right)
		}
	}

	treeWalk(t.root)
}
