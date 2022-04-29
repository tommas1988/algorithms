package btree

// TODO: support 2-3 tree
// TODO: wrapper method to get left and right child at node key

type Algorithm int

// TODO: [STY] golang const and iota
const (
	TopDown Algorithm = iota
	BottomUp
)

// a Btree handle, and non of it`s fields are exported
type Btree struct {
	root          *btreeNode
	minDegree     int
	maxDegree     int
	insertHandler func(t *Btree, key int, value int)
	deleteHandler func(t *Btree, key int) bool
}

// A container that hold key, value and a left child node,
// which all keys in this child are less than key
type entry struct {
	key   int
	value int
	node  *btreeNode
}

// TODO: use a cycle to store entries
type btreeNode struct {
	degree  int // degree of btree node
	entries []entry
	leaf    bool
}

func New(minDegree int, alg Algorithm) *Btree {
	btree := &Btree{
		root:      nil,
		minDegree: minDegree,
		maxDegree: 2 * minDegree}

	if alg == TopDown {
		btree.insertHandler = topDownInsertHandler
		btree.deleteHandler = topDownDeleteHandler
	} else if alg == BottomUp {
		panic("Bottom up algorithm is not implement yet")
	}
	return btree
}

func (t *Btree) Search(key int) (int, bool) {
	if t.root == nil {
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

func (t *Btree) Insert(key int, value int) {
	t.insertHandler(t, key, value)
}

func (t *Btree) Delete(key int) bool {
	return t.deleteHandler(t, key)
}

func (t *Btree) newNode() *btreeNode {
	return &btreeNode{degree: 0, entries: make([]entry, t.maxDegree)}
}

// return left child node of entry at i
func (node *btreeNode) left(i int) *btreeNode {
	return node.entries[i].node
}

// return left child node of this entry
func (node *btreeNode) right(i int) *btreeNode {
	return node.entries[i+1].node
}

// return index of entry which the entry.key is equal to the search key,
// or the entry.key is smallest but greater than the search key,
// or the last entry index when search key is greater than all the keys in this node
func (n *btreeNode) findKey(key int) (int, bool) {
	// set initial key index range
	i, j := 0, n.degree-2
	for i <= j {
		m := i + (j-i)/2
		e := n.entries[m]
		if key == e.key {
			return m, true
		}

		if key > e.key {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i, false
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
	for i := n.degree - 1; i >= idx; i-- {
		n.entries[i+1] = n.entries[i]
	}
	// new entry for the move up key
	n.entries[idx] = entry{key: median.key, value: median.value, node: left}
	// update next entry to point the new created child node
	n.entries[idx+1].node = right
	n.degree++
}

/**
 * merge key at i and keys of left child into right child
 * return merged node
 */
func (n *btreeNode) mergeChildAt(i int) *btreeNode {
	left := n.entries[i].node
	right := n.entries[i+1].node
	lastEntry := left.entries[left.degree-1]
	lastEntry.key = n.entries[i].key
	lastEntry.value = n.entries[i].value
	left.degree++

	// append right entries to left
	copy(left.entries[left.degree-1:], right.entries[0:right.degree])
	left.degree += (right.degree - 1)

	// right child is the merged node, left child pointer will lost after this merge process
	right.degree = left.degree
	right.entries = left.entries

	n.deleteKeyAt(i)

	return right
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
		for i := child.degree - 1; i >= 0; i-- {
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
