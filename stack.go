package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type stack struct {
	top  *Node
	size int
}

func (p *stack) push(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.top == nil {
		p.top = newNode
	} else {
		newNode.next = p.top
		p.top = newNode
	}
	p.size++
	return nil
}

func (p *stack) pop() (string, error) {
	var item string

	if p.top == nil {
		return "", errors.New("Empty Stack!")
	}

	item = p.top.item
	if p.size == 1 {
		p.top = nil
	} else {
		p.top = p.top.next
	}
	p.size--
	return item, nil
}

func (p *stack) printAllNodes() error {
	currentNode := p.top
	if currentNode == nil {
		fmt.Println("Stack is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

// func main() {
// 	newStack := stack{}
// 	newStack.push("bob")
// 	newStack.push("tom")
// 	newStack.push("romeo")
// 	newStack.push("dio")
// 	newStack.printAllNodes()

// 	// activity 2, alternatively can travel down the node instead
// 	tempStack := stack{}

// 	for newStack.size > 0 {
// 		tempString, _ := newStack.pop()

// 		println("Manual Print:", tempString)
// 		tempStack.push(tempString)
// 	}

// 	for tempStack.size > 0 {
// 		x, _ := tempStack.pop()
// 		newStack.push(x)
// 	}

// 	newStack.printAllNodes()

// 	println("Check bool:", balancedParenthesis("(("))
// }

// leetcode innit, tested to work with all edge cases in leetcode
func balancedParenthesis(input string) bool {
	if len(input) % 2 != 0 {
		return false
	}

	newStack := stack{}

	for _, char := range input {
		if char == '{' {
			newStack.push("}")
		} else if char == '[' {
			newStack.push("]")
		} else if char == '(' {
			newStack.push(")")
		} else {
			temp, _ := newStack.pop()

			if temp != string(char) {
				return false
			}
		}
	}

	if newStack.size != 0 {
		return false
	}

	return true
}
