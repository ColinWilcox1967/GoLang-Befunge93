package main

import (
"fmt"
"os"
"flag"
"io/ioutil"
"strings"

"github.com/eiannone/keyboard"
"github.com/colinwilcox1967/intstack"
interpreter "github.com/colinwilcox1967/golang-befunge93/interpreter"
instructionptr "github.com/colinwilcox1967/golang-befunge93/instructionptr"


)
	
const BEFUNGE93_VERSION = "v1.0"

// error status
const (
	KErrorNone 					 		= 0
	KErrorUnableToOpenFile 		 		= 1
	KErrorMemoryBlockSizeToSmall 		= 2
	KErrorProblemParsingFile     		= 3
	KErrorSyntaxError			 		= 4
	KErrorBracketsNotPaired      		= 5
	KErrorUnexpectedClosingBracketFound = 6
	KErrorInvalidCharacterFound			= 7
)

const (
	DEFAULT_MEMORY_SIZE = 30000 // bytes
	DEFAULT_CODE_SIZE 	= 10000 // bytes
	DEFAULT_SOURCE_FILE = "A.BF" // to be changed
	MINIMUM_MMEORY_BLOCK_SIZE = 1000 // min menory size
)

var (
	sourceFileName string = DEFAULT_SOURCE_FILE
	memorySizePtr	     *int    
	stack 				intstack.IntStack
	instructionPtr		instructionptr.InstructionPtr
)

var (
	memoryPtr int			// Pointer to current location in memory
	memory []uint8
	instructions []uint8

	// useful limits and sizes
	codeSize int 		= DEFAULT_CODE_SIZE
	sourceFileSize int
	fileSize 	   int
	nestingLevel   int // count of the bracket nesting level from the root
)


func main () {
	displayBanner ()
	Initialise ()

stack.Push (65)
	interpreter.PopAndDisplayAsASCII (&stack)

	
	// read command line arguments, 'MEMORY' & 'FILE'
	flag.StringVar (&sourceFileName, "file", DEFAULT_SOURCE_FILE,"Name of default code file")

	
	memorySizePtr = flag.Int ("memory", DEFAULT_MEMORY_SIZE, "Working memory size in bytes")
	flag.Parse ()

	fmt.Printf ("Reading from source file '%s' ...\n", strings.ToUpper (sourceFileName))
	fmt.Printf ("Initialising memory block : %d bytes\n", *memorySizePtr)

	if *memorySizePtr <= MINIMUM_MMEORY_BLOCK_SIZE {
		showStatus (KErrorMemoryBlockSizeToSmall, "", *memorySizePtr)
	}

	// init memory block
	memory = make ([]uint8, *memorySizePtr, *memorySizePtr)
	
	fmt.Println ("Reading File ....")
	if data, err := readFileToMemory (sourceFileName); err == nil {
		codeSize = len(data)
		
		instructions = make ([]uint8, codeSize, codeSize)
	

		fmt.Println ("Parsing file ...")
		if status := parseFile (data); status != KErrorNone {
			showStatus (status, sourceFileName,0)
		} else {
			fmt.Println ("Executing file ...")
			if status := executeFile (); status != KErrorNone {
				showStatus (KErrorSyntaxError, sourceFileName,0)
			}
		}
	} else {
		showStatus (KErrorUnableToOpenFile, sourceFileName,0)
	}
}

// private methods
func displayBanner () {
	fmt.Printf ("BrainFuck Interpreter %s\n\n", BEFUNGE93_VERSION)
}

func Initialise () {
	stack.Reset ()
	instructionPtr.ResetInstructionPtr()
}


func readFileToMemory (filename string) ([]uint8, error) {
	file, err := os.Open (sourceFileName)
	defer file.Close ()
	
	if err == nil {
		data, err := ioutil.ReadAll (file)
		if err == nil {
	   		return data, nil
    	}
    }
	return nil, err	
}

func parseFile (data []uint8) int {
	return KErrorNone
}

func executeFile () int {
	return KErrorNone
}


func dump (){
	fmt.Printf ("%d:", instructionPtr)
	for i:=0; i <5; i++{
		fmt.Printf ("%02d ", memory[i])
	}
	fmt.Println ()
}

func showStatus (status int, extraInfoString string, extraInfoValue int) {

	var message string = "\nError: "

	if status != KErrorNone {
		switch (status) {
			case KErrorUnableToOpenFile: 
										message += fmt.Sprintf ("Unable to open file '%s'\n", strings.ToUpper(extraInfoString))
			case KErrorMemoryBlockSizeToSmall:
										message += fmt.Sprintf ("Specified memory size is too small (Minimum is %d bytes)\n", extraInfoValue) 
			case KErrorProblemParsingFile:
										message += fmt.Sprintf ("Problem parsing file '%s'\n", strings.ToUpper(extraInfoString))
			case KErrorSyntaxError:
										message += fmt.Sprintf ("Syntax error in file\n")
			case KErrorBracketsNotPaired:
										message += fmt.Sprintf ("Mismatched opening and closing brackets\n")
			case KErrorUnexpectedClosingBracketFound:
										message += fmt.Sprintf ("Closing backet found before any opening bracket\n")
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

