package main

import (
"fmt"
"os"
"flag"
"strings"
"strconv"
"github.com/eiannone/keyboard"

instructionptr "github.com/colinwilcox1967/GoLang-Befunge93/instructionptr"
fileutilities "github.com/colinwilcox1967/golangfileutilities"
intstack "github.com/colinwilcox1967/intstack"
interpreter "github.com/colinwilcox1967/GoLang-Befunge93/interpreter"
)
	
const BEFUNGE93_VERSION = "v1.0"

// error status
const (
	KErrorNone 					 		= 0
	KErrorUnableToOpenFile 		 		= 1
	KErrorProblemParsingFile     		= 3
	KErrorSyntaxError			 		= 4
	KErrorInvalidCharacterFound			= 7
)

const (
	MAX_LINE_LENGTH = 80
	MAX_SOURCE_LINES = 25
	
	KErrorSourceLineTooLong = 100
	KErrorTooManySourceLines = 101
)

const (
	DEFAULT_SOURCE_FILE = "TEST.B93"
)

const (
	MOVE_UP int  = 0
	MOVE_DOWN int = 1
	MOVE_BACKWARDS int = 2
	MOVE_FORWARDS int = 3
)

var (
	sourceFileName string = DEFAULT_SOURCE_FILE
	stack 				intstack.IntStack
	instructionPtr		instructionptr.InstructionPtr
	endProgram bool
	currentDirection int
	inStringMode bool
)

func main () {
	displayBanner ()
	
	// read command line arguments, FILE'
	flag.StringVar (&sourceFileName, "file", DEFAULT_SOURCE_FILE,"Name of default source file")
	flag.Parse ()

	fmt.Printf ("Reading from source file '%s' ...\n", strings.ToUpper (sourceFileName))

	err, lineData := fileutilities.ReadFileAsLines (sourceFileName)
	if err != nil { // sigtnificant error just abort
		showStatus (KErrorUnableToOpenFile, sourceFileName, 0)
	}

	// now check the dimensions of the program
	status := checkProgramDimensions (lineData)

	if status == KErrorNone {
		// the fetch execute cycle
		initialise ()

		for !endProgram {
			var xPos int = instructionPtr.GetXPos ()
			var yPos int = instructionPtr.GetYPos ()
			var maxX int = len (lineData[yPos])
			var maxY int = len (lineData) - 1 // row zero indexed

			// check whether we are running string mode
			if inStringMode {
				// just push each character onto the stack for later
				
				if (lineData[yPos][xPos] == '"') {
					inStringMode = false
				} else {
					stack.Push(int(lineData[yPos][xPos]))
				}
			} else {
				switch (lineData[yPos][xPos]) {
					case '0','1','2','3','4','5','6','7','8','9': // push digits onto stack
								var str string = fmt.Sprintf ("%c", lineData[yPos][xPos])
								val,_ := strconv.Atoi(str)
								stack.Push (val)
					case '<': // move instruction pointer back, wrapping if necessary
							   	currentDirection = MOVE_BACKWARDS
				
					case '>': // move instruction point right, wrapping if necessary
								currentDirection = MOVE_FORWARDS
				
					case 'v': // move instruction pointer down, wrapping if necessary
								currentDirection = MOVE_DOWN
				
					case '^': // move instruction pointer up, wrapping if necessary
								currentDirection = MOVE_UP
				
					case '+': // add two two numbers then push
								interpreter.AddTopStack (&stack)
				
					case '-': // subtract top two numbers then push
								interpreter.SubtractTopStack (&stack)
				
					case '*': // multiply top two numbers then push
							    interpreter.MultiplyTopStack (&stack)
				
					case '/': // divide top two numbers then push
							     interpreter.DivideTopStack (&stack)
				
					case '%': // mod top two numbers then push
								 var val1 = stack.Pop ()
								 var val2 = stack.Pop ()
								 stack.Push (val2 % val1)

					case '!': // //logical not
								 var val = stack.Pop ()
								 if val == 0 {
								 	stack.Push (1)
								 } else {
								 	stack.Push(0)
								 }
					case '`': // greater than
								 var val1 = stack.Pop ()
								 var val2 = stack.Pop ()
								 if val2 > val1 {
								 	stack.Push (1)
								 } else {
								 	stack.Push(0)
								 }
					case '?': // move in a random cardinal direction
				
					case '_': // pop and move left or right
								 var val = stack.Pop ()
								 if val == 0 {
								 	instructionPtr.MoveInstructionPointerCardinal(MOVE_FORWARDS, maxX, maxY)
								 	currentDirection = MOVE_FORWARDS
								 } else {
								 	instructionPtr.MoveInstructionPointerCardinal(MOVE_BACKWARDS, maxX, maxY)
								 	currentDirection = MOVE_BACKWARDS
								 }
					case '|': // pop and move up or down
								var val = stack.Pop ()
								 if val == 0 {
								 	instructionPtr.MoveInstructionPointerCardinal(MOVE_DOWN, maxX, maxY)
								 	currentDirection = MOVE_DOWN
								 } else {
								 	instructionPtr.MoveInstructionPointerCardinal(MOVE_UP, maxX, maxY)
								 	currentDirection = MOVE_UP
								 }
					case '"': // enter string mode
								inStringMode = true
							
					case ':': // duplicate top of stack
								interpreter.DuplicateTopStack (&stack)

					case '\\': // swap top two on stack
								interpreter.ExchangeTopStack (&stack)
				
					case '$': // pop and discard
								 _ = stack.Pop ()
						
					case '.': // display top of stack as int
								interpreter.PopAndDisplayAsInt (&stack)
				
					case ',': // pop and display as character
								interpreter.PopAndDisplayAsASCII (&stack)
					case '#':
					case 'p': // put value
								var y = stack.Pop ()
								var x = stack.Pop ()
								var val = byte(stack.Pop ())
					 			lineData[y] = replaceCharAtStringIndex (lineData[y], val, x)

					case 'g': // get character
							  var y = stack.Pop ()
							  var x = stack.Pop ()

							  lineData[y] = replaceCharAtStringIndex (lineData[y], lineData[yPos][xPos], x)
					
					case '&': // input number
				//				char, _, err := keyboard.GetSingleKey () //? need multi digit?
				//				stack.Push (int(char))
					case '~': // ask for character and push
								char, _, _ := keyboard.GetSingleKey ()
								stack.Push (int(char))
					case '@': // end program
							endProgram = true
					default:
						// ignore do nothing
				}
			}

			// move to next cell in current direction
			switch currentDirection {
				case MOVE_FORWARDS:
									if xPos < len(lineData[yPos]) {
										instructionPtr.MoveInstructionPtrRelative (1,0) // one place to the right
									} else {
										instructionPtr.MoveInstructionPtr (0, yPos) // wrap round horizontally
									}
				case MOVE_BACKWARDS:
									if xPos > 0 {
										instructionPtr.MoveInstructionPtrRelative (-1, 0) // one place to the left
									} else {
										instructionPtr.MoveInstructionPtr (len(lineData[yPos])-1, yPos) // wrap round from the right
									}
				case MOVE_UP:
								   if yPos > 0 {
										instructionPtr.MoveInstructionPtrRelative (0, -1) // one place to the up
								   } else {
										instructionPtr.MoveInstructionPtr (xPos, 0) // wrap round from the top
								   }
				case MOVE_DOWN:
								   if yPos < len(lineData) {
								   		instructionPtr.MoveInstructionPtrRelative (0, 1) // one place down
								   } else {
								   		instructionPtr.MoveInstructionPtr (xPos,0) // wrap round from below
								   }
				default:
					fmt.Printf ("Unknown direction specified for instruction pointer (%d)\n", currentDirection)

			}

		}	
	}	
}

