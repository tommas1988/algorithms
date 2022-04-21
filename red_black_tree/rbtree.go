package red_black_tree

type color int
const (
	red color iota
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

func New() *RedBlackTree {
	return &RedBlackTree{
		nilNode: &node{
			color: black
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
		root := &node{
			key: key,
			value: value,
			color: black,
		}
		// make root as child of nil node
		rbtree.nilNode.left = rbtree.nilNode.left = root
		rbtree.root = root

		return
	}

	// declare parent, grandparent and greate grandparent node
	// that will use in rebalance process
	var p, gp, ggp *node
	p = rbtree.nilNode
	node := rbtree.root
	for node != rbtree.nilNode {
		// update
		if key == node.key {
			node.value = value
			return
		}

		if node.left.color == red || node.right.color == red {
			// TODO: expaine why gp and ggp get value when process node split

			// split full node
			node.color = red
			node.left.color = node.right.color = black

			if p.color = red {
				gp, p, node = rebalance(ggp, gp, p, node)
			}
		}

		ggp, gp, p = gp, p, node
		node = key < node.key ? node.left : node.right
	}

	// TODO: why insert red node
	node = &node{key, value, rbtree.nilNode, rbtree.nilNode, red}
	if key < p.key {
		p.left = node
	} else {
		p.right = node
	}

	if p.color == red {
		rebalance(ggp, gp, p, node)
	}

	rbtree.root.color = black
}

// compare with 2-3-4 B-tree
func (rbtree *RedBlackTree) Delete(key int) bool {
	if rbtree.root == nil {
		return false
	}

	var (
		node *node,
		p *node,
		gp *node
	)
	var deleteNode *node = nil

	if rbtree.root.key == key {
		deletedNode = rbtree.root
		// find preprocessor
		node = root.left
	} else if rbtree.root.key < key {
		node = root.left
	} else {
		node = root.right
	}
	p = rbtree.root
	gp = rbtree.nilNode

	for node != rbtree.nilNode {
		if node.color == black &&
			node.left.color == black &&
			node.right.color == black {
			// minimum node, merge with parent or sibling
			var sibling *node
			if p.left == node {
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
					node.color = sibling.color = red
				} else {
					// sibling is 3 degree node, move up a node to parent and merge with old parent
					if sibling.right.color == black {
						p.right = rightRotate(sibling)
						sibling.color = red
						sibling.left.color = black
					}

					if gp.left == p {
						gp.left == leftRotate(p)
					} else {
						gp.right == leftRotate(p)
					}
					node.color = red
				}
			} else {
				sbiling = p.left
				if sibling.color == red {
					if gp.left == p {
						gp.left = rightRotate(p)
					} else {
						gp.right = rightRotate(p)
					}
					continue
				} else if sibling.left.color == black && sibling.right.color == black {
					p.color = black
					node.color = sibling.color = red
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
					node.color = red
				}
			}
		}

		gp = p
		p = node
		if node.key == key {
			deleteNode = node
			// find preprocessor
			node = node.left
		} else if deleteNode != nil {
			// find preprocessor
			if node.right = rbtree.nilNode {
				break
			}
			node = node.right
		} else if (key < node.key) {
			node = node.left
		} else {
			node = node.right
		}

	}

	if deleteNode == nil {
		return false
	}

	deleteNode.key = node.key
	deleteNode.value = node.value
	if parent.right = node {
		panret.right = rbtree.nilNode
	} else {
		parent.left = rbtree.nilNode
	}
	return true
}

func rebalance(ggp, gp, p, c *node) (gp, p, c *node) {
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
		gp = ggp
	} else {
		if p.left == c {
			gp.right = rightRotate(p)
			c, p = p, c
		}

		gp.color = red
		p.color = black
		if (ggp.left = gp) {
			ggp.left = leftRotate(gp)
		} else {
			ggp.right = leftRotate(gp)
		}
		gp = ggp
	}

	return gp, p, c
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
