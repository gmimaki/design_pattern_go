package main

type Mediator interface {
	createColleagues()
	colleagueChanged()
}

type Colleague interface {
	setMediator(Mediator)
	setColleagueEnabled(bool)
}

type colleagueButton struct {
	mediator Mediator
	enabled  bool
}

func (s *colleagueButton) setMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *colleagueButton) setColleagueEnabled(enabled bool) {
	s.enabled = enabled
}

type colleagueTextField struct {
	mediator   Mediator
	enabled    bool
	background string
}

func (s *colleagueTextField) setMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *colleagueTextField) setColleagueEnabled(enabled bool) {
	s.enabled = enabled
	if enabled {
		s.background = "white"
	} else {
		s.background = "lightgray"
	}
}

func (s *colleagueTextField) textValueChanged(e string) { // action型として定義する？
	s.mediator.colleagueChanged()
}

type colleaguCheckbox struct {
	mediator Mediator
	enabled  bool
}

func (s *colleaguCheckbox) setMediator(mediator Mediator) {
	s.mediator = mediator
}
func (s *colleaguCheckbox) setColleagueEnabled(enabled bool) {
	s.enabled = enabled
}
func (s *colleaguCheckbox) itemStateChanged(e string) {
	s.mediator.colleagueChanged()
}

type loginFrame struct {
}
