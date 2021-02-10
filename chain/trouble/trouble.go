package trouble

import "fmt"

func NewTrouble(number int) Trouble {
	return &trouble{
		number: number,
	}
}

type Trouble interface {
	GetNumber() int
	ToString() string
}

type trouble struct {
	number int
}

func (s *trouble) GetNumber() int {
	return s.number
}

func (s *trouble) ToString() string {
	return fmt.Sprintf("[Trouble %d]", s.number)
}