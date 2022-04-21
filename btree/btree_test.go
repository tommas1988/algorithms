package btree

import "testing"

func TestInsertion(t *testing.T) {
	var tests = []struct {
		key rune
		expected [][][]rune
	}{
		{'F', [][][]rune{
			{{'F'}},
		}},
		{'S', [][][]rune{
			{{'F', 'S'}},
		}},
		{'Q', [][][]rune{
			{{'F', 'Q', 'S'}},
		}},
		{'K', [][][]rune{
			{{'Q'}},
			{{'F', 'K'}, {'S'}},
		}},
		{'C', [][][]rune{
			{{'Q'}},
			{{'C', 'F', 'K'}, {'S'}},
		}},
		{'L', [][][]rune{
			{{'F', 'Q'}},
			{{'C'}, {'K', 'L'}, {'S'}},
		}},
		{'H', [][][]rune{
			{{'F', 'Q'}},
			{{'C'}, {'H', 'K', 'L'}, {'S'}},
		}},
		{'T', [][][]rune{
			{{'F', 'Q'}},
			{{'C'}, {'H', 'K', 'L'}, {'S', 'T'}},
		}},
		{'V', [][][]rune{
			{{'F', 'Q'}},
			{{'C'}, {'H', 'K', 'L'}, {'S', 'T', 'V'}},
		}},
		{'W', [][][]rune{
			{{'F', 'Q', 'T'}},
			{{'C'}, {'H', 'K', 'L'}, {'S'}, {'V', 'W'}},
		}},
		{'M', [][][]rune{
			{{'Q'}},
			{{'F', 'K'}, {'T'}},
			{{'C'}, {'H'}, {'L', 'M'}, {'S'}, {'V', 'W'}},
		}},
		{'R', [][][]rune{
			{{'Q'}},
			{{'F', 'K'}, {'T'}},
			{{'C'}, {'H'}, {'L', 'M'}, {'R', 'S'}, {'V', 'W'}},
		}},
		{'N', [][][]rune{
			{{'Q'}},
			{{'F', 'K'}, {'T'}},
			{{'C'}, {'H'}, {'L', 'M', 'N'}, {'R', 'S'}, {'V', 'W'}},
		}},
		{'P', [][][]rune{
			{{'Q'}},
			{{'F', 'K', 'M'}, {'T'}},
			{{'C'}, {'H'}, {'L'}, {'N', 'P'}, {'R', 'S'}, {'V', 'W'}},
		}},
		{'A', [][][]rune{
			{{'K', 'Q'}},
			{{'F'}, {'M'}, {'T'}},
			{{'A', 'C'}, {'H'}, {'L'}, {'N', 'P'}, {'R', 'S'}, {'V', 'W'}},
		}},
		{'B', [][][]rune{
			{{'K', 'Q'}},
			{{'F'}, {'M'}, {'T'}},
			{{'A', 'B', 'C'}, {'H'}, {'L'}, {'N', 'P'}, {'R', 'S'}, {'V', 'W'}},
		}},
		{'X', [][][]rune{
			{{'K', 'Q'}},
			{{'F'}, {'M'}, {'T'}},
			{{'A', 'B', 'C'}, {'H'}, {'L'}, {'N', 'P'}, {'R', 'S'}, {'V', 'W', 'X'}},
		}},
		{'Y', [][][]rune{
			{{'K', 'Q'}},
			{{'F'}, {'M'}, {'T', 'W'}},
			{{'A', 'B', 'C'}, {'H'}, {'L'}, {'N', 'P'}, {'R', 'S'}, {'V'}, {'X', 'Y'}},
		}},
		{'D', [][][]rune{
			{{'K', 'Q'}},
			{{'B', 'F'}, {'M'}, {'T', 'W'}},
			{{'A'}, {'C', 'D'}, {'H'}, {'L'}, {'N', 'P'}, {'R', 'S'}, {'V'}, {'X', 'Y'}},
		}},
		{'Z', [][][]rune{
			{{'K', 'Q'}},
			{{'B', 'F'}, {'M'}, {'T', 'W'}},
			{{'A'}, {'C', 'D'}, {'H'}, {'L'}, {'N', 'P'}, {'R', 'S'}, {'V'}, {'X', 'Y', 'Z'}},
		}},
		{'E', [][][]rune{
			{{'K', 'Q'}},
			{{'B', 'F'}, {'M'}, {'T', 'W'}},
			{{'A'}, {'C', 'D', 'E'}, {'H'}, {'L'}, {'N', 'P'}, {'R', 'S'}, {'V'}, {'X', 'Y', 'Z'}},
		}},
	}

	btree := New(2)
	for _, test := range tests {
		btree.Insert(int(test.key), int(test.key))

		actual := convertBtreeToKeyArray(btree)
		if !compare(actual, test.expected) {
			t.Errorf("btree.Insert(int(%c), int(%[1]c)) with result: %v", test.key, actual)
			return
		}
	}
}

