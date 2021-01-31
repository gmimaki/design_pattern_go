package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"./gamer"
)

func main() {
	gamer := gamer.NewGamer(100) // 所持金100でスタート
	memento := gamer.CreateMemento()

	for i := 0; i < 100; i++ {
		fmt.Println("=== " + strconv.Itoa(i))
		fmt.Println(fmt.Sprintf("現状: money=%d fruits=%s", gamer.GetMoney(), strings.Join(gamer.GetFruits(), ", ")))

		gamer.Bet()
		fmt.Println(fmt.Sprintf("所持金は%d円になりました", gamer.GetMoney()))

		if gamer.GetMoney() > memento.GetMoney() {
			fmt.Println("	所持金が増えたので現在の状態を保存しておこう")
			memento = gamer.CreateMemento()
		} else if gamer.GetMoney() < memento.GetMoney()/2 {
			fmt.Println("	だいぶ減ったので、以前の状態に復元しよう")
			gamer.RestoreMemento(memento)
		}

		time.Sleep(time.Second * 1)
		fmt.Println("")
	}
}
