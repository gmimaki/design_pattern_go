package main

import (
	"chain/trouble"
	"chain/support"
)

func main() {
	alice := support.NewNoSupport("Alice")
	bob := support.NewLimitSupport("Bob", 100)
	charlie := support.NewSpecialSupport("Charlie", 429)
	diana := support.NewLimitSupport("diana", 200)
	elmo := support.NewOddSupport("elmo")
	fred := support.NewLimitSupport("fred", 300)

	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)

	for i := 0; i < 500; i += 33 {
		alice.Handle(trouble.NewTrouble(i))
	}
}
