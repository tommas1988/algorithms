package btree

type stackEntry struct {
	node  *btreeNode
	index int
}

type nodeStack struct {
	entries []stackEntry
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
	return false
}

func newStack() *nodeStack {
	stack := &nodeStack{
		entries: nil,
		size:    0,
	}
	stack.entries = make([]stackEntry, 1)

	return stack
}

func (ns *nodeStack) clear() {
	ns.size = 0
}

func (ns *nodeStack) push(node *btreeNode, index int) {
	if ns.size >= len(ns.entries) {
		ns.entries = append(ns.entries, stackEntry{node, index})
	} else {
		ns.entries[ns.size] = stackEntry{node, index}
	}
	ns.size++
}

func (ns *nodeStack) pop() (*btreeNode, int) {
	entry := ns.entries[ns.size-1]
	ns.size--
	return entry.node, entry.index
}

func (ns *nodeStack) isEmpty() bool {
	return ns.size == 0
}
