# Part of the foodpanda lessons related to DSA stuff

Publishing this to github for visibility sake since I am just starting my career, gotta look decent at least. I have other notes relating to DSA beyond what is shown here and practise ✨Leetcode✨ sometimes.

## Table of Contents
- Linked-Lists
- Stacks
- Queues
- Trees

## Linked-Lists

### Activity #1: Create a simple main program that uses a linked list structure

1. Create an instance of the linked list.
2. Add 4 items to it by calling the `addNode()` function.
3. Now call the `printAllNodes` function to verify that the items were added correctly to the linked list.

### Activity #2: Adding more functions to the linked list structure

1. `remove()` that removes an item at a specified position.
2. `addAtPos()` that adds an item at a specified position.
3. `get()` that retrieves the item at a specified position.
4. Now write code in your main program to test the above 3 functions.

## Stacks

### Activity #1: Create a simple main program that uses a stack structure

1. Create an instance of the stack.
2. Push 4 items it by calling the push() function.
3. Now call the printAllNodes() function to verify the nodes were pushed correctly into the stack. (The items should be printed in the reverse manner that they go into the stack.)

### Activity #2: Printing contents of the stack in main program

Sometimes, there is no function in the stack structure that allows printing of all its nodes. Now you would write code in the main program that will allow you to print the contents in a stack. (hint you may need another instance of certain data structure)

### Activity #3: Checking for unbalanced parenthesis in a program

Stacks can be used to check for unbalanced parenthesis in a program. Write a program that checks a user input of a program and displays if the program contains unbalanced parenthesis.

## Queues

### Activity #1: Create a simple main program that uses a queue structure

1. Create an instance of the queue.
2. Add 4 items to it by calling the enqueue() function.
3. Now call the printAllNodes function to verify that the items were added correctly to the queue.

### Activity #2: Determining Palindromes

Let’s write some code in your driver program (i.e. your main) to determine if an input string is a palindrome. A palindrome is any sequence of characters that reads the same forwards and backwards.
E.g. words like “refer”, “noon”, “level”, “my gym”, “top spot”

(hint: you need to make use of a queue and stack structure)

### Activity #3: Implementing a Priority Queue
Now improve on queue.go to become a priority queue. As mentioned, elements with higher priority will be processed before elements of lower priority, therefore they will appear in the queue nearer to
front compared to those with lower priority. If they are of same priority level, they will be processed in order of their enqueue.

Let higher priority be represented by a smaller number in our case. Larger priority number indicate less importance.

(hint: need to modify enqueue function, as well as node struct of queue to include priority level)

## Trees

### Activity #1: Create a simple main program that uses a tree structure

1. Create an instance of the tree.
2. Insert a few items it by calling the insert() function.
3. Now call the inOrder() traversal function to verify the nodes were inserted correctly into the tree.

### Activity #2: Search function for the tree structure

Now implement a search function to search for an item in a binary search tree.

### Activity #3: PreOrder and PostOrder traversals

You have seen how InOrder traversal can be implemented, as provided in the code base. Now implement the other two traversals: PreOrder and PostOrder.

### Activity #4: Count no of nodes function for the tree

Implement a function that counts total number of nodes in the tree.

### Activity #5: Remove function for the tree structure

Now implement a remove function to remove an item from the binary search tree.
