package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func (p *linkedList) addNode(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.head == nil {
		p.head = newNode
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	p.size += 1
	return nil
}

func (p *linkedList) printAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func (p *linkedList) addAtPos() error {
	
	return nil
}

func (p *linkedList) remove(position int) error{
	var current *Node = p.head
	var prev *Node

	if position > p.size - 1 || 0 > position{
		return errors.New("Position out of index bounds")
	}

	for position != 0 {
		prev = current
		current = current.next
		position--
	}

	//if pos was 0 at start will have error so
	if prev == nil {
		p.head = p.head.next
		p.size--
		return nil
	}

	prev.next = current.next
	p.size--

	return nil
}

func main() {
	newLinkedList := linkedList {}
	newLinkedList.addNode("lol1")
	newLinkedList.addNode("lol2")
	newLinkedList.addNode("lol3")
	newLinkedList.addNode("lol4")

	newLinkedList.printAllNodes()

	println("removing index 0")
	newLinkedList.remove(0)
	newLinkedList.printAllNodes()
	println("removing index 1")
	newLinkedList.remove(1)
	newLinkedList.printAllNodes()
}
