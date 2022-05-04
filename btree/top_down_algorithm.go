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
		root.setKey(0, key)
		root.setValue(0, value)
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
		key, value, left, right := t.root.split(t)
		newRoot.addKey(0, key, value, left, right)
		t.root = newRoot
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

	deleteNode := struct {
		node  *btreeNode
		index int
	}{}
	node := t.root
	for true {
		var i int
		var found bool
		if deleteNode.node != nil {
			i = node.degree - 2
			found = true
		} else {
			i, found = node.findKey(key)
		}

		// Since current leaf node is guaranteed to have one more key
		// than the minimum degree node, it`s safe to delete the key
		// or just break the loop when the key is not in the tree
		if node.leaf {
			if deleteNode.node != nil {
				deleteNode.node.setKey(deleteNode.index, node.key(i))
				deleteNode.node.setValue(deleteNode.index, node.value(i))
			}

			if found {
				node.deleteKey(i)
			}
			break
		}

		left := node.left(i)
		right := node.right(i)

		if left.degree == t.minDegree && right.degree == t.minDegree {
			node = node.mergeChild(i)
			continue
		}

		var child *btreeNode
		if deleteNode.node != nil || found {
			// preprocessor
			child = left

			if found {
				deleteNode.node = node
				deleteNode.index = i
			}
		} else if key < node.key(i) {
			child = left
		} else {
			child = right
		}

		if child.degree == t.minDegree {
			// replace key from sibling with key at i of node, and merge replaced key into child node
			if child == left {
				child.appendKey(node.key(i), node.value(i), right.left(0))
				node.setKey(i, right.key(i))
				node.setValue(i, right.value(i))
				right.deleteKey(0)
			} else {
				lastKeyIndex := left.degree - 2
				child.addKey(0, node.key(i), node.value(i), left.right(lastKeyIndex), node.left(0))
				node.setKey(i, left.key(lastKeyIndex))
				node.setValue(i, left.value(lastKeyIndex))
				left.deleteKey(lastKeyIndex)
			}
		}

		node = child
	}

	// when root is empty after deletion process, assign the only child of root as the new root
	// this is the only way to decrease the height of tree
	if t.root.degree == 1 {
		t.root = t.root.left(0)
	}

	return deleteNode.node != nil
}
