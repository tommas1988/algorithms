package btree

/*
 * Top-down(one-pass) algorithm implementation
 *
 * Travel down the tree searching for the position where key belongs,
 * or update the value if key is already in the tree, and split the each
 * full node that come along the way
 */
func topDownInsertHandler(t *Btree, key int, value int) {
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
			for j := node.degree - 1; j >= i; j-- {
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

/*
 * Top-down(one-trip) algorithm implementation
 *
 * Traval down the tree deleting the key in the tree,
 * and merge each node down the path with their sibling
 * if this merged node is not full
 */
func topDownDeleteHandler(t *Btree, key int) bool {
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
			 * 1. when child nodes are both min degree node, mrege key and child nodes,
			 * and recursive delete the key in the merged node
			 * 2. otherwise replace deleted key with the predecessor or successor,
			 * depending on which child is not the minimum degree node, and recursive delete the replacement key.
			 */
			left := node.entries[i].node
			right := node.entries[i+1].node

			// TODO: is this right to delete node[i] directly
			if left.degree == t.minDegree && right.degree == t.minDegree {
				node = node.mergeChildAt(i)
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
				child = node.mergeChildAt(i)
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
