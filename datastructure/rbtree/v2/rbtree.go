package rbtree

type color int

const (
	red color = iota
	black
)

type node struct {
	key    int
	value  int
	parent *node
	left   *node
	right  *node
	color  color
}

type RBTree struct {
	root     *node
	sentinel *node
}

func New() *RBTree {
	sentinel := node{
		color: red,
	}
	return &RBTree{
		root: &sentinel,
	}
}

func (t *RBTree) Search(key int) (value int, ok bool) {
	if n := t.search(key); n != nil {
		return n.value, true
	}
	return 0, false
}

func (t *RBTree) Insert(key int, value int) bool {
	if t.root == t.sentinel {
		t.root = &node{
			key:    key,
			value:  value,
			parent: t.sentinel,
			left:   t.sentinel,
			right:  t.sentinel,
			color:  black,
		}
		t.sentinel.left = t.root
		return true
	}

	var parent *node
	n := t.root
	for n != t.sentinel {
		if key == n.key {
			return false
		}

		parent = n
		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}

	insertNode := &node{
		key:    key,
		value:  value,
		parent: parent,
		left:   t.sentinel,
		right:  t.sentinel,
		color:  red,
	}

	if key < parent.key {
		parent.left = insertNode
	} else {
		parent.right = insertNode
	}

	t.fixupNodeColor(insertNode)

	// update root in case of root changed by root rotation
	t.root = t.sentinel.left

	return true
}

func (t *RBTree) Update(key int, value int) bool {
	if n := t.search(key); n != nil {
		n.value = value
		return true
	}
	return false
}

func (t *RBTree) Delete(key int) bool {
	if n := t.search(key); n != nil {
		parent := n.parent
		if n.left == t.sentinel {
			if n == parent.left {
				parent.left = n.right
			} else {
				parent.right = n.right
			}
		} else if n.right == t.sentinel {
			if n == parent.left {
				parent.left = n.left
			} else {
				parent.right = n.left
			}
		} else {
			predecessor := t.predecessor(n)
			n.key = predecessor.key
			n.value = predecessor.value

			if n.color == red {
				n.color = predecessor.color
			}

			if n.color == black && predecessor.color == black {
				n = predecessor
				n.color = red
				for n.parent != t.sentinel {
					if n.parent.color == red {
						n.parent.color = black
						break
					} else {
						if n.parent.left == n {
							if n.parent.right.color == red {
								n.parent.leftRotate()
								n.parent.color = black
								break
							} else {
								n = n.parent
							}
						} else {
							if n.parent.left.color == red {
								n.parent.rightRotate()
								n.parent.color = black
								break
							} else {
								n = n.parent
							}
						}
					}
				}
				// root will be red, if all the nodes up to the root are signle node from B-Tree perspective
				t.root.color = black
			}

			// remove predecessor
			if predecessor.parent.left == predecessor {
				predecessor.left = t.sentinel
			} else {
				predecessor.right = t.sentinel
			}
		}

		return true
	}

	return false
}

func (t *RBTree) search(key int) *node {
	n := t.root
	for n != t.sentinel {
		if key == n.key {
			return n
		}

		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}

	return nil
}

func (t *RBTree) fixupNodeColor(n *node) {
	if n.color == black || n.parent.color == black {
		return
	}

	parent := n.parent
	gp := parent.parent
	if gp.left.color != gp.right.color {
		if gp.left == parent {
			if n.key > parent.key {
				parent.leftRotate()
			}
			gp.rightRotate()
		} else {
			if n.key < parent.key {
				parent.rightRotate()
			}
			gp.leftRotate()
		}
	} else {
		gp.color = red
		gp.left.color, gp.right.color = black, black

		if gp.parent == t.sentinel {
			gp.color = black
		}

		if gp.parent.color == red {
			t.fixupNodeColor(gp)
		}
	}
}

func (t *RBTree) predecessor(n *node) *node {
	n = n.left
	for n.right != t.sentinel {
		n = n.right
	}
	return n
}

func (n *node) leftRotate() {
	n.leftRotateV1()
}

func (n *node) rightRotate() {
	n.rightRotateV1()
}

// Implement rotation by changing node pointer
func (n *node) leftRotateV1() {
	parent := n.parent
	gp := parent.parent
	if gp.left == parent {
		gp.left = n
	} else {
		gp.right = n
	}
	n.parent = gp

	n.right, parent.left = parent, n.right
	parent.parent = n
	n.right.parent = parent

	n.color, parent.color = parent.color, n.color
}

func (n *node) rightRotateV1() {
	parent := n.parent
	gp := parent.parent
	if gp.left == parent {
		gp.left = n
	} else {
		gp.right = n
	}
	n.parent = gp

	n.left, parent.right = parent, n.left
	parent.parent = n
	n.left.parent = parent

	n.color, parent.color = parent.color, n.color
}

// Implement rotation by switching fields with parent node
func (n *node) leftRotateV2() {
	parent := n.parent
	n.key, parent.key = parent.key, n.key
	n.value, parent.value = parent.value, n.value

	parent.left = n.left
	n.left.parent = parent
	n.left = n.right
	n.right = parent.right
	parent.right.parent = n
	parent.right = n
}
