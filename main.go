package main

import (
	"flag"
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	instructionptr "github.com/colinwilcox1967/GoLang-Befunge93/instructionptr"
	interpreter "github.com/colinwilcox1967/GoLang-Befunge93/interpreter"
	fileutilities "github.com/colinwilcox1967/golangfileutilities"
	intstack "github.com/colinwilcox1967/intstack"
)

const BEFUNGE93_VERSION = "v1.0"

// error status
const (
	KErrorNone                  = 0
	KErrorUnableToOpenFile      = 1
	KErrorProblemParsingFile    = 3
	KErrorSyntaxError           = 4
	KErrorInvalidCharacterFound = 7
)

const (
	MAX_LINE_LENGTH  = 80
	MAX_SOURCE_LINES = 25

	KErrorSourceLineTooLong  = 100
	KErrorTooManySourceLines = 101
)

const (
	MOVE_UP    int = 0
	MOVE_DOWN  int = 1
	MOVE_LEFT  int = 2
	MOVE_RIGHT int = 3
)

const (
	DEFAULT_SOURCE_FILE = "TEST.B93"
)

var (
	sourceFileName   string = DEFAULT_SOURCE_FILE
	stack            intstack.IntStack
	instructionPtr   instructionptr.InstructionPtr
	endProgram       bool
	currentDirection int
	inStringMode     bool
	skipCell         bool // support '#'
)

func main() {
	displayBanner()

	// read command line arguments, FILE'
	flag.StringVar(&sourceFileName, "file", DEFAULT_SOURCE_FILE, "Name of default source file")
	flag.Parse()

	fmt.Printf("Reading from source file '%s' ...\n", strings.ToUpper(sourceFileName))

	err, lineData := fileutilities.ReadFileAsLines(sourceFileName)
	if err != nil { // sigtnificant error just abort
		showStatus(KErrorUnableToOpenFile, sourceFileName, 0)
	}

	// now check the dimensions of the program
	status := checkProgramDimensions(lineData)

	if status == KErrorNone {
		// the fetch execute cycle
		initialise()

		for !endProgram {
			var xPos int = instructionPtr.GetXPos()
			var yPos int = instructionPtr.GetYPos()

			// make sure we havent been sent offgrid by incorrect code
			if offGrid (lineData, xPos, yPos) {
				fmt.Printf ("Program terminated - sent off grid at cell (Row = %d, Column = %d)\n", yPos, xPos)
				stack.DisplayStack ()
				endProgram = true
				continue
			}

			// check whether we are running string mode
			if inStringMode {
				// just push each character onto the stack for later

				if lineData[yPos][xPos] == '"' {
					inStringMode = false
				} else {
					stack.Push(int(lineData[yPos][xPos]))
				}
			} else {
				switch lineData[yPos][xPos] {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': // push digits onto stack
					var str string = fmt.Sprintf("%c", lineData[yPos][xPos])
					val, _ := strconv.Atoi(str)
					stack.Push(val)
				case '<': // move instruction pointer back, wrapping if necessary
					currentDirection = MOVE_LEFT

				case '>': // move instruction point right, wrapping if necessary
					currentDirection = MOVE_RIGHT

				case 'v': // move instruction pointer down, wrapping if necessary
					currentDirection = MOVE_DOWN

				case '^': // move instruction pointer up, wrapping if necessary
					currentDirection = MOVE_UP

				case '+': // add two two numbers then push
					interpreter.AddTopStack(&stack)

				case '-': // subtract top two numbers then push
					interpreter.SubtractTopStack(&stack)

				case '*': // multiply top two numbers then push
					interpreter.MultiplyTopStack(&stack)

				case '/': // divide top two numbers then push
					interpreter.DivideTopStack(&stack)

				case '%': // mod top two numbers then push
					interpreter.ModTopStack(&stack)

				case '!': // //logical not
					interpreter.LogicalNot(&stack)

				case '`': // greater than
					var val1 = stack.Pop()
					var val2 = stack.Pop()
					if val2 > val1 {
						stack.Push(1)
					} else {
						stack.Push(0)
					}
				case '?': // move in a random cardinal direction
					// get a seed
					seed := rand.NewSource(time.Now().UnixNano())
					r := rand.New(seed)
					currentDirection = r.Intn(4) // choose a random cardinal direction 0 .. 3 -> MOVE_UP to MOVE_RIGHT
				case '_': // pop and move left or right
					var val = stack.Pop()
					if val == 0 {
						MoveToNextCellInCurrentDirection(MOVE_RIGHT, xPos, yPos, len(lineData[yPos]), len(lineData))
						currentDirection = MOVE_RIGHT
					} else {
						MoveToNextCellInCurrentDirection(MOVE_LEFT, xPos, yPos, len(lineData[yPos]), len(lineData))
						currentDirection = MOVE_LEFT
					}
				case '|': // pop and move up or down
					var val = stack.Pop()
					if val == 0 {
						MoveToNextCellInCurrentDirection(MOVE_DOWN, xPos, yPos, len(lineData[yPos]), len(lineData))
						currentDirection = MOVE_DOWN
					} else {
						MoveToNextCellInCurrentDirection(MOVE_UP, xPos, yPos, len(lineData[yPos]), len(lineData))
						currentDirection = MOVE_UP
					}
				case '"': // enter string mode
					inStringMode = true

				case ':': // duplicate top of stack
					interpreter.DuplicateTopStack(&stack)

				case '\\': // swap top two on stack
					interpreter.ExchangeTopStack(&stack)

				case '$': // pop and discard
					_ = stack.Pop()

				case '.': // display top of stack as int
					interpreter.PopAndDisplayAsInt(&stack)

				case ',': // pop and display as character
					interpreter.PopAndDisplayAsASCII(&stack)
				case '#': // skip next cell
					skipCell = true
				case 'p': // put value
					var y = stack.Pop()
					var x = stack.Pop()
					var val = byte(stack.Pop())
					lineData[y] = replaceCharAtStringIndex(lineData[y], val, x)

				case 'g': // get character
					var y = stack.Pop()
					var x = stack.Pop()

					lineData[y] = replaceCharAtStringIndex(lineData[y], lineData[yPos][xPos], x)

				case '&': // input number
					char, _, _ := keyboard.GetSingleKey() //? need multi digit?
					stack.Push(int(char))
				case '~': // ask for character and push
					char, _, _ := keyboard.GetSingleKey()
					stack.Push(int(char))
				case '@': // end program
					endProgram = true
				default:
					// ignore do nothing
				}
			}

			// if we skip the cell then move ahead in current direction before doing anything else
			if skipCell {
				MoveToNextCellInCurrentDirection(currentDirection, xPos, yPos, len(lineData[yPos]), len(lineData))
				skipCell = false
			}

			MoveToNextCellInCurrentDirection(currentDirection, xPos, yPos, len(lineData[yPos]), len(lineData))

		}
	}
}

