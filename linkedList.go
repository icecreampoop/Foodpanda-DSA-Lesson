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

func (p *linkedList) get(position int) string {
	var temp *Node = p.head
	
	for position != 0 {
		temp = temp.next
		position--
	}

	return temp.item
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

func (p *linkedList) addAtPos(position int, item string) error {
	var tempHead = p.head
	var prev *Node

	// if pos given > whats inside just add to end
	if position >= p.size {
		for tempHead.next != nil {
			tempHead = tempHead.next
		}

		// once reached end, to make sure the new node is added at the end instead
		prev = tempHead
		tempHead = nil
	} else {
		for position != 0 {
			prev = tempHead
			tempHead = tempHead.next
			position--
		}
	}

	// meaning if adding at pos 0, the first element
	if prev == nil {
		p.head = &Node{
			item: item,
			next: p.head,
		}
	} else {
		prev.next = &Node{
			item: item,
			next: tempHead,
		}
	}

	p.size++
	return nil
}

func (p *linkedList) remove(position int) error {
	var current *Node = p.head
	var prev *Node

	if position > p.size-1 || 0 > position {
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
	newLinkedList := linkedList{}
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

	println("adding at index 0")
	newLinkedList.addAtPos(0, "valo1")
	newLinkedList.printAllNodes()
	println("adding at index 2")
	newLinkedList.addAtPos(2, "valo3")
	newLinkedList.printAllNodes()
	println("adding to end")
	newLinkedList.addAtPos(1000, "valo5")
	newLinkedList.printAllNodes()

	for x := 0; x < newLinkedList.size; x++ {
		println("Item at position", x, "is:", newLinkedList.get(x))
	}
}
