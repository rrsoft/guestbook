// stack test
package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := new(Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3.15)
	stack.Push(12.35)
	stack.Push("ranran")
	stack.Push("rrsoft")
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Count())
	stack.Clear()
	fmt.Println(stack.Count())
}
