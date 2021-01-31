package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Memento interface {
	getMoney() int //本当はこれだけpublic
	addFruits(string)
	getFruits() []string
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

func (m *memento) getMoney() int {
	return m.money
}
func (m *memento) addFruits(fruit string) {
	m.fruits = append(m.fruits, fruit)
}
func (m *memento) getFruits() []string {
	return m.fruits
}

type Gamer interface {
	getMoney() int
	getFruits() []string
	bet()
	createMemento() Memento
	restoreMemento(m Memento)
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
func (g *gamer) getMoney() int {
	return g.money
}
func (g *gamer) getFruits() []string {
	return g.fruits
}
func (g *gamer) bet() {
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

func (g *gamer) createMemento() Memento {
	m := NewMemento(g.money)
	// フルーツは美味しいものだけ
	for _, v := range g.fruits {
		if strings.Contains(v, "美味しい") {
			m.addFruits(v)
		}
	}

	return m
}

func (g *gamer) restoreMemento(m Memento) {
	g.money = m.getMoney()
	g.fruits = m.getFruits()
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

func main() {
	gamer := NewGamer(100) // 所持金100でスタート
	memento := gamer.createMemento()

	for i := 0; i < 100; i++ {
		fmt.Println("=== " + strconv.Itoa(i))
		fmt.Println(fmt.Sprintf("現状: money=%d fruits=%s", gamer.getMoney(), strings.Join(gamer.getFruits(), ", ")))

		gamer.bet()
		fmt.Println(fmt.Sprintf("所持金は%d円になりました", gamer.getMoney()))

		if gamer.getMoney() > memento.getMoney() {
			fmt.Println("	所持金が増えたので現在の状態を保存しておこう")
			memento = gamer.createMemento()
		} else if gamer.getMoney() < memento.getMoney()/2 {
			fmt.Println("	だいぶ減ったので、以前の状態に復元しよう")
			gamer.restoreMemento(memento)
		}

		time.Sleep(time.Second * 1)
		fmt.Println("")
	}
}
