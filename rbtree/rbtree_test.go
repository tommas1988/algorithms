package rbtree

import "testing"

const testFailedPanic = "test_failed_panic"

func TestTopDownInsertion(t *testing.T) {
	var tests = []struct {
		key      int
		value    int
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

	defer testFailedHandler()

	tree := New(TopDown)
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

				// TODO: [STY] go try-catch-finally panic recover defer
				panic(testFailedPanic)
			}
		})
	}
}

func TestTopDownDeletion(t *testing.T) {
	keys := []int{41, 38, 31, 12, 19, 8}

	var tests = []struct {
		key      int
		expected []node
	}{
		{8, []node{
			{38, 38, nil, nil, black},
			{19, 19, nil, nil, red},
			{12, 12, nil, nil, black},
			{31, 31, nil, nil, black},
			{41, 41, nil, nil, black},
		}},
		{12, []node{
			{38, 38, nil, nil, black},
			{19, 19, nil, nil, black},
			{31, 31, nil, nil, red},
			{41, 41, nil, nil, black},
		}},
		{19, []node{
			{38, 38, nil, nil, black},
			{31, 31, nil, nil, black},
			{41, 41, nil, nil, black},
		}},
		{31, []node{
			{38, 38, nil, nil, black},
			{41, 41, nil, nil, red},
		}},
		{38, []node{
			{41, 41, nil, nil, black},
		}},
		{41, []node{}},
	}

	defer testFailedHandler()

	tree := New(TopDown)
	for _, k := range keys {
		tree.Insert(k, k)
	}
	initail := []node{
		{38, 38, nil, nil, black},
		{19, 19, nil, nil, red},
		{12, 12, nil, nil, black},
		{8, 8, nil, nil, red},
		{31, 31, nil, nil, black},
		{41, 41, nil, nil, black},
	}

	actualkeys := make([]int, 0, len(keys))
	preorderTreeWalk(tree, func(actual *node, i int) {
		actualkeys = append(actualkeys, actual.key)
		expected := initail[i]
		if actual.key != expected.key ||
			actual.value != expected.value ||
			actual.color != expected.color {
			t.Errorf("Unexpected initail RedBlackTree structual: %v", actualkeys)

			panic(testFailedPanic)
		}
	})

	for _, test := range tests {
		if !tree.Delete(test.key) {
			t.Errorf("RedBlackTree.Delete(%d) = false", test.key)
			return
		}

		actualkeys = make([]int, 0, len(test.expected))
		preorderTreeWalk(tree, func(actual *node, i int) {
			actualkeys = append(actualkeys, actual.key)
			expected := test.expected[i]
			if actual.key != expected.key ||
				actual.value != expected.value ||
				actual.color != expected.color {
				t.Errorf("RedBlackTree.Delete(%d) produce tree structual: %v", test.key, actualkeys)

				panic(testFailedPanic)
			}
		})
	}
}

func TestBottomUpInsertion(t *testing.T) {
	var tests = []struct {
		key      int
		value    int
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
			{2, 2, nil, nil, red},
			{1, 1, nil, nil, black},
			{3, 3, nil, nil, black},
			{8, 8, nil, nil, red},
			{6, 6, nil, nil, black},
			{5, 5, nil, nil, red},
			{7, 7, nil, nil, red},
			{9, 9, nil, nil, black},
		}},
	}

	defer testFailedHandler()

	tree := New(BottomUp)
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

				// TODO: [STY] go try-catch-finally panic recover defer
				panic(testFailedPanic)
			}
		})
	}
}

func testFailedHandler() {
	// prevent from printing stack trace for test failed panic
	if err := recover(); err != nil && err != testFailedPanic {
		panic(err)
	}
}

func preorderTreeWalk(t *RedBlackTree, cb func(node *node, i int)) {
	if t.root == t.nilNode {
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
