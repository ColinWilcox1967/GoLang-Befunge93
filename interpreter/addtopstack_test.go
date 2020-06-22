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

func TestAddTopStack (t *testing.T) {

	// test 1
	showNextTestHeader ()
	stack.Reset ()
	stack.Push (1)

	status = interpreter.AddTopStack (&stack)
	if status {
		fmt.Printf ("Stack size %d with value %d  (Expected 1, Value 1)\n", stack.Size (), stack.Peek ())
		t.Error ("Added top two items to stack when there was only one, should have failed")
	} 

	// test 2 - two values on stack replaced with a single sum
	showNextTestHeader ()
	stack.Push (2)
	status = interpreter.AddTopStack (&stack) 
	if !status {
		var val = stack.Peek ()
		fmt.Printf ("Stack size = %d Value = %d (Expected 1, Value 3)\n", stack.Size (), val)
		if val != 3 {
			t.Error ("Expecting a value of 3.")
		}
	} 

	// test 3
	showNextTestHeader ()
	stack.Reset ()
	stack.Push (1)
	stack.Push (2)
	stack.Push (4) // three items
 
	size := stack.Size ()
	if size != 3 {
		t.Errorf ("Expecting three items on the stack, only got %d\n", size)
	}

	status = interpreter.AddTopStack (&stack)
	if status {
		size = stack.Size ()
		if size != 2 {
			t.Errorf ("After first operation expecting two items on stack, left with %d\n", size)
		} else {
			if stack.Peek () == 6 {
				// do it again
				status = interpreter.AddTopStack (&stack)
				if status {
					size = stack.Size ()
					if size == 1 {
						if stack.Peek () != 7 {
							t.Errorf ("Expecting value of 7 on top of stack, received %d\n", stack.Peek ())
						} 
					} else {
						t.Errorf ("Expecting one item on the stack, got %d\n", size)
					}
				} else {
					t.Errorf ("After second operation expecting two items on stack, left with %d\n", size)
				}
			} else {
				t.Errorf ("Expected 6 on top of stack, peeked %d\n", stack.Peek ())
			}
		}
	}
}

// helpers
func showNextTestHeader () {
	testNumber++
	fmt.Println(fmt.Sprintf (">>> %s-%d", getCallerFunctionName (), testNumber))
}

func getCallerFunctionName () string {
  pc, _, _, ok := runtime.Caller(1)
    details := runtime.FuncForPC(pc)
    if ok && details != nil {
    	name := details.Name ()
    	var dotPos = strings.Index (name , ".")
   		return name[dotPos+1:]
  	}
	return ""
}


