package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Observer interface {
	Update(numberGenerator)
}

type NumberGenerator interface {
	addObserver(Observer)
	deleteObserver(Observer)
	notifyObservers()
	getNumber() int
	execute()
}

func NewNumberGenerator() NumberGenerator {
	return &numberGenerator{}
}

type numberGenerator struct {
	observers []Observer
	number    int
}

func (s *numberGenerator) addObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}
func (s *numberGenerator) deleteObserver(observer Observer) {
	newObservers := []Observer{}
	for _, v := range s.observers {
		if !(v == observer) { // TODO どうやったらイコール確認できる？
			newObservers = append(newObservers, v)
		}
	}
}
func (s *numberGenerator) notifyObservers() {
	for _, v := range s.observers {
		v.Update(*s)
	}
}
func (s *numberGenerator) getNumber() int {
	return s.number
}
func (s *numberGenerator) execute() {
}

func NewRandomNumberGenerator() NumberGenerator {
	return &randomNumberGenerator{}
}

type randomNumberGenerator struct {
	numberGenerator
}

func (s *randomNumberGenerator) getNumber() int {
	return s.number
}
func (s *randomNumberGenerator) execute() {
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		s.numberGenerator.number = rand.Intn(50)
		s.notifyObservers()
	}
}

func NewDigitObserver() Observer {
	return &digitObserver{}
}

type digitObserver struct {
}

func (s *digitObserver) Update(generator numberGenerator) {
	fmt.Println(fmt.Sprintf("DigitObserver: %d", generator.getNumber()))
	time.Sleep(time.Second * 1)
}

func NewGraphObserver() Observer {
	return &graphObserver{}
}

type graphObserver struct {
}

func (s *graphObserver) Update(generator numberGenerator) {
	fmt.Println("GraphObserver:")
	count := generator.getNumber()
	for i := 0; i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println("")
	time.Sleep(time.Second * 1)
}

func main() {
	generator := NewRandomNumberGenerator()
	observer1 := NewDigitObserver()
	observer2 := NewGraphObserver()
	generator.addObserver(observer1)
	generator.addObserver(observer2)

	generator.execute()
}
