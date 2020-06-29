package interpreter

import (
	"fmt"
	interpreter "github.com/colinwilcox1967/GoLang-Befunge93/interpreter"
	intstack "github.com/colinwilcox1967/intstack"
	"runtime"
	"strings"
	"testing"
)

var (
	status     bool
	stack      intstack.IntStack
	testNumber int
)

func TestDuplicateTop(t *testing.T) {

	// test 1 - empty stack
	showNextTestHeader()
	stack.Reset()

	status = interpreter.DuplicateTopStack(&stack)
	if status {
		t.Error("Stack was empty should have failed")
	}

	// test 2 - one value on stack duplicated
	showNextTestHeader()
	stack.Push(2)
	status = interpreter.DuplicateTopStack(&stack)
	if status {
		if stack.Size() == 2 {
			val1 := stack.Pop()
			val2 := stack.Pop()
			if val1 != val2 {
				fmt.Printf("Expecting top two items on stack to be same (Expected %d, Read %d&%d\n", 2, val1, val2)
				t.Error("Expecting two values of 2.")
			} else {
				if val1 != 2 || val2 != 2 {
					t.Errorf("Expecting value to be 2, read %d&%d", val1, val2)
				}
			}
		}
	}
}

// helpers
func showNextTestHeader() {
	var name string

	testNumber++
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		name = details.Name()
		var dotPos = strings.Index(name, ".")
		name = name[dotPos+1:]
	}

	fmt.Println(fmt.Sprintf(">>> %s-%d", name, testNumber))
}
