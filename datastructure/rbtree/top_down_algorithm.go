package rbtree

// Top-down algorithm
func topDownInsertHandler(rbtree *RedBlackTree, key int, value int) {
	if rbtree.root == rbtree.nilNode {
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
func topDownDeleteHandler(rbtree *RedBlackTree, key int) bool {
	if rbtree.root == rbtree.nilNode {
		return false
	}

	var (
		currNode *node
		p        *node
		gp       *node
	)
	var deleteNode *node = nil

	p = rbtree.root
	gp = rbtree.nilNode
	if rbtree.root.key == key {
		deleteNode = rbtree.root
		// find preprocessor
		currNode = rbtree.root.left
	} else if key < rbtree.root.key {
		currNode = rbtree.root.left
	} else {
		currNode = rbtree.root.right
	}

	for currNode != rbtree.nilNode {
		if currNode.color == black &&
			currNode.left.color == black &&
			currNode.right.color == black {
			// minimum node, merge with parent or sibling
			redifyNode(gp, p, currNode)
		}

		var child *node
		if currNode.key == key {
			deleteNode = currNode
			// find preprocessor
			if currNode.left == rbtree.nilNode {
				break
			}
			child = currNode.left
		} else if deleteNode != nil {
			// find preprocessor
			if currNode.right == rbtree.nilNode {
				break
			}
			child = currNode.right
		} else if key < currNode.key {
			child = currNode.left
		} else {
			child = currNode.right
		}

		gp = p
		p = currNode
		currNode = child
	}

	if deleteNode == nil {
		return false
	}

	if deleteNode.left == rbtree.nilNode {
		// force right child to black, in case of red node
		deleteNode.right.color = black
		if deleteNode == rbtree.root {
			p = rbtree.nilNode
		}
		if p.left == deleteNode {
			p.left = deleteNode.right
		} else {
			p.right = deleteNode.right
		}
	} else {
		deleteNode.key = currNode.key
		deleteNode.value = currNode.value
		if p.right == currNode {
			// right child is nil node, but left child may be a red node
			p.right = currNode.left
		} else {
			p.left = currNode.left
		}
	}

	// reset root
	rbtree.root = rbtree.nilNode.left

	return true
}
