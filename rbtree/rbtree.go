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

// TODO: need a argument to setup top-down or bottom-up algrithm
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
