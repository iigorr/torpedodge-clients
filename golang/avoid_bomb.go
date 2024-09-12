package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/davecgh/go-spew/spew"
	"github.com/rebirth-in-ruins/torpedodge/server/game"
)

type Point struct {
	X, Y   int
	rating float64
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y, rating: 0}
}

type AvoidBombStrategy struct {
	ManualStrategy

	me       *game.Player
	nextPos  Point
	gridSize int
}

func NewAvoidBombStrategy() AvoidBombStrategy {
	return AvoidBombStrategy{
		ManualStrategy: NewManualStrategy(),
	}
}

func (s *AvoidBombStrategy) updateData(state game.GameStateResponse) error {
	for _, player := range state.Players {
		if player.Name == playerName {
			s.me = &player
			break
		}
	}
	if s.me == nil {
		return fmt.Errorf("player with name %s not found", playerName)
	}

	s.gridSize = state.Settings.GridSize

	s.calcNextPos()
	return nil
}

func (s AvoidBombStrategy) point2Dir(dest Point) string {
	dX := dest.X - s.me.X
	if dX == 1 {
		return "RIGHT"
	} else if dX == -1 {
		return "LEFT"
	}

	dY := dest.Y - s.me.Y
	if dY == 1 {
		return "DOWN"
	} else if dY == -1 {
		return "UP"
	}

	return "STOP"
}

func (s *AvoidBombStrategy) calcNextPos() {
	s.nextPos = NewPoint(s.me.X, s.me.Y)

	switch s.me.Rotation {
	case "LEFT":
		s.nextPos.X = max(0, s.me.X-1)
	case "RIGHT":
		s.nextPos.X = min(s.gridSize-1, s.me.X+1)
	case "UP":
		s.nextPos.Y = max(0, s.me.Y-1)
	case "DOWN":
		s.nextPos.Y = min(s.gridSize-1, s.me.Y+1)
	}
}

func (s AvoidBombStrategy) possibleMoves() []Point {
	moves := []Point{}

	if s.me.X != 0 {
		moves = append(moves, NewPoint(s.me.X-1, s.me.Y))
	}
	if s.me.X != s.gridSize-1 {
		moves = append(moves, NewPoint(s.me.X+1, s.me.Y))
	}
	if s.me.Y != 0 {
		moves = append(moves, NewPoint(s.me.X, s.me.Y-1))
	}
	if s.me.Y != s.gridSize-1 {
		moves = append(moves, NewPoint(s.me.X, s.me.Y+1))
	}

	return moves
}

func (s AvoidBombStrategy) rateMove(move Point, state game.GameStateResponse) float64 {
	score := float64(0)

	hitX, hitY := hitMap(state, 1)
	spew.Printf("hitX: %v\n", hitX)
	spew.Printf("hitY: %v\n", hitY)

	if hitX[move.X] {
		spew.Printf("X-Hit on point (%d, %d)\n", move.X, move.Y)
		return float64(-1)
	}

	if hitY[move.Y] {
		spew.Printf("Y-Hit on point (%d, %d)\n", move.X, move.Y)
		return float64(-1)
	}

	return score
}

func hitMap(state game.GameStateResponse, fuse int) (hitX []bool, hitY []bool) {
	hitX = make([]bool, state.Settings.GridSize)
	hitY = make([]bool, state.Settings.GridSize)

	for _, airstrike := range state.Airstrikes {
		if airstrike.FuseCount == fuse {
			hitX[airstrike.X] = true
			hitY[airstrike.Y] = true

			continue
		}

		if hitX[airstrike.X] {
			hitY[airstrike.Y] = true
		}
		if hitY[airstrike.Y] {
			hitX[airstrike.X] = true
		}
	}

	for _, bomb := range state.Bombs {
		if bomb.FuseCount == fuse {
			hitX[bomb.X] = true
			hitY[bomb.Y] = true

			continue
		}

		if hitX[bomb.X] {
			hitY[bomb.Y] = true
		}
		if hitY[bomb.Y] {
			hitX[bomb.X] = true
		}
	}

	return
}

func (s *AvoidBombStrategy) nextMove(state game.GameStateResponse) string {
	err := s.updateData(state)
	if err != nil {
		fmt.Println(err)
		return "QUIT"
	}

	// spew.Printf("Orientation: %s\n", s.me.Rotation)

	spew.Dump(s.me)
	spew.Dump(state.Airstrikes)
	spew.Dump(state.Bombs)

	if s.ManualStrategy.hasInput {
		return s.ManualStrategy.nextMove(state)
	}

	moves := s.possibleMoves()
	randomPick := moves[rand.IntN(len(moves))]

	spew.Printf("Pos (%d, %d)\n", s.me.X, s.me.Y)
	spew.Printf("Next Move: (%d, %d)\n", randomPick.X, randomPick.Y)

	moveRating := s.rateMove(randomPick, state)
	spew.Printf("Rating of Move %f\n", moveRating)

	return s.point2Dir(randomPick)
}
