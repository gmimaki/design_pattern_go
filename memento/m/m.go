package m

type Memento interface {
	GetMoney() int //本当はこれだけpublic
	AddFruits(string)
	GetFruits() []string
}

func NewMemento(money int) Memento {
	return &memento{
		money: money,
	}
}

type memento struct {
	money  int
	fruits []string
}

func (m *memento) GetMoney() int {
	return m.money
}
func (m *memento) AddFruits(fruit string) {
	m.fruits = append(m.fruits, fruit)
}
func (m *memento) GetFruits() []string {
	return m.fruits
}
