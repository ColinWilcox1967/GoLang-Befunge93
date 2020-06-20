package befunge93instructionptr

type InstructionPtr struct {
	xPos, yPos int
}

func (iptr *InstructionPtr)ResetInstructionPtr () {
	iptr.xPos = 0
	iptr.yPos = 0
}

func (iptr *InstructionPtr)SetInstructionPtr (x,y int) {
	iptr.xPos = x
	iptr.yPos = y
}

func (iptr *InstructionPtr)GetXPos () int {
	return iptr.xPos
}

func (iptr *InstructionPtr)GetYPos () int {
	return iptr.yPos
}
