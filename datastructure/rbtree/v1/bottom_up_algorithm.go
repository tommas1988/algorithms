package rbtree

func bottomUpInsertHandler(rbtree *RedBlackTree, key int, value int) {
	if rbtree.root == rbtree.nilNode {
		root := &node{key, value, rbtree.nilNode, rbtree.nilNode, black}

		// make root as child of nil node
		rbtree.nilNode.left = root
		rbtree.root = root
		return
	}

	nodePath := []*node{rbtree.nilNode}
	treeNode := rbtree.root
	insertNode := &node{key, value, rbtree.nilNode, rbtree.nilNode, red}
	for true {
		if treeNode.key == key {
			// update
			treeNode.value = value
			return
		}

		nodePath = append(nodePath, treeNode)

		var child *node
		if key < treeNode.key {
			child = treeNode.left
			if child == rbtree.nilNode {
				treeNode.left = insertNode
				break
			}
		} else {
			child = treeNode.right
			if child == rbtree.nilNode {
				treeNode.right = insertNode
				break
			}
		}
		treeNode = child
	}

	// back to root but not include root
	for i := len(nodePath) - 1; i > 1; {
		p := nodePath[i]

		if p.color == black {
			break
		}

		gp := nodePath[i-1]
		if gp.left.color == red && gp.right.color == red {
			gp.color = red
			gp.left.color = black
			gp.right.color = black
			i = i - 2
			continue
		}

		ggp := nodePath[i-2]
		var c *node
		if p.left.color == red {
			c = p.left
		} else {
			c = p.right
		}

		rebalance(ggp, gp, p, c)
		break
	}

	// reset root
	rbtree.root = rbtree.nilNode.left
	rbtree.root.color = black
}

func bottomUpDeleteHandler(rbtree *RedBlackTree, key int) bool {
	if rbtree.root == rbtree.nilNode {
		return false
	}

	var deleteNode *node
	nodePath := []*node{rbtree.nilNode}
	treeNode := rbtree.root
	for treeNode != rbtree.nilNode {
		nodePath = append(nodePath, treeNode)

		if treeNode.key == key {
			deleteNode = treeNode
			// path to preprocessor
			treeNode = treeNode.left
		} else if deleteNode != nil {
			// path to preprocessor
			treeNode = treeNode.right
		} else if key < treeNode.key {
			treeNode = treeNode.left
		} else {
			treeNode = treeNode.right
		}
	}

	if deleteNode == nil {
		return false
	}

	// When left child of delete node is nil node, replace delete node with its right child.
	// Otherwise replace delete node with preprocessor, then delete preprocessor.
	lastNodeIdx := len(nodePath) - 1
	lastNode := nodePath[lastNodeIdx]
	var replaceNode *node
	if lastNode == deleteNode {
		replaceNode = lastNode.right
	} else {
		deleteNode.key = lastNode.key
		deleteNode.value = lastNode.value
		replaceNode = lastNode.left
	}

	parent := nodePath[lastNodeIdx-1]
	if parent.left == lastNode {
		parent.left = replaceNode
	} else {
		parent.right = replaceNode
	}

	// reset root
	rbtree.root = rbtree.nilNode.left

	if lastNode.color == red {
		return true
	} else if replaceNode.color == red {
		replaceNode.color = black
		return true
	}

	// fix up red black tree
	nodePath[lastNodeIdx] = replaceNode

	for i := lastNodeIdx; i > 1; {
		c := nodePath[i] // must be black node
		p := nodePath[i-1]
		gp := nodePath[i-2]
		var sibling *node
		if p.left == c {
			sibling = p.right
		} else {
			sibling = p.left
		}

		if p.color == black &&
			sibling.color == black &&
			sibling.left.color == black &&
			sibling.right.color == black {
			// merge parent and sibling into a 2-3-4 tree node
			redifyNode(gp, p, c)
			// back to child 2-3-4 tree node or revert nil node color
			c.color = black
			i = i - 1
		} else {
			redifyNode(gp, p, c)
			// back to child 2-3-4 tree node or revert nil node color
			c.color = black
			break
		}
	}

	// reset root
	rbtree.root = rbtree.nilNode.left

	return true
}