// private methods
func displayBanner () {
	fmt.Printf ("Befunge93 Interpreter %s\n\n", BEFUNGE93_VERSION)
}

func initialise () {
	stack.Reset ()
	instructionPtr.ResetInstructionPtr()
	endProgram = false
	currentDirection = MOVE_FORWARDS
	inStringMode = false
}

func checkProgramDimensions (data []string) int {
	var status int = KErrorNone

	// too many lines
	if len (data) > MAX_SOURCE_LINES {
		fmt.Printf ("Program contains %d source lines (Maximum is %d please reduce)\n", len(data), MAX_SOURCE_LINES)
		return KErrorTooManySourceLines
	}

	// theres a line thats too long?
	for lineNumber, line := range data {
		if len(line) > MAX_LINE_LENGTH {
			fmt.Printf ("Source line %d is too long (%d chars). Maximum length is %d\n", lineNumber+1, len(line), MAX_LINE_LENGTH)
			status = KErrorSourceLineTooLong
		}
	}
	
	return status
}

func parseFile (data []uint8) int {
	return KErrorNone
}

func executeFile () int {
	return KErrorNone
}


func replaceCharAtStringIndex(input string, replacement byte, index int) string {
    return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}

func showStatus (status int, extraInfoString string, extraInfoValue int) {

	var message string = "\nError: "

	if status != KErrorNone {
		switch (status) {
			case KErrorUnableToOpenFile: 
										message += fmt.Sprintf ("Unable to open file '%s'\n", strings.ToUpper(extraInfoString))
			case KErrorProblemParsingFile:
										message += fmt.Sprintf ("Problem parsing file '%s'\n", strings.ToUpper(extraInfoString))
			case KErrorSyntaxError:
										message += fmt.Sprintf ("Syntax error in file\n")
			case KErrorInvalidCharacterFound:
									    message += fmt.Sprintf ("Invalid character found whilst executing code : '%s'\n", extraInfoString)
			default:
				// shouldnt get here but catch it anyway
				fmt.Println ("Unknown Error Detected (%d)\n", status)
		}

		fmt.Println (message)
		os.Exit(status)
	}	
}

