package btree

// TODO: support 2-3 tree

type Algorithm int
type moveDirection int

// TODO: [STY] golang const and iota
const (
	TopDown Algorithm = iota
	BottomUp
)

const (
	toLeft moveDirection = iota
	toRight
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
		btree.insertHandler = bottomUpInsertHanlder
		btree.deleteHandler = bottomUpDeleteHandler
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

func (t *Btree) newRoot(key int, value int, left *btreeNode, right *btreeNode) {
	leaf := false
	if left == nil {
		leaf = true
	}
	root := &btreeNode{
		degree:  2,
		entries: make([]entry, t.maxDegree),
		leaf:    leaf,
	}

	root.entries[0].key = key
	root.entries[0].value = value
	root.entries[0].node = left
	root.entries[1].node = right

	t.root = root
}

/**
 * return index of entry which entry.key is equal to the search key,
 * or of entry that entry.key is smallest but greater than the search key,
 * or of last entry that contains key when search key is greater than all the keys in this node
 */
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

	if i == n.degree-1 {
		i = i - 1
	}

	return i, false
}

func (n *btreeNode) key(i int) int {
	return n.entries[i].key
}

func (n *btreeNode) value(i int) int {
	return n.entries[i].value
}

func (n *btreeNode) setKey(i int, key int) {
	n.entries[i].key = key
}

func (n *btreeNode) setValue(i int, value int) {
	n.entries[i].value = value
}

func (n *btreeNode) left(i int) *btreeNode {
	return n.entries[i].node
}

func (n *btreeNode) right(i int) *btreeNode {
	return n.entries[i+1].node
}

func (n *btreeNode) addKey(idx int, key int, value int, left *btreeNode, right *btreeNode) {
	// make a room for added entry
	for i := n.degree - 1; i >= idx; i-- {
		n.entries[i+1] = n.entries[i]
	}
	n.entries[idx].key = key
	n.entries[idx].value = value
	n.entries[idx].node = left
	n.degree++

	n.entries[idx+1].node = right
}

func (n *btreeNode) appendKey(key int, value int, right *btreeNode) {
	i := n.degree - 1
	n.entries[i].key = key
	n.entries[i].value = value
	n.entries[i+1].node = right
	n.degree++
}

func (n *btreeNode) removeKey(i int) {
	if i < n.degree-2 {
		copy(n.entries[i:n.degree], n.entries[i+1:n.degree])
	}
	n.degree--
}

func (n *btreeNode) split(btree *Btree) (key int, value int, left *btreeNode, right *btreeNode) {
	// Entry that split child entries. And being merged up into current node
	splitEntry := n.entries[btree.minDegree-1]

	left = n
	left.degree = btree.minDegree

	right = btree.newNode()
	right.leaf = n.leaf
	right.degree = btree.minDegree

	copy(right.entries, left.entries[btree.minDegree:])

	return splitEntry.key, splitEntry.value, left, right
}

/**
 * merge key at i and keys of left child into right child
 * return merged node
 */
func (n *btreeNode) mergeChild(idx int) *btreeNode {
	left := n.left(idx)
	right := n.right(idx)

	var mergeNode *btreeNode
	if idx == n.degree-2 {
		// merge right into left
		left.setKey(left.degree-1, n.key(idx))
		left.setValue(left.degree-1, n.value(idx))

		// append right entries to left
		copy(left.entries[left.degree:], right.entries[0:right.degree])
		left.degree += right.degree

		mergeNode = left
	} else {
		// merge left into right
		c := left.degree // number of left keys and parent key
		for i := 0; i < c; i++ {
			right.entries[c+i] = right.entries[i]
		}
		copy(right.entries[0:left.degree], left.entries[0:left.degree])
		right.setKey(left.degree-1, n.key(idx))
		right.setValue(left.degree-1, n.value(idx))
		right.degree += left.degree

		mergeNode = right
	}

	return mergeNode
}

/**
 * replace key from a child to its sibling
 */
func (n *btreeNode) moveChildKey(i int, dir moveDirection) {
	left := n.left(i)
	right := n.right(i)
	if dir == toLeft {
		left.appendKey(n.key(i), n.value(i), right.left(0))
		n.setKey(i, right.key(0))
		n.setValue(i, right.value(0))
		right.removeKey(0)
	} else {
		lastKeyIdx := left.degree - 2
		right.addKey(0, n.key(i), n.value(i), left.right(lastKeyIdx), right.left(0))
		n.setKey(i, left.key(lastKeyIdx))
		n.setValue(i, left.value(lastKeyIdx))
		left.removeKey(lastKeyIdx)
	}
}
