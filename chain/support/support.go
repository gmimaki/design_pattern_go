package support

import (
	"chain/trouble"
	"fmt"
)

type SupportInterface interface {
	Resolve(trouble.Trouble) bool
	SetNext(SupportInterface) SupportInterface
	Handle(trouble.Trouble)
}

type support struct {
	name string
	own  SupportInterface
	next SupportInterface
}

func (s *support) SetNext(next SupportInterface) SupportInterface {
	s.next = next
	fmt.Println(s.next)
	return next
}

func (s *support) Handle(t trouble.Trouble) {
	if s.own.Resolve(t) {
		s.done(t)
	} else if s.next != nil {
		s.next.Handle(t)
	} else {
		s.fail(t)
	}
}

func (s *support) toString() string {
	return fmt.Sprintf("[%s]", s.name)
}

func (s *support) Resolve(t trouble.Trouble) bool {
	return false
}

func (s *support) done(t trouble.Trouble) {
	fmt.Println(fmt.Sprintf("%s is resolved by %s", t.ToString(), s.toString()))
}

func (s *support) fail(t trouble.Trouble) {
	fmt.Println(fmt.Sprintf("%s cannot be resolved", t.ToString()))
}

func NewNoSupport(name string) *noSupport {
	noSupport := &noSupport{
		support: &support{
			name: name,
		},
	}
	noSupport.own = noSupport
	return noSupport
}

type noSupport struct {
	*support
}

func (s *noSupport) Resolve(t trouble.Trouble) bool {
	fmt.Println("No Support")
	return false
}

func NewLimitSupport(name string, limit int) *limitSupport {
	limitSupport := &limitSupport{
		support: &support{
			name: name,
		},
		limit: limit,
	}
	limitSupport.own = limitSupport
	return limitSupport
}

type limitSupport struct {
	*support
	limit int
}

func (s *limitSupport) resolve(t trouble.Trouble) bool {
	return t.GetNumber() < s.limit
}

func NewOddSupport(name string) *oddSupport {
	oddSupport := &oddSupport{
		support: &support{
			name: name,
		},
	}
	oddSupport.own = oddSupport
	return oddSupport
}

type oddSupport struct {
	*support
}

func (s *oddSupport) Resolve(t trouble.Trouble) bool {
	return t.GetNumber()%2 == 1
}

func NewSpecialSupport(name string, number int) *specialSupport {
	specialSupport := &specialSupport{
		support: &support{
			name: name,
		},
		number: number,
	}
	specialSupport.own = specialSupport
	return specialSupport
}

type specialSupport struct {
	*support
	number int
}

func (s *specialSupport) Resolve(t trouble.Trouble) bool {
	fmt.Println("Special Support")
	fmt.Println(t.GetNumber())
	return t.GetNumber() == s.number
}
