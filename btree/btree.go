package btree

// TODO: support 2-3 tree

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
	copy(n.entries[i:n.degree], n.entries[i+1:n.degree])
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
func (n *btreeNode) mergeChild(i int) *btreeNode {
	left := n.left(i)
	right := n.right(i)

	// merge key into child
	lastEntry := left.entries[left.degree-1]
	lastEntry.key = n.entries[i].key
	lastEntry.value = n.entries[i].value
	left.degree++

	// append right entries to left
	copy(left.entries[left.degree-1:], right.entries[0:right.degree])
	left.degree += (right.degree - 1)

	n.removeKey(i)

	// right child is the merged node, left child pointer will lost after this merge process
	right.degree = left.degree
	right.entries = left.entries
	return right

}