func TestDeletion(t *testing.T) {
	keys := []rune{'D', 'E', 'G', 'J', 'K', 'M', 'N', 'O', 'P', 'R', 'S', 'X',
		'Y', 'Z', 'T', 'A', 'C', 'U', 'V', 'B', 'Q', 'L', 'F'}

	var tests = []struct{
		key rune
		expected [][][]rune
	}{
		{'S', [][][]rune{
			{{'M'}},
			{{'C', 'G'}, {'P', 'T', 'X'}},
			{{'A', 'B'}, {'D', 'E', 'F'}, {'J', 'K', 'L'}, {'N', 'O'}, {'Q', 'R'}, {'U', 'V'}, {'Y', 'Z'}},
		}},
		{'B', [][][]rune{
			{{'P'}},
			{{'D', 'G', 'M'}, {'T', 'X'}},
			{{'A', 'C'}, {'E', 'F'}, {'J', 'K', 'L'}, {'N', 'O'}, {'Q', 'R'}, {'U', 'V'}, {'Y', 'Z'}},
		}},
		{'Z', [][][]rune{
			{{'M'}},
			{{'D', 'G'}, {'P', 'T'}},
			{{'A', 'C'}, {'E', 'F'}, {'J', 'K', 'L'}, {'N', 'O'}, {'Q', 'R'}, {'U', 'V', 'X', 'Y'}},
		}},
		{'M', [][][]rune{
			{{'D', 'G', 'L', 'P', 'T'}},
			{{'A', 'C'}, {'E', 'F'}, {'J', 'K'}, {'N', 'O'}, {'Q', 'R'}, {'U', 'V', 'X', 'Y'}},
		}},
		{'T', [][][]rune{
			{{'D', 'G', 'L', 'P', 'U'}},
			{{'A', 'C'}, {'E', 'F'}, {'J', 'K'}, {'N', 'O'}, {'Q', 'R'}, {'V', 'X', 'Y'}},
		}},
		{'P', [][][]rune{
			{{'D', 'G', 'L', 'U'}},
			{{'A', 'C'}, {'E', 'F'}, {'J', 'K'}, {'N', 'O', 'Q', 'R'}, {'V', 'X', 'Y'}},
		}},
	}

	btree := New(3)
	for _, key := range keys {
		btree.Insert(int(key), int(key))
	}

	for _, test := range tests {
		if !btree.Delete(int(test.key)) {
			t.Errorf("btree.Delete(int(%c)) = false", test.key)
		}

		actual := convertBtreeToKeyArray(btree)
		if !compare(actual, test.expected) {
			t.Errorf("btree.Delete(int(%c)) with result: %v", test.key, actual)
			return
		}
	}
}

func convertBtreeToKeyArray(btree *Btree) [][][]rune {
	var context = struct {
		nodes [2][]*btreeNode
		current, next int
	}{
		[2][]*btreeNode{
			{btree.root},
			{},
		},
		0,
		1,
	}

	result := [][][]rune{}
	isLeaf := false

	for !isLeaf {
		current, next := context.current, context.next
		groups := [][]rune{}
		nextNodes := []*btreeNode{}
		for _, node := range context.nodes[current] {
			keys := []rune{}
			for i := 0; i < node.degree-1; i++ {
				e := node.entries[i]
				keys = append(keys, rune(e.key))
				nextNodes = append(nextNodes, e.node)
			}

			groups = append(groups, keys)
			nextNodes = append(nextNodes, node.entries[node.degree-1].node)
			isLeaf = node.leaf
		}

		context.nodes[next] = nextNodes
		context.current, context.next = next, current
		result = append(result, groups)
	}

	return result
}

func compare(actual [][][]rune, expected [][][]rune) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i, aGroups := range actual {
		eGroups := expected[i]
		if len(aGroups) != len(eGroups) {
			return false
		}

		for j, aKeys := range aGroups {
			eKeys := eGroups[j]
			if len(aKeys) != len(eKeys) {
				return false
			}

			for n, key := range aKeys {
				if key != eKeys[n] {
					return false
				}
			}
		}
	}

	return true
}
