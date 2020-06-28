package instructionptr

import "fmt"

type InstructionPtr struct {
	xPos, yPos int
}

const (
	MOVE_UP        int = 0
	MOVE_DOWN      int = 1
	MOVE_BACKWARDS int = 2
	MOVE_FORWARDS  int = 3
)

func (iptr *InstructionPtr) ResetInstructionPtr() {
	iptr.xPos = 0
	iptr.yPos = 0
}

func (iptr *InstructionPtr) MoveInstructionPtr(x, y int) {
	iptr.xPos = x
	iptr.yPos = y
}

func (iptr *InstructionPtr) MoveInstructionPtrRelative(deltaX, deltaY int) {
	iptr.xPos += deltaX
	iptr.yPos += deltaY
}

func (iptr *InstructionPtr) GetXPos() int {
	return iptr.xPos
}

func (iptr *InstructionPtr) GetYPos() int {
	return iptr.yPos
}

func (iptr *InstructionPtr) MoveInstructionPointerCardinal(direction, maxX, maxY int) {
	var xPos = iptr.GetXPos()
	var yPos = iptr.GetYPos()

	switch direction {
	case MOVE_UP:
		yPos--
		if yPos < 0 {
			yPos = maxY
		}
	case MOVE_DOWN:
		yPos++
		if yPos > maxY {
			yPos = 0
		}
	case MOVE_BACKWARDS:
		xPos--
		if xPos < 0 {
			xPos = maxX
		}
	case MOVE_FORWARDS:
		xPos++
		if xPos == maxX {
			xPos = 0
		}
	default:
		// Ignore and do nothing
	}
	iptr.MoveInstructionPtr(yPos, xPos)
}

// helper
func (iptr *InstructionPtr) ShowInstructionPtr() {
	fmt.Printf("Iptr: (%d,%d)\n", iptr.xPos, iptr.yPos)
}
