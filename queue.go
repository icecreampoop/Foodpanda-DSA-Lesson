package main

import (
	"errors"
	"fmt"
)

type qNode struct {
	item string
	next *qNode
	prio int
}

type queue struct {
	front *qNode
	back  *qNode
	size  int
}

func (p *queue) enqueue(name string, priority int) error {
	newNode := &qNode{
		item: name,
		prio: priority,
		next: nil,
	}
	if p.front == nil {
		p.front = newNode
		p.back = newNode

	} else {
		if newNode.prio < p.front.prio {
			newNode.next = p.front
			p.front = newNode
		} else {
			current := p.front
			for current.next != nil && current.next.prio <= newNode.prio {
                current = current.next
            }
			newNode.next = current.next
			current.next = newNode

			if newNode.next == nil {
				p.back = newNode
			}
		}

	}
	p.size++
	return nil
}

func (p *queue) dequeue() (string, error) {
	var item string

	if p.front == nil {
		return "", errors.New("empty queue!")
	}

	item = p.front.item
	if p.size == 1 {
		p.front = nil
		p.back = nil
	} else {
		p.front = p.front.next
	}
	p.size--
	return item, nil
}

func (p *queue) printAllNodes() error {
	currentNode := p.front
	if currentNode == nil {
		fmt.Println("Queue is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func (p *queue) isEmpty() bool {
	return p.size == 0
}

func main() {
	newQueue := queue{}

	//newQueue.enqueue("iron")
	//newQueue.enqueue("bronze")
	//newQueue.enqueue("silver")
	//newQueue.enqueue("gold")

	newQueue.printAllNodes()

	// using 2 pointer is better but
	// palinQueue := queue{}
	// palinStack := stack{}
	// var temp1, temp2 string
	// padiBool := true

	// for _, char := range "referr" {
	// 	palinQueue.enqueue(string(char))
	// 	palinStack.push(string(char))
	// }

	// for x := 0; x < palinQueue.size; x++ {
	// 	temp1, _ = palinQueue.dequeue()
	// 	temp2, _ = palinStack.pop()

	// 	if temp1 != temp2{
	// 		fmt.Println("This is not a Padini")
	// 		padiBool = false
	// 		break
	// 	}
	// }

	// if padiBool {
	// 	println("This is Padini")
	// }

	pq := queue{}
    pq.enqueue("low-priority", 5)
    pq.enqueue("medium-priority", 3)
    pq.enqueue("high-priority", 1)
    pq.enqueue("another-medium", 3)
    pq.enqueue("another-low", 6)

    fmt.Println("Queue after enqueuing elements:")
    pq.printAllNodes()
}
