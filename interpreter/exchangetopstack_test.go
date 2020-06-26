package interpreter

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

func TestExchangeTop (t *testing.T) {

	// test 1 - empty stack
	showNextTestHeader ()
	stack.Reset ()

	status = interpreter.ExchangeTopStack (&stack)
	if status {
		t.Error ("Stack was empty should have failed")
	}

	// test 2 - one value on stack duplicated
	showNextTestHeader ()
	stack.Push (2)
	stack.Push (3)
	status = interpreter.ExchangeTopStack (&stack) 
	if status {
		if stack.Size () == 2 {
			val1 := stack.Pop ()
			val2 := stack.Pop ()
			if val1 != 2 && val2 != 3 {
				fmt.Printf ("Expecting top two items on stack to be swapped (Expected %d&%d, Read %d&%d\n", 2, 3, val1, val2)
				t.Error ("Expecting two values to be swapped")
			} 
		}
	}

	// test 3 - three items on stack, top two exchanged, third is untouched
	showNextTestHeader ()
	
	stack.Push (5)
	stack.Push (2)
	stack.Push (3)
	
	status = interpreter.ExchangeTopStack (&stack) 
	if status {
		if stack.Size () == 3 { // nothing lost
			val1 := stack.Pop ()
			val2 := stack.Pop ()
			val3 := stack.Pop ()
			if val1 != 2 && val2 != 3 {
				fmt.Printf ("Expecting top two items on stack to be swapped (Expected %d&%d, Read %d&%d\n", 2, 3, val1, val2)
				t.Error ("Expecting two values to be swapped")
			} else {
				if val3 != 5 {
				fmt.Printf ("Expecting third stack value to be %d (Read %d)\n", 5, val3)
				t.Error ("Expecting third value on stack to be unchanged")
					
				}
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

