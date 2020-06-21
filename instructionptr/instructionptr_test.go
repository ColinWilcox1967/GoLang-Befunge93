package befunge93instructionptr

import (
	"testing"
	"fmt"
	instructionptr "github.com/colinwilcox1967/GoLang-Befunge93/instructionptr"
)

var ptr instructionptr.InstructionPtr


func TestResetInstructionPtr (t *testing.T) {
	ptr.ResetInstructionPtr ()
	ptr.ShowInstructionPtr ()
}

func TestSetPositionX (t *testing.T) {
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
	ptr.ResetInstructionPtr ()
	originalXPos := ptr.GetXPos ()
	originalYPos := ptr.GetYPos ()

	// test 1 - no move
	ptr.MoveInstructionPtrRelative (0,0)

	newXPos := ptr.GetXPos ()
	newYPos := ptr.GetYPos ()

	fmt.Printf ("Original Pos (%d,%d), New Pos (%d, %d)\n", originalXPos, originalYPos, newXPos, newYPos)

	if (originalXPos != newXPos) || (originalYPos != newYPos) {
		t.Errorf ("Expected pointer to remain in same place")
	}

	// test 2 - move x positive only
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



