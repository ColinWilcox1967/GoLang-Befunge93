package befunge93interpreter

import (
	"fmt"
	"testing"
	"runtime"
	"strings"
	intstack "github.com/colinwilcox1967/intstack"
	interpreter "github.com/colinwilcox1967/GoLang-Befunge93/interpreter"
)

var (
	status bool
	stack intstack.IntStack
	testNumber int
)

func TestDivideTopStack (t *testing.T) {

	// test 1
	showNextTestHeader ()
	stack.Reset ()
	stack.Push (4)

	status = interpreter.DivideTopStack (&stack)
	if status {
		fmt.Printf ("Stack size %d with value %d  (Expected 1, Value 1)\n", stack.Size (), stack.Peek ())
		t.Error ("Dividing top two items to stack when there was only one, should have failed")
	} 

	// test 2 - two values on stack replaced with a single product
	showNextTestHeader ()
	stack.Push (2)
	status = interpreter.DivideTopStack (&stack) 
	if status {
		var val = stack.Peek ()
		if val != 2 {
			fmt.Printf ("Stack size = %d Value = %d (Expected 1, Value 2)\n", stack.Size (), val)
			t.Error ("Expecting a value of 2.")
		}
	} else {
		t.Error ("Division failed")
	}

	// test 3 - two values on stack replaced with a single product
	showNextTestHeader ()
	stack.Push (0)
	stack.Push (0)
	status = interpreter.DivideTopStack (&stack) 
	if status {
		var val = stack.Peek ()
		if val != 0 {
			fmt.Printf ("Stack size = %d Value = %d (Expected 1, Value 0)\n", stack.Size (), val)
			t.Error ("Expecting a value of 0.")
		}
	} else {
		t.Error ("Division failed")
	}

	// test 4
	showNextTestHeader ()
	stack.Reset ()
	stack.Push (1)
	stack.Push (4)
	stack.Push (2) // three items
 
	size := stack.Size ()
	if size != 3 {
		t.Errorf ("Expecting three items on the stack, only got %d\n", size)
	}

	status = interpreter.DivideTopStack (&stack)

	if status {
		size = stack.Size ()
		if size != 2 {
			t.Errorf ("After first operation expecting two items on stack, left with %d\n", size)
		} else {
			if stack.Peek () == 2 {
				// do it again
				status = interpreter.DivideTopStack (&stack)
				if status {
					size = stack.Size ()
					if size == 1 {
						if stack.Peek () != 0{
							t.Errorf ("Expecting value of 0 on top of stack, received %d\n", stack.Peek ())
						} 
					} else {
						t.Errorf ("Expecting one item on the stack, got %d\n", size)
					}
				} else {
					t.Errorf ("After second operation expecting two items on stack, left with %d\n", size)
				}
			} else {
				t.Errorf ("Expected 2 on top of stack, peeked %d\n", stack.Peek ())
			}
		}
	}
}

// helpers
func showNextTestHeader () {
	var name string

	testNumber++
	pc, _, _, ok := runtime.Caller(1)
    details := runtime.FuncForPC(pc)
    if ok && details != nil {
    	name = details.Name ()
    	var dotPos = strings.Index (name , ".")
   		name = name[dotPos+1:]
  	}

	fmt.Println(fmt.Sprintf (">>> %s-%d", name, testNumber))
}

