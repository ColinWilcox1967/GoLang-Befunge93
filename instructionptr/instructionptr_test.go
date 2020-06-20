package befunge93instructionptr

import (
	"testing"
	"fmt"
	instructionptr "github.com/colinwilcox1967/GoLang-Befunge93/instructionptr"
)

var ptr instructionptr.InstructionPtr


func TestResetInstructionPtr (t * testing.T) {
	ptr.ResetInstructionPtr ()
	ptr.ShowInstructionPtr ()
}

func TestSetPositionX (t * testing.T) {
	
	ptr.ResetInstructionPtr ()
	for x := 1 ; x<10; x++ {
		ptr.SetInstructionPtr (x, ptr.GetYPos ())
		ptr.ShowInstructionPtr ()
	}

	x := ptr.GetXPos ()
	y := ptr.GetYPos ()
	fmt.Printf ("Final Pos : x=%d, y=%d\n", x,y)
}

func TestSetPositionY (t * testing.T) {
	
	ptr.ResetInstructionPtr ()
	for y := 1 ; y<10; y++ {
		ptr.SetInstructionPtr (ptr.GetXPos (), y)
		ptr.ShowInstructionPtr ()
	}

	x := ptr.GetXPos ()
	y := ptr.GetYPos ()
	fmt.Printf ("Final Pos : x=%d, y=%d\n", x,y)
}
// helper functions


