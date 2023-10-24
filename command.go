package main

import "fmt"

type Command interface {
	execute()
}

type SitDown struct {
}

func (c1 *SitDown) execute() {
	fmt.Println("Sit down, motherfucker")
}

type StandUp struct {
}

func (c2 *StandUp) execute() {
	fmt.Println("Stand Up, motherfucker")

}

type Button struct {
	command Command
}

func (b Button) press() {
	b.command.execute()
}

type RemoteController struct {
	sitDown Button
	standUp Button
}

func testCommand() {
	downCommand := &SitDown{}
	downButton := Button{downCommand}

	upCommand := &StandUp{}
	upButton := Button{upCommand}

	controller := RemoteController{downButton, upButton}

	controller.standUp.press()
	controller.sitDown.press()

}
