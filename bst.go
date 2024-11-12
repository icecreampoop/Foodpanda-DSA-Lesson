package main

import (
	"errors"
	"fmt"
)

type BinaryNode struct {
	item  string      // to store the data item
	left  *BinaryNode // pointer to point to left node
	right *BinaryNode // pointer to point to right node
}

type BST struct {
	root *BinaryNode
}

func (bst *BST) insertNode(t **BinaryNode, item string) error {

	if *t == nil {
		newNode := &BinaryNode{
			item:  item,
			left:  nil,
			right: nil,
		}
		*t = newNode
		return nil
	}

	if item < (*t).item {
		bst.insertNode(&((*t).left), item)
	} else {
		bst.insertNode(&((*t).right), item)
	}

	return nil
}

func (bst *BST) insert(item string) {
	bst.insertNode(&bst.root, item)

}

func (t *BST) search(root *BinaryNode, item string) (*BinaryNode, error) {
	if root == nil {
		return nil, errors.New("Cannot be found")
	}	

	if root.item == item {
		return root, nil
	}

	if item > root.item {
		return t.search(root.right, item)
	} else {
		return t.search(root.left, item)
	}
}

func (bst *BST) inOrderTraverse(t *BinaryNode) {
	if t != nil {
		bst.inOrderTraverse(t.left)
		fmt.Println(t.item)
		bst.inOrderTraverse(t.right)
	}
}

func (bst *BST) preOrderTraverse(t *BinaryNode) {
	if t != nil {
		fmt.Println(t.item)
		bst.preOrderTraverse(t.left)
		bst.preOrderTraverse(t.right)
	}
}

func (bst *BST) countNodes() int {
	return countHelper(bst.root)
}

func countHelper(root *BinaryNode) int {
	if root == nil {
		return 0
	}

	return countHelper(root.left) + countHelper(root.right) + 1
}

func (bst *BST) postOrderTraverse(t *BinaryNode) {
	if t != nil {
		bst.postOrderTraverse(t.left)
		bst.postOrderTraverse(t.right)
		fmt.Println(t.item)
	}
}

func (bst *BST) inOrder() {
	bst.inOrderTraverse(bst.root)
}

func (bst *BST) getItem(t *BinaryNode) string {
	for t.right != nil {
		t = t.right
	}
	return t.item
}

func (bst BST) removeNode(root **BinaryNode, item string) {
    if *root == nil {
        return
    }

    if item < (*root).item {
        bst.removeNode(&(*root).left, item)
    } else if item > (*root).item {
        bst.removeNode(&(*root).right, item)
    } else { // Node with the value found
        // Case 1: Node with only one child or no child
        if (*root).left == nil {
            *root = (*root).right
        } else if (*root).right == nil {
            *root = (*root).left
        } else {
            // Case 2: Node with two children
            minNode := findMin((*root).right)
            (*root).item = minNode.item
            bst.removeNode(&(*root).right, minNode.item)
        }
    }
}

// Helper function to find the minimum node in a subtree
func findMin(node *BinaryNode) *BinaryNode {
    for node.left != nil {
        node = node.left
    }
    return node
}

func main() {
	newTree := BST{}

	newTree.insert("aran")
	newTree.insert("lumi")
	newTree.insert("evan")
	newTree.insert("demon slayer")
	newTree.insert("merc")
	newTree.inOrder()

	ans,err := newTree.search(newTree.root, "evan")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Element found at", ans.item)
	}

	newTree.preOrderTraverse(newTree.root)
	newTree.postOrderTraverse(newTree.root)
	fmt.Println("Number of Nodes:", newTree.countNodes())

	newTree.removeNode(&newTree.root, "merc")
	newTree.inOrder()
	newTree.removeNode(&newTree.root, "evan")
	newTree.inOrder()
}
