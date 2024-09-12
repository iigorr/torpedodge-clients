// Package main implements a bot that circles around and drops a bomb occasionally.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/rebirth-in-ruins/torpedodge/server/game"
)

const (
	urlStr = "wss://gameserver.resamvi.io/play"
)

var (
	directions = []string{"LEFT", "BOMBLEFT", "DOWN", "DOWN", "RIGHT", "RIGHT", "UP", "UP"}

	name = []string{
		"Ella Atkinson",
		"Duke Krueger",
		"Kamari Solis",
		"Ronin Santana",
		"Myra Fuller",
		"Andre Montes",
		"Roselyn Frederick",
		"Kase Gregory",
		"Alaya Herman",
		"Juelz Frost",
		"Paula Reyna",
		"Reginald Kirby",
		"Skyla Holt",
		"Niko Garrison",
		"Cadence Graham",
	}
)

func main() {
	for {
		_ = run() // When we died and lost connection, retry again.
	}
}

func run() error {
	conn, _, err := websocket.Dial(context.Background(), urlStr, nil)
	if err != nil {
		return fmt.Errorf("could not dial: %w", err)
	}
	defer conn.CloseNow()

	// Join as player
	name := name[rand.IntN(len(name))]
	err = wsjson.Write(context.Background(), conn, "JOIN "+name)
	if err != nil {
		return fmt.Errorf("could not write: %w", err)
	}

	i := 0

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 12 * time.Second)

		// RECEIVE NEXT STATE
		var state game.GameStateResponse
		err = wsjson.Read(ctx, conn, &state)
		if err != nil {
			return fmt.Errorf("could not write: %w", err)
		}

		// Sail in a circle
		action := directions[i % len(directions)]

		slog.Info(action)

		// SEND ACTION
		err = wsjson.Write(ctx, conn, action)
		if err != nil {
			return fmt.Errorf("could not write: %w", err)
		}

		i++
		cancel()
	}
}

