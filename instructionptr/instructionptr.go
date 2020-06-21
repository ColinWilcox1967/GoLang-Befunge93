package befunge93instructionptr

import "fmt"

type InstructionPtr struct {
	xPos, yPos int
}

func (iptr *InstructionPtr)ResetInstructionPtr () {
	iptr.xPos = 0
	iptr.yPos = 0
}

func (iptr *InstructionPtr)MoveInstructionPtr (x,y int) {
	iptr.xPos = x
	iptr.yPos = y
}

func (iptr *InstructionPtr)MoveInstructionPtrRelative (deltaX, deltaY int) {
	iptr.xPos += deltaX
	iptr.yPos += deltaY
}

func (iptr *InstructionPtr)GetXPos () int {
	return iptr.xPos
}

func (iptr *InstructionPtr)GetYPos () int {
	return iptr.yPos
}

// helper
func (iptr *InstructionPtr)ShowInstructionPtr () {
	fmt.Printf ("Iptr: (%d,%d)\n", iptr.xPos, iptr.yPos)
}
