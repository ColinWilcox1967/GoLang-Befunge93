package main

import (
"fmt"
"os"
"flag"
"io/ioutil"
"strings"

"github.com/colinwilcox1967/intstack"
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
	DEFAULT_SOURCE_FILE = "TEST.B93"
)

var (
	sourceFileName string = DEFAULT_SOURCE_FILE
	stack 				intstack.IntStack
	instructionPtr		instructionptr.InstructionPtr
)

var (
	// useful limits and sizes
	sourceFileSize int
	fileSize 	   int
)


func main () {
	displayBanner ()
	Initialise ()

	
	// read command line arguments, 'MEMORY' & 'FILE'
	flag.StringVar (&sourceFileName, "file", DEFAULT_SOURCE_FILE,"Name of default code file")
	flag.Parse ()

	fmt.Printf ("Reading from source file '%s' ...\n", strings.ToUpper (sourceFileName))

	if data, err := readFileToMemory (sourceFileName); err == nil {
			
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
	fmt.Printf ("Befunge93 Interpreter %s\n\n", BEFUNGE93_VERSION)
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

