package main

import "fmt"

func NewTrouble(number int) Trouble {
	return &trouble{
		number: number,
	}
}

type Trouble interface {
	getNumber() int
	toString() string
}

type trouble struct {
	number int
}

func (s *trouble) getNumber() int {
	return s.number
}

func (s *trouble) toString() string {
	return fmt.Sprintf("[Trouble %d]", s.number)
}

type SupportInterface interface {
	resolve(Trouble) bool
	setNext(SupportInterface) SupportInterface
	handle(Trouble)
}

type support struct {
	name string
	own  SupportInterface
	next SupportInterface
}

func (s *support) setNext(next SupportInterface) SupportInterface {
	s.next = next
	fmt.Println(s.next)
	return next
}

func (s *support) handle(t Trouble) {
	if s.own.resolve(t) {
		s.done(t)
	} else if s.next != nil {
		s.next.handle(t)
	} else {
		s.fail(t)
	}
}

func (s *support) toString() string {
	return fmt.Sprintf("[%s]", s.name)
}

func (s *support) resolve(t Trouble) bool {
	return false
}

func (s *support) done(t Trouble) {
	fmt.Println(fmt.Sprintf("%s is resolved by %s", t.toString(), s.toString()))
}

func (s *support) fail(t Trouble) {
	fmt.Println(fmt.Sprintf("%s cannot be resolved", t.toString()))
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

func (s *noSupport) resolve(t Trouble) bool {
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

func (s *limitSupport) resolve(t Trouble) bool {
	return t.getNumber() < s.limit
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

func (s *oddSupport) resolve(t Trouble) bool {
	return t.getNumber()%2 == 1
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

func (s *specialSupport) resolve(t Trouble) bool {
	fmt.Println("Special Support")
	fmt.Println(t.getNumber())
	return t.getNumber() == s.number
}

func main() {
	alice := NewNoSupport("Alice")
	bob := NewLimitSupport("Bob", 100)
	charlie := NewSpecialSupport("Charlie", 429)
	diana := NewLimitSupport("diana", 200)
	elmo := NewOddSupport("elmo")
	fred := NewLimitSupport("fred", 300)

	alice.setNext(bob).setNext(charlie).setNext(diana).setNext(elmo).setNext(fred)

	for i := 0; i < 500; i += 33 {
		alice.handle(NewTrouble(i))
	}
}
