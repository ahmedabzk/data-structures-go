package main

import(
	"fmt"
)

type Node struct{
	data int
	next *Node

}

type LinkedList struct{
	length int 
	head *Node
	tail *Node
}

func (i LinkedList) len() int{
	return i.length
}

func (i LinkedList) pushBack(n *Node){
	if i.head == nil{
		i.head = n
		i.tail = n
		i.length++
	}else{
		i.tail.next = n
		i.tail = n
		i.length++
	}
}

func main(){
	node := &Node{data: 20}
	list := LinkedList{}

	list.pushBack(node)
	fmt.Println(node)
	fmt.Println(list)
	


}