package main 
import(
	"fmt"
)

var stackSize int

type Stack struct{
	items []interface{}
	top int
}

// push an element to the stack
func (s *Stack) push(value interface{}){
	// check if the stack is full
	if s.top == stackSize -1{
		fmt.Println("stack overflow..stack is full")
		return
	}else{
		s.items = append(s.items, value)
		s.top++
		fmt.Printf("inserted value %v\n",value)
	}
	stackSize++
	return
}

// pop an element

func (s *Stack) pop() interface{}{
	// check if the stack is empty
	if s.top == -1{
		fmt.Println("stack is empty, no value to pop")
		return nil
	}else{
		fmt.Printf("poped item %v\n",s.items[s.top-1])
		s.items = s.items[:s.top-1]
		s.top--
	}
	stackSize--
	return nil
}

// display stack items
func (s *Stack) display(){
	// check if stack is empty
	if s.top == -1{
		fmt.Println("stack is empty")
	}else{
		for _, val := range s.items{
			fmt.Printf("items in the stack %v\n",val)
		}
	}
	
}

// return at the top element without removing
func (s *Stack) peek() interface{}{
	// check if stack is empty
	if s.top == -1{
		fmt.Println("stack is empty")
		return nil
	}
	fmt.Printf("peeked item %v\n",s.items[s.top-1])
	return nil
}

func main(){
	stack := Stack{items: make([]interface{}, 0)}
	stack.push(5)
	stack.push(6)
	stack.push(8)
	stack.push(9)
	fmt.Println(stack.items)
	stack.peek()
	stack.pop()
	stack.pop()
	fmt.Println(stack.items)

}