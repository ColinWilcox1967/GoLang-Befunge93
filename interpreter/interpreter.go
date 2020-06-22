package befunge93interpreter

import (
	"fmt"
	intstack "github.com/colinwilcox1967/intstack"
)

// extra stack operations
func AddTopStack (s *intstack.IntStack) bool {
	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		s.Push (val1+val2)
		return true
	}

	return false
}

func SubtractTopStack (s *intstack.IntStack) bool {
	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		s.Push (val2-val1)
		return true
	}

	return false
}

func MultiplyTopStack (s *intstack.IntStack) bool {
	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		s.Push (val2*val1)
		return true
	}
	return false
}

func DivideTopStack (s *intstack.IntStack) bool {
	if s.Size () >= 2 {
		val1 := s.Pop ()
		val2 := s.Pop ()

		if val1 == 0 {
			s.Push (0)
		} else {
			s.Push (val2/val1)
		}
		return true
	}
	return false
}

func DuplicateTop (s *intstack.IntStack) bool {
	if s.Size() >= 1 {
		s.Push (s.Peek ())
		return true
	}
	return false
}

func ExchangeTop (s *intstack.IntStack) bool {

	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		s.Push (val1)
		s.Push (val2)
		return true
	}
	return false
}

func PopAndDisplayAsInt (s *intstack.IntStack) {
	var val = s.Pop ()
	fmt.Printf ("%d ", val)
}

func PopAndDisplayAsASCII (s *intstack.IntStack) {
	var val = s.Pop ()
	fmt.Printf ("%c ", val)
}

