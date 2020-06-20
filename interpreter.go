package main

package golang-befunge93

// extra stack operations
func (s *IntStack)AddTopStack () bool {
	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		s.Push (val1+val2)
		return true
	}

	return false
}

func (s *IntStack)SubtractTopStack () bool {
	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		s.Push (val2-val1)
		return true
	}

	return false
}

func (s *IntStack)MultiplyTopStack () bool {
	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		s.Push (val2*val1)
		return true
	}
	return false
}

func (s *IntStack)DivideTopStack () bool {
	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		if val1 == 0 {
			s.Push (0)
		} else {
			s.Push (val2/val1)
		}
		return true
	}
	return false
}

func (s *IntStack)DuplicateTop () bool {
	if s.Size() >= 1 {
		s.Push (s.Peek ())
		return true
	}
	return false
}

func (s *IntStack)ExchangeTop () bool {

	if s.Size () >= 2 {
		var val1 = s.Pop ()
		var val2 = s.Pop ()
		s.Push (val1)
		s.Push (val2)
		return true
	}
	return false
}

func (s* IntStack)PopAndDisplayAsInt () {
	var val = s.Pop
	fmt.Printf ("%d ", val1)
}

func (s* IntStack)PopAndDisplayAsASCII () {
	var val = s.Pop
	fmt.Printf ("%c ", val1)
}

