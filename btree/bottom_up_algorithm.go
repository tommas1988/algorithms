package btree

type keyPos struct {
	node  *btreeNode
	index int
}

type nodeStack struct {
	entries []keyPos
	size    int
}

func bottomUpInsertHanlder(t *Btree, key int, value int) {
	// empty tree
	if t.root == nil {
		t.newRoot(key, value, nil, nil)
		return
	}

	node := t.root
	stack := newStack()
	for true {
		i, found := node.findKey(key)
		if found {
			// update value
			node.setValue(i, value)
			return
		}

		if node.degree == t.maxDegree {
			stack.push(node, i)
		} else {
			stack.clear()
			stack.push(node, i)
		}

		if node.leaf {
			break
		}

		if key < node.key(i) {
			node = node.left(i)
		} else {
			node = node.right(i)
		}
	}

	var left, right *btreeNode = nil, nil
	for true {
		node, idx := stack.pop()
		isFullNode := node.degree == t.maxDegree
		if node.degree == t.maxDegree {
			k, v, l, r := node.split(t)
			if key < k {
				node = l
			} else {
				node = r
				idx = idx - t.minDegree
			}

			if key < node.key(idx) {
				node.addKey(idx, key, value, left, right)
			} else {
				node.addKey(idx+1, key, value, left, right)
			}
			key, value, left, right = k, v, l, r
		} else {
			if key < node.key(idx) {
				node.addKey(idx, key, value, left, right)
			} else {
				node.addKey(idx+1, key, value, left, right)
			}
		}

		if stack.isEmpty() {
			if isFullNode {
				t.newRoot(key, value, left, right)
			}

			break
		}
	}
}

func bottomUpDeleteHandler(t *Btree, key int) bool {
	if t.root == nil {
		return false
	}

	var deleteKeyPos keyPos
	stack := newStack()
	node := t.root
	found := false
	var keyIdx int
	for true {
		var child *btreeNode
		if found {
			// path to the preprocessor
			keyIdx = node.degree - 2
			child = node.right(keyIdx)
		} else {
			keyIdx, found = node.findKey(key)
			if found {
				deleteKeyPos.node = node
				deleteKeyPos.index = keyIdx
				// path to the preprocessor
				child = node.left(keyIdx)
			} else if key < node.key(keyIdx) {
				child = node.left(keyIdx)
			} else {
				child = node.right(keyIdx)
			}
		}

		if node.degree <= t.minDegree {
			stack.push(node, keyIdx)
		} else {
			stack.clear()
			stack.push(node, keyIdx)
		}

		if node.leaf {
			if found {
				// replace delete key with preprocessor
				deleteKeyPos.node.setKey(deleteKeyPos.index, node.key(keyIdx))
				deleteKeyPos.node.setValue(deleteKeyPos.index, node.value(keyIdx))
			}
			break
		}

		node = child
	}

	if !found {
		return false
	}

	for !stack.isEmpty() {
		node, idx := stack.pop()
		if node.degree > t.minDegree {
			node.removeKey(idx)
		} else {
			var sibling *btreeNode
			parent, i := stack.peek()
			if parent.left(i) == node {
				sibling = parent.right(i)
			} else {
				sibling = parent.left(i)
			}

			node.removeKey(idx)

			if sibling.degree == t.minDegree {
				mergeNode := parent.mergeChild(i)
				if t.root == parent {
					t.root = mergeNode
					break
				}
			} else {
				if parent.left(i) == node {
					parent.moveChildKey(i, toLeft)
				} else {
					parent.moveChildKey(i, toRight)
				}
				break
			}
		}
	}

	return true
}

func newStack() *nodeStack {
	stack := &nodeStack{
		entries: nil,
		size:    0,
	}
	stack.entries = make([]keyPos, 1)

	return stack
}

func (ns *nodeStack) clear() {
	ns.size = 0
}

func (ns *nodeStack) push(node *btreeNode, index int) {
	if ns.size >= len(ns.entries) {
		ns.entries = append(ns.entries, keyPos{node, index})
	} else {
		ns.entries[ns.size] = keyPos{node, index}
	}
	ns.size++
}

func (ns *nodeStack) pop() (*btreeNode, int) {
	entry := ns.entries[ns.size-1]
	ns.size--
	return entry.node, entry.index
}

func (ns *nodeStack) peek() (*btreeNode, int) {
	entry := ns.entries[ns.size-1]
	return entry.node, entry.index
}

func (ns *nodeStack) isEmpty() bool {
	return ns.size == 0
}