// private methods
func displayBanner() {
	fmt.Printf("Befunge93 Interpreter %s", BEFUNGE93_VERSION)

	_, filename, _, _ := runtime.Caller(1)
	fileStat, err := os.Stat(filename)
	if err == nil {
		modifiedString := fmt.Sprintf("%s", fileStat.ModTime())
		pos := strings.Index(modifiedString, " ")
		modifiedString = modifiedString[:pos]
		fmt.Printf(" (%s)", modifiedString)
	}

	fmt.Printf("\n\n")
}

func initialise() {
	stack.Reset()
	instructionPtr.ResetInstructionPtr()
	endProgram = false
	currentDirection = MOVE_RIGHT
	inStringMode = false
	skipCell = false
}

func checkProgramDimensions(data []string) int {
	var status int = KErrorNone

	// too many lines
	if len(data) > MAX_SOURCE_LINES {
		fmt.Printf("Program contains %d source lines (Maximum is %d please reduce)\n", len(data), MAX_SOURCE_LINES)
		return KErrorTooManySourceLines
	}

	// theres a line thats too long?
	for lineNumber, line := range data {
		if len(line) > MAX_LINE_LENGTH {
			fmt.Printf("Source line %d is too long (%d chars). Maximum length is %d\n", lineNumber+1, len(line), MAX_LINE_LENGTH)
			status = KErrorSourceLineTooLong
		}
	}

	return status
}

func replaceCharAtStringIndex(input string, replacement byte, index int) string {
	return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}

func MoveToNextCellInCurrentDirection(currentDirection, xPos, yPos, xMax, yMax int) {
	// move to next cell in current direction
	switch currentDirection {
	case MOVE_RIGHT:
		if xPos < xMax {
			instructionPtr.MoveInstructionPtrRelative(1, 0) // one place to the right
		} else {
			instructionPtr.MoveInstructionPtr(0, yPos) // wrap round horizontally
		}
	case MOVE_LEFT:
		if xPos > 0 {
			instructionPtr.MoveInstructionPtrRelative(-1, 0) // one place to the left
		} else {
			instructionPtr.MoveInstructionPtr(xMax-1, yPos) // wrap round from the right
		}
	case MOVE_UP:
		if yPos > 0 {
			instructionPtr.MoveInstructionPtrRelative(0, -1) // one place to the up
		} else {
			instructionPtr.MoveInstructionPtr(xPos, 0) // wrap round from the top
		}
	case MOVE_DOWN:
		if yPos < yMax {
			instructionPtr.MoveInstructionPtrRelative(0, 1) // one place down
		} else {
			instructionPtr.MoveInstructionPtr(xPos, 0) // wrap round from below
		}
	default:
		fmt.Printf("Unknown direction specified for instruction pointer (%d)\n", currentDirection)

	}
}

func showStatus(status int, extraInfoString string, extraInfoValue int) {

	var message string = "\nError: "

	if status != KErrorNone {
		switch status {
		case KErrorUnableToOpenFile:
			message += fmt.Sprintf("Unable to open file '%s'\n", strings.ToUpper(extraInfoString))
		case KErrorProblemParsingFile:
			message += fmt.Sprintf("Problem parsing file '%s'\n", strings.ToUpper(extraInfoString))
		case KErrorSyntaxError:
			message += fmt.Sprintf("Syntax error in file\n")
		case KErrorInvalidCharacterFound:
			message += fmt.Sprintf("Invalid character found whilst executing code : '%s'\n", extraInfoString)
		default:
			// shouldnt get here but catch it anyway
			fmt.Printf("Unknown Error Detected (%d)\n", status)
		}

		fmt.Println(message)
		os.Exit(status)
	}
}

func offGrid (lineData []string, x,y int) bool {
	// row bad
	if y < 0 || y == len(lineData) {
		return true
	}

	// column is bad ?
	if x <0 || x == len(lineData[y]) {
		return true
	}

	

	return false
}




