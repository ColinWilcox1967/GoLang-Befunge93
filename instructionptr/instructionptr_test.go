package befunge93instructionptr

import (
	"testing"
	"fmt"
	"strings"
	"runtime"
	instructionptr "github.com/colinwilcox1967/GoLang-Befunge93/instructionptr"
)


var (
	ptr instructionptr.InstructionPtr
	testNumber int = 0
)


func TestResetInstructionPtr (t *testing.T) {
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	ptr.ShowInstructionPtr ()
}

func TestSetPositionX (t *testing.T) {
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	for x := 1; x<10; x++ {
		ptr.MoveInstructionPtr (x, ptr.GetYPos ())
		ptr.ShowInstructionPtr ()
	}

	x := ptr.GetXPos ()
	y := ptr.GetYPos ()
	fmt.Printf ("Final Pos : x=%d, y=%d\n", x,y)

	if x != 9 || y != 0 {
		t.Errorf ("Expected pointer to be at (9,0)\n")
	}
}

func TestSetPositionY (t *testing.T) {
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	for y := 1 ; y<10; y++ {
		ptr.MoveInstructionPtr (ptr.GetXPos (), y)
		ptr.ShowInstructionPtr ()

	}

	x := ptr.GetXPos ()
	y := ptr.GetYPos ()
	fmt.Printf ("Final Pos : x=%d, y=%d\n", x,y)

	if x != 0 || y != 9 {
		t.Errorf ("Expected final pointer to be at (0,9)\n")
	}
}

func TestMovePositionRelative (t *testing.T) {
	
	// test 1 - no move
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos := ptr.GetXPos ()
	originalYPos := ptr.GetYPos ()

	ptr.MoveInstructionPtrRelative (0,0)

	newXPos := ptr.GetXPos ()
	newYPos := ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)

	if (originalXPos != newXPos) || (originalYPos != newYPos) {
		t.Errorf ("Expected pointer to remain in same place")
	}

	// test 2 - move x positive only
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos = ptr.GetXPos ()
	originalYPos = ptr.GetYPos ()

	ptr.MoveInstructionPtrRelative (1,0)

	newXPos = ptr.GetXPos ()
	newYPos = ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)

	if newXPos != 1 || newYPos != 0 {
		t.Error ("Expected final position to be (1,0)")
	}

	// test 3 - move x negative  only
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos = ptr.GetXPos ()
	originalYPos = ptr.GetYPos ()


	ptr.MoveInstructionPtrRelative (-1,0)

	newXPos = ptr.GetXPos ()
	newYPos = ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)
	if newXPos != -1 || newYPos != 0 {
		t.Error ("Expected final position to be (-1,0)")
	}

	// test 4 - move Y positive only
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos = ptr.GetXPos ()
	originalYPos = ptr.GetYPos ()

	ptr.MoveInstructionPtrRelative (0,1)

	newXPos = ptr.GetXPos ()
	newYPos = ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)
	if newXPos != 0 || newYPos != 1 {
		t.Error ("Expected final position  to be (0,1)")
	}

	// test 5 - move Y negative only
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos = ptr.GetXPos ()
	originalYPos = ptr.GetYPos ()


	ptr.MoveInstructionPtrRelative (0,-1)

	newXPos = ptr.GetXPos ()
	newYPos = ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)
	if newXPos != 0 || newYPos != -1 {
		t.Error ("Expected final position to be (0,-1)")
	}

	// test 6 - move X,Y positive
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos = ptr.GetXPos ()
	originalYPos = ptr.GetYPos ()

	ptr.MoveInstructionPtrRelative (1,1)

	newXPos = ptr.GetXPos ()
	newYPos = ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)

	if newXPos != 1 || newYPos != 1 {
		t.Error ("Expected final position to be (1,1)")
	}

	// test 7 - move x,y negative
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos = ptr.GetXPos ()
	originalYPos = ptr.GetYPos ()

	ptr.MoveInstructionPtrRelative (-1,-1)

	newXPos = ptr.GetXPos ()
	newYPos = ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)

	if newXPos != -1 || newYPos != -1 {
		t.Error ("Expected final position to be (-1,-1)")
	}

	// test 8 - move X positive, Y negative
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos = ptr.GetXPos ()
	originalYPos = ptr.GetYPos ()

	ptr.MoveInstructionPtrRelative (1,-1)

	newXPos = ptr.GetXPos ()
	newYPos = ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)

	if newXPos != 1 || newYPos != -1 {
		t.Error ("Expected new position to be (1,-1)")
	}

	// test 9 - move x negative, y positive
	showNextTestHeader ()
	ptr.ResetInstructionPtr ()
	originalXPos = ptr.GetXPos ()
	originalYPos = ptr.GetYPos ()

	ptr.MoveInstructionPtrRelative (-1,1)

	newXPos = ptr.GetXPos ()
	newYPos = ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)
	if newXPos != -1 || newYPos != 1 {
		t.Error ("Expected final position to be (-1,1)")
	}
}

// helpers
func showNextTestHeader () {
	testNumber++
	fmt.Println(fmt.Sprintf ("%s-%d", getCurrentFunctionName (), testNumber))
}

func getCurrentFunctionName () string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	
	//strip prefix upto leading '.'
	dotPos := strings.Index (frame.Function,".")
	
    return frame.Function [dotPos+1:]
}



