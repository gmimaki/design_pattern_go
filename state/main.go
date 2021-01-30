package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	ds = NewDayState()
	ns = NewNightState()
)

// State
type State interface {
	doClock(Context, int)
	doUse(Context)
	doAlarm(Context)
	doPhone(Context)
	toString() string
}

// NewDayState
func NewDayState() State {
	return &dayState{}
}

type dayState struct {
}

func (s *dayState) doClock(ctx Context, hour int) {
	if hour < 9 || hour >= 17 {
		ctx.changeState(ns)
	}
}

func (s *dayState) doUse(ctx Context) {
	ctx.recordLog("金庫使用(昼間)")
}
func (s *dayState) doAlarm(ctx Context) {
	ctx.callSecurityCenter("非常ベル(昼間)")
}
func (s *dayState) doPhone(ctx Context) {
	ctx.callSecurityCenter("通常の通話(昼間)")
}
func (s *dayState) toString() string {
	return "昼間"
}

func NewNightState() State {
	return &nightState{}
}

type nightState struct {
}

func (s *nightState) doClock(ctx Context, hour int) {
	if 9 <= hour && hour < 17 {
		ctx.changeState(ds)
	}
}

func (s *nightState) doUse(ctx Context) {
	ctx.callSecurityCenter("非常:夜間の金庫使用！")
}
func (s *nightState) doAlarm(ctx Context) {
	ctx.callSecurityCenter("非常ベル(夜間)")
}
func (s *nightState) doPhone(ctx Context) {
	ctx.recordLog("夜間の通話録音")
}
func (s *nightState) toString() string {
	return "夜間"
}

// Context
type Context interface {
	setClock(int)
	changeState(State)
	callSecurityCenter(string)
	recordLog(string)
	actionPerform(actionSource)
}

type actionSource string

const (
	buttonUse   actionSource = "buttonUse"
	buttonAlarm actionSource = "buttonAlarm"
	buttonPhone actionSource = "buttonPhone"
)

var (
	buttons = []actionSource{buttonUse, buttonAlarm, buttonPhone}
)

func NewSafeFrame(state State, msg string) Context {
	return &safeFrame{
		state:      state,
		textClock:  "",
		textScreen: msg,
	}
}

type safeFrame struct {
	state      State
	textClock  string
	textScreen string
}

func (s *safeFrame) setClock(hour int) {
	clockString := "現在時刻は"
	if hour < 10 {
		clockString = fmt.Sprintf("%s0%d:00", clockString, hour)
	} else {
		clockString = fmt.Sprintf("%s%d:00", clockString, hour)
	}

	fmt.Println(clockString)
	s.textClock = clockString
	s.state.doClock(s, hour)
}

func (s *safeFrame) changeState(state State) {
	fmt.Println(fmt.Sprintf("%sから%sへ状態が変化しました", s.state.toString(), state.toString()))
	s.state = state
}
func (s *safeFrame) callSecurityCenter(message string) {
	s.textScreen = fmt.Sprintf(`call!! %s`, message)
	fmt.Println(s.textScreen)
}
func (s *safeFrame) recordLog(message string) {
	s.textScreen = fmt.Sprintf(`record... %s`, message)
	fmt.Println(s.textScreen)
}

func (s *safeFrame) actionPerform(action actionSource) {
	switch action {
	case buttonUse:
		s.state.doUse(s)
	case buttonAlarm:
		s.state.doAlarm(s)
	case buttonPhone:
		s.state.doPhone(s)
	}
}

func main() {
	frame := NewSafeFrame(ds, "State Sample")

	for true {
		for hour := 0; hour < 24; hour++ {
			frame.setClock(hour)

			time.Sleep(time.Second * 2)

			s := rand.NewSource(time.Now().Unix())
			r := rand.New(s)

			frame.actionPerform(buttons[r.Intn(len(buttons))])
		}
	}
}
