package main

import "github.com/rebirth-in-ruins/torpedodge/server/game"

type Strategy interface {
	nextMove(game.GameStateResponse) string
}
