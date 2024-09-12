package main

import "github.com/rebirth-in-ruins/torpedodge/server/game"

type CircleSailStrategy struct {
	i int
}

var (
	circleDirections = []string{"LEFT", "BOMBLEFT", "DOWN", "DOWN", "RIGHT", "RIGHT", "UP", "UP"}
)

func (s *CircleSailStrategy) nextMove(strat game.GameStateResponse) string {
	s.i++
	return circleDirections[s.i%len(circleDirections)]
}
