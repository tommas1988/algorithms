package rbtree

type Algorithm int
type color int

const (
	red color = iota
	black
)

const (
	TopDown Algorithm = iota
	BottomUp
)

// TODO: or RBTree ?
type RedBlackTree struct {
	root          *node
	nilNode       *node
	insertHanlder func(rbtree *RedBlackTree, key int, value int)
	deleteHandler func(rbtree *RedBlackTree, key int) bool
}

type node struct {
	key   int
	value int
	left  *node
	right *node
	color color
}

func New(alg Algorithm) *RedBlackTree {
	nilNode := node{
		color: black,
	}
	rbtree := RedBlackTree{
		root:    &nilNode,
		nilNode: &nilNode,
	}

	if TopDown == alg {
		rbtree.insertHanlder = topDownInsertHandler
		rbtree.deleteHandler = topDownDeleteHandler
	} else {
		rbtree.insertHanlder = bottomUpInsertHandler
		rbtree.deleteHandler = bottomUpDeleteHandler
	}

	return &rbtree
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

func (rbtree *RedBlackTree) Insert(key int, value int) {
	rbtree.insertHanlder(rbtree, key, value)
}

func (rbtree *RedBlackTree) Delete(key int) bool {
	return rbtree.deleteHandler(rbtree, key)
}

// TODO: need comments
// parent and current are both red node, need rebalance
func rebalance(ggp, gp, p, c *node) (_gp, _p, _c *node) {
	if gp.left == p {
		if p.right == c {
			gp.left = leftRotate(p)
			c, p = p, c
		}

		gp.color = red
		p.color = black
		if ggp.left == gp {
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
		if ggp.left == gp {
			ggp.left = leftRotate(gp)
		} else {
			ggp.right = leftRotate(gp)
		}
		_gp, _p, _c = ggp, p, c
	}

	return _gp, _p, _c
}

// color current node to red
func redifyNode(gp *node, p *node, c *node) {
	if p.left == c {
		sibling := p.right
		if sibling.color == red {
			// sibling belongs to up level B-tree node
			if gp.left == p {
				gp.left = leftRotate(p)
			} else {
				gp.right = leftRotate(p)
			}
			sibling.color = black
			p.color = red
			sibling = p.right
		}

		if sibling.left.color == black && sibling.right.color == black {
			// merge B-tree node
			p.color = black
			c.color = red
			sibling.color = red
		} else {
			// move left most node up to parent B-tree node,
			// move origin parent down to current B-tree node
			if sibling.left.color == red {
				// recolor in following rotate process
				p.right = rightRotate(sibling)
				sibling = p.right
			}

			if gp.left == p {
				gp.left = leftRotate(p)
			} else {
				gp.right = leftRotate(p)
			}
			sibling.color = p.color
			sibling.right.color = black
			p.color = black
			c.color = red
		}
	} else {
		sibling := p.left
		if sibling.color == red {
			// sibling belongs to up level B-tree node
			if gp.left == p {
				gp.left = rightRotate(p)
			} else {
				gp.right = rightRotate(p)
			}
			sibling.color = black
			p.color = red
			sibling = p.left
		}

		if sibling.left.color == black && sibling.right.color == black {
			// merge B-tree node
			p.color = black
			c.color = red
			sibling.color = red
		} else {
			// move right most node up to parent B-tree node,
			// move origin parent down to current B-tree node
			if sibling.right.color == red {
				// recolor in following rotate process
				p.left = leftRotate(sibling)
				sibling = p.left
			}

			if gp.left == p {
				gp.left = rightRotate(p)
			} else {
				gp.right = rightRotate(p)
			}
			sibling.color = p.color
			sibling.left.color = black
			p.color = black
			c.color = red
		}
	}
}

/**
 * \
 */
func leftRotate(root *node) *node {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	return newRoot
}

/**
 * /
 */
func rightRotate(root *node) *node {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	return newRoot
}
