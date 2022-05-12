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
	return true
}
