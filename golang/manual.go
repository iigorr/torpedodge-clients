package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/rebirth-in-ruins/torpedodge/server/game"
)

type ManualStrategy struct {
	nextMovement string
	doBomb       bool
	hasInput     bool
}

func NewManualStrategy() ManualStrategy {
	return ManualStrategy{nextMovement: "LEFT", doBomb: false}
}

func (s *ManualStrategy) listen() {
	go func() {
		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer func() {
			_ = keyboard.Close()
		}()

		fmt.Println("Press ESC to quit")
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)

			switch key {
			case keyboard.KeyEsc:
				s.nextMovement = "QUIT"
			case keyboard.KeyArrowDown:
				s.nextMovement = "DOWN"
			case keyboard.KeyArrowUp:
				s.nextMovement = "UP"
			case keyboard.KeyArrowLeft:
				s.nextMovement = "LEFT"
			case keyboard.KeyArrowRight:
				s.nextMovement = "RIGHT"
			case 0x00:
				if char == 'h' {
					s.nextMovement = ""
				} else if char == 'x' {
					s.nextMovement = "LASER"
				}
			case keyboard.KeySpace:
				s.doBomb = !s.doBomb

			}
			s.hasInput = true
			fmt.Printf("Next Move %s\r\n", s.calcAction())

			if key == keyboard.KeyEsc {
				break
			}
		}
	}()

}

func (s ManualStrategy) calcAction() string {
	if s.doBomb {
		return "BOMB" + s.nextMovement
	} else {
		return s.nextMovement
	}
}

func (s *ManualStrategy) nextMove(strat game.GameStateResponse) string {
	action := s.calcAction()

	s.doBomb = false
	s.hasInput = false

	return action
}
