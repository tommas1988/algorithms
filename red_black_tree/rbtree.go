package red_black_tree

type color int
const (
	red color = iota
	black
)

// TODO: or RBTree ?
type RedBlackTree struct {
	root *node
	nilNode *node
}

type node struct {
	key int
	value int
	left *node
	right *node
	color color
}

// TODO: need a argument to setup top-down or bottom-up algrithm
func New() *RedBlackTree {
	return &RedBlackTree{
		nilNode: &node{
			color: black,
		}}
}

func (rbtree *RedBlackTree) Search(key int) (int, bool) {
	node := rbtree.root
	// TODO: use nil(sentinel) node ?
	for node != rbtree.nilNode {
		if key == node.key {
			return node.value, true
		} else if key > node.key {
			node = node.right
		} else {
			node = node.left
		}
	}

	return 0, false
}

// Top-down algorithm
func (rbtree *RedBlackTree) Insert(key int, value int) {
	if rbtree.root == nil {
		root := &node{key, value, rbtree.nilNode, rbtree.nilNode, black}

		// make root as child of nil node
		rbtree.nilNode.left = root
		rbtree.root = root
		return
	}

	// declare parent, grandparent and greate grandparent node
	// that will use in rebalance process
	var currNode, p, gp, ggp *node
	currNode = rbtree.root
	p = rbtree.nilNode
	for currNode != rbtree.nilNode {
		// update
		if key == currNode.key {
			currNode.value = value
			return
		}

		if currNode.left.color == red && currNode.right.color == red {
			// TODO: expaine why gp and ggp get value when process node split

			// split full node
			currNode.color = red
			currNode.left.color = black
			currNode.right.color = black

			if p.color == red {
				gp, p, currNode = rebalance(ggp, gp, p, currNode)
			}
		}

		ggp, gp, p = gp, p, currNode
		if key < currNode.key {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}
	}

	// TODO: why insert red node
	currNode = &node{key, value, rbtree.nilNode, rbtree.nilNode, red}
	if key < p.key {
		p.left = currNode
	} else {
		p.right = currNode
	}

	if p.color == red {
		rebalance(ggp, gp, p, currNode)
	}

	// reset root and it`s color
	rbtree.root = rbtree.nilNode.left
	rbtree.root.color = black
}

// compare with 2-3-4 B-tree
func (rbtree *RedBlackTree) Delete(key int) bool {
	if rbtree.root == nil {
		return false
	}

	var (
		currNode *node
		p *node
		gp *node
	)
	var deletedNode *node = nil

	if rbtree.root.key == key {
		deletedNode = rbtree.root
		// find preprocessor
		currNode = rbtree.root.left
	} else if rbtree.root.key < key {
		currNode = rbtree.root.left
	} else {
		currNode = rbtree.root.right
	}
	p = rbtree.root
	gp = rbtree.nilNode

	for currNode != rbtree.nilNode {
		if currNode.color == black &&
			currNode.left.color == black &&
			currNode.right.color == black {
			// minimum node, merge with parent or sibling
			var sibling *node
			if p.left == currNode {
				sibling = p.right
				if sibling.color == red {
					// sibling is in up level, refind the same level sibling
					if gp.left == p {
						gp.left = leftRotate(p)
					} else {
						gp.right = leftRotate(p)
					}
					continue
				} else if sibling.left.color == black && sibling.right.color == black {
					// sibling is also a minimum node, merge them with parent
					p.color = black
					currNode.color = red
					sibling.color = red
				} else {
					// sibling is 3 degree node, move up a node to parent and merge with old parent
					if sibling.right.color == black {
						p.right = rightRotate(sibling)
						sibling.color = red
						sibling.left.color = black
					}

					if gp.left == p {
						gp.left = leftRotate(p)
					} else {
						gp.right = leftRotate(p)
					}
					currNode.color = red
				}
			} else {
				sibling = p.left
				if sibling.color == red {
					if gp.left == p {
						gp.left = rightRotate(p)
					} else {
						gp.right = rightRotate(p)
					}
					continue
				} else if sibling.left.color == black && sibling.right.color == black {
					p.color = black
					currNode.color = red
					sibling.color = red
				} else {
					if sibling.left.color == black {
						p.left = leftRotate(sibling)
						sibling.color = red
						sibling.right.color = black
					}

					if gp.left == p {
						gp.left = rightRotate(p)
					} else {
						gp.right = rightRotate(p)
					}
					currNode.color = red
				}
			}
		}

		gp = p
		p = currNode
		if currNode.key == key {
			deletedNode = currNode
			// find preprocessor
			currNode = currNode.left
		} else if deletedNode != nil {
			// find preprocessor
			if currNode.right == rbtree.nilNode {
				break
			}
			currNode = currNode.right
		} else if (key < currNode.key) {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}

	}

	if deletedNode == nil {
		return false
	}

	deletedNode.key = currNode.key
	deletedNode.value = currNode.value
	if p.right == currNode {
		p.right = rbtree.nilNode
	} else {
		p.left = rbtree.nilNode
	}
	return true
}

func rebalance(ggp, gp, p, c *node) (_gp, _p, _c *node) {
	if gp.left == p {
		if p.right == c {
			gp.left = leftRotate(p)
			c, p = p, c
		}

		gp.color = red
		p.color = black
		if (ggp.left == gp) {
			ggp.left = rightRotate(gp)
		} else {
			ggp.right = rightRotate(gp)
		}
		_gp, _p, _c = ggp, p, c
	} else {
		if p.left == c {
			gp.right = rightRotate(p)
			c, p = p, c
		}

		gp.color = red
		p.color = black
		if (ggp.left == gp) {
			ggp.left = leftRotate(gp)
		} else {
			ggp.right = leftRotate(gp)
		}
		_gp, _p, _c = ggp, p, c
	}

	return _gp, _p, _c
}

/**
 * \
 */
func leftRotate(root *node) *node {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	return newRoot;
}

/**
 * /
 */
func rightRotate(root *node) *node {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root;
	return newRoot;
}
