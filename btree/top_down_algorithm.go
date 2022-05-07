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
		t.newRoot(key, value, nil, nil)
		return
	}

	// handle root is full
	// the only way to increase the height of a b-tree
	if t.maxDegree == t.root.degree {
		key, value, left, right := t.root.split(t)
		t.newRoot(key, value, left, right)
	}

	node := t.root
	for true {
		i, found := node.findKey(key)
		if found {
			// update value
			node.setValue(i, value)
			return
		}

		if node.leaf {
			if key < node.key(i) {
				node.addKey(i, key, value, nil, nil)
			} else {
				node.addKey(i+1, key, value, nil, nil)
			}

			return
		}

		var child *btreeNode
		if key < node.key(i) {
			child = node.left(i)
		} else {
			child = node.right(i)
			i++ // increment i for add key process when child is full
		}

		// split child if child node is full
		if t.maxDegree == child.degree {
			k, v, l, r := child.split(t)
			node.addKey(i, k, v, l, r)
			if key < node.key(i) {
				node = node.left(i)
			} else {
				node = node.right(i)
			}
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

	var keyIdx int
	var found bool
	node := t.root
	for true {
		if !found {
			keyIdx, found = node.findKey(key)
		}

		// Since current leaf node is guaranteed to have one more key
		// than the minimum degree node, it`s safe to delete the key
		// or just break the loop when the key is not in the tree
		if node.leaf {
			if found {
				node.removeKey(keyIdx)
			}
			break
		}

		left := node.left(keyIdx)
		right := node.right(keyIdx)

		if left.degree == t.minDegree && right.degree == t.minDegree {
			mergeNode := node.mergeChild(keyIdx)
			node.removeKey(keyIdx)
			node = mergeNode

			keyIdx = t.minDegree - 1 // update delete key index
			continue
		}

		var child *btreeNode
		if found {
			// replace with preprocessor and delete preprocessor recursively
			child = left
			lastKeyIdx := child.degree - 2 // last key
			node.setKey(keyIdx, child.key(lastKeyIdx))
			node.setValue(keyIdx, child.value(lastKeyIdx))
		} else if key < node.key(keyIdx) {
			child = left
		} else {
			child = right
		}

		if child.degree == t.minDegree {
			if child == left {
				node.moveChildKey(keyIdx, toLeft)
			} else {
				node.moveChildKey(keyIdx, toRight)
			}
		}

		if found {
			// update delete key index
			keyIdx = child.degree - 2
		}

		node = child
	}

	// when root is empty after deletion process, assign the only child of root as the new root
	// this is the only way to decrease the height of tree
	if t.root.degree == 1 {
		t.root = t.root.left(0)
	}

	return found
}
