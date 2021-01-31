package gamer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	me "../m"
)

type Gamer interface {
	GetMoney() int
	GetFruits() []string
	Bet()
	CreateMemento() me.Memento
	RestoreMemento(m me.Memento)
}

func NewGamer(money int) Gamer {
	return &gamer{
		money: money,
	}
}

type gamer struct {
	money  int
	fruits []string
}

func (g *gamer) fruitsName() []string {
	return []string{"リンゴ", "ぶどう", "バナナ", "みかん"}
}
func (g *gamer) GetMoney() int {
	return g.money
}
func (g *gamer) GetFruits() []string {
	return g.fruits
}
func (g *gamer) Bet() {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(6)
	i++

	switch i {
	case 1:
		g.money = g.money + 100
		fmt.Println("所持金が増えました")
	case 2:
		g.money = g.money / 2
		fmt.Println("所持金が半減しました")
	case 6:
		f := g.getFruit()
		fmt.Println(fmt.Sprintf("フルーツ %s をもらいました", f))
		g.fruits = append(g.fruits, f)
	default:
		fmt.Println("何も起こりませんでした")
	}
}

func (g *gamer) CreateMemento() me.Memento {
	m := me.NewMemento(g.money)
	// フルーツは美味しいものだけ
	for _, v := range g.fruits {
		if strings.Contains(v, "美味しい") {
			m.AddFruits(v)
		}
	}

	return m
}

func (g *gamer) RestoreMemento(m me.Memento) {
	g.money = m.GetMoney()
	g.fruits = m.GetFruits()
}

func (g *gamer) getFruit() string {
	prefix := ""
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(2)
	i++

	if i == 1 {
		prefix = "美味しい"
	}

	rand.Seed(time.Now().UnixNano())
	fruitsIndex := rand.Intn(len(g.fruitsName()))

	return prefix + g.fruitsName()[fruitsIndex]
}
