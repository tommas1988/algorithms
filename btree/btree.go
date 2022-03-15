package btree

// TODO: check all the comments make sure they are clear

// a Btree handle, and non of it`s fields are exported
type Btree struct {
	root *btreeNode
	minDegree int
	maxDegree int
}

// A container that hold key, value and a left child node,
// which all keys in this child are less than key
type entry struct {
	key int
	value int
	node *btreeNode
}

type btreeNode struct {
	degree int // degree of btree node
	entries []entry
	leaf bool
}

func New(minDegree int) *Btree {
	return &Btree{root: nil, minDegree: minDegree, maxDegree: 2*minDegree}
}

func (t *Btree) newNode() *btreeNode {
	return &btreeNode{degree: 0, entries: make([]entry, t.maxDegree)}
}

func (t *Btree) Search(key int) (int, bool) {
	if (t.root == nil) {
		return 0, false
	}

	node := t.root
	for true {
		i, ok := node.findKey(key)
		if ok {
			return node.entries[i].value, true
		}

		if node.leaf {
			break
		}

		node = node.entries[i].node
	}

	return 0, false
}

// return entry index that the key is equal to entry.key
// or the smallest entry.key that greater than key
// or the last entry index, which key is greater than all keys in entries
func (n *btreeNode) findKey(key int) (int, bool) {
	i, j := 0, n.degree-2
	for i <= j {
		m := i + (j-i)/2
		e := n.entries[m]
		if key == e.key {
			return m, true
		}

		if key > e.key {
			i = m+1
		} else {
			j = m-1
		}
	}

	return i, false
}

/*
 * Top-down(one-pass) algorithm implementation
 *
 * Travel down the tree searching for the position where key belongs,
 * or update the value if key is already in the tree, and split the each
 * full node that come along the way
 */
func (t *Btree) Insert (key int, value int) {
	// empty tree
	if t.root == nil {
		root := t.newNode()
		root.entries[0].key = key
		root.entries[0].value = value
		root.degree = 2
		root.leaf = true
		t.root = root
		return
	}

	// handle root is full
	// the only way to increase the height of a b-tree
	if t.maxDegree == t.root.degree {
		newRoot := t.newNode()
		newRoot.degree = 1
		newRoot.entries[0].node = t.root
		newRoot.splitChildAt(0, t)

		t.root = newRoot
	}

	node := t.root
	for true {
		i, ok := node.findKey(key)
		if ok {
			// update value
			node.entries[i].value = value
			return
		}

		if node.leaf {
			for j := node.degree-1; j >= i; j-- {
				node.entries[j+1] = node.entries[j]
			}
			node.entries[i].key = key
			node.entries[i].value = value
			node.entries[i].node = nil
			node.degree++
			return
		}

		child := node.entries[i].node
		// split child if child node is full
		if t.maxDegree == child.degree {
			node.splitChildAt(i, t)
			// update i the point to the right child of moved up key
			if key > node.entries[i].key {
				i++
			}
			node = node.entries[i].node
		} else {
			node = child
		}
	}
}

func (n *btreeNode) splitChildAt(idx int, btree *Btree) {
	child := n.entries[idx].node
	// split child at the median key position
	median := child.entries[btree.minDegree-1]

	left := child
	left.degree = btree.minDegree

	right := btree.newNode()
	right.leaf = child.leaf
	right.degree = btree.minDegree

	copy(right.entries, left.entries[btree.minDegree:])

	// move the median key up to the parent
	for i := n.degree-1; i >= idx; i-- {
		n.entries[i+1] = n.entries[i]
	}
	// new entry for the move up key
	n.entries[idx] = entry{key: median.key, value: median.value, node: left}
	// update next entry to point the new created child node
	n.entries[idx+1].node = right
	n.degree++
}

/*
 * Top-down(one-trip) algorithm implementation
 *
 * Traval down the tree deleting the key in the tree,
 * and merge each node down the path with their sibling
 * if this merged node is not full
 */
func (t *Btree) Delete (key int) bool {
	if t.root == nil {
		return false
	}

	node := t.root
	result := false
	for true {
		i, found := node.findKey(key)
		result = found

		// cause current node is guaranteed to have one more key
		// than the minimum degree node, it`s safe to delete the key
		// or break the loop when the key is not in the tree
		if node.leaf {
			if found {
				node.deleteKeyAt(i)
			}
			break
		}

		if found {
			/*
			 * Cases of found the deletion key belongs to internal node
			 *
			 * 1. when child nodes can be merged, which the keys of this merged node
			 * are not greater than btree.maxDegree - 1, then break the loop.
			 * 2. otherwise replace deleted key with the predecessor or successor,
			 * depending on which child is not the minimum degree node, and recursive delete the replacement key.
			 */
			left := node.entries[i].node
			right := node.entries[i+1].node

			if left.degree + right.degree <= t.maxDegree+1 {
				node.mergeChildAt(i, false)
				break
			} else if left.degree > t.minDegree {
				predecessor := left.entries[left.degree-2]
				node.entries[i].key = predecessor.key
				node.entries[i].value = predecessor.value

				node = left
				key = predecessor.key
			} else {
				successor := right.entries[0]
				node.entries[i].key = successor.key
				node.entries[i].value = successor.value

				node = right
				key = successor.key
			}
		} else if node.entries[i].node.degree == t.minDegree {
			/*
			 * Cases of key is not present in internal node and
			 * the child which on the path the key should present is minimum node
			 *
			 * 1. if immediate sibling of child is minimum node, merge child and this sibling
			 * with the key of current node
			 * 2. otherwise move a key from current node down into child, and move a key from
			 * immediate sibling of child up into current node
			 */
			var child, sibling *btreeNode

			isLastChild := i == node.degree-1
			child = node.entries[i].node

			if isLastChild {
				sibling = node.entries[i-1].node
			} else {
				sibling = node.entries[i+1].node
			}

			if sibling.degree == t.minDegree {
				// back to the entry that contains the last key
				if isLastChild {
					i--
				}
				node.mergeChildAt(i, true)
				// update child
				child = node.entries[i].node
			} else if isLastChild {
				node.moveChildKey(i-1, false)
			} else {
				node.moveChildKey(i, true)
			}

			node = child
		} else {
			node = node.entries[i].node
		}
	}


	// when root is empty after deletion process, assign the only child of root as the new root
	// this is the only way to decrease the height of tree
	if t.root.degree == 1 {
		t.root = t.root.entries[0].node
	}

	return result
}

func (n *btreeNode) mergeChildAt(i int, includeParentKey bool) {
	left := n.entries[i].node
	right := n.entries[i+1].node
	if includeParentKey {
		lastEntry := left.entries[left.degree-1]
		lastEntry.key = n.entries[i].key
		lastEntry.value = n.entries[i].value
		left.degree++
	}

	// append right entries to left
	copy(left.entries[left.degree-1:], right.entries[0:right.degree])
	left.degree += (right.degree-1)

	// right child is the merged node, left child pointer will lost after this merge process
	right.degree = left.degree
	right.entries = left.entries

	n.deleteKeyAt(i)
}

func (n *btreeNode) moveChildKey(i int, moveToLeft bool) {
	if moveToLeft {
		child := n.entries[i].node
		sibling := n.entries[i+1].node
		child.entries[child.degree-1].key = n.entries[i].key
		child.entries[child.degree-1].value = n.entries[i].value
		child.entries[child.degree].node = sibling.entries[0].node
		child.degree++

		n.entries[i].key = sibling.entries[0].key
		n.entries[i].value = sibling.entries[0].value

		sibling.deleteKeyAt(0)
	} else {
		child := n.entries[i+1].node
		sibling := n.entries[i].node
		for i := child.degree-1; i >=0; i-- {
			child.entries[i+1] = child.entries[i]
		}
		child.entries[0].key = n.entries[i].key
		child.entries[0].value = n.entries[i].value
		child.entries[0].node = sibling.entries[sibling.degree-1].node
		child.degree++

		n.entries[i].key = sibling.entries[sibling.degree-2].key
		n.entries[i].value = sibling.entries[sibling.degree-2].value

		sibling.degree--
	}
}

func (n *btreeNode) deleteKeyAt(i int) {
	copy(n.entries[i:n.degree], n.entries[i+1:n.degree])
	n.degree--
}
