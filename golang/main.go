// Package main implements a bot that circles around and drops a bomb occasionally.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/rebirth-in-ruins/torpedodge/server/game"
)

const (
	gameserverURL = "wss://gameserver.resamvi.io/play"
	playerName    = "iigorr"
)

func main() {
	//s := ManualStrategy{}
	//s.listen()

	s := NewAvoidBombStrategy()
	s.ManualStrategy.listen()

	if err := run(&s); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run(strat Strategy) error {
	conn, _, err := websocket.Dial(context.Background(), gameserverURL, nil)
	if err != nil {
		return fmt.Errorf("could not dial server: %w (url: %v)", err, gameserverURL)
	}
	defer conn.CloseNow()

	// Send initial JOIN message with your name
	err = wsjson.Write(context.Background(), conn, "JOIN "+playerName+".go")
	if err != nil {
		return fmt.Errorf("could not join server: %w", err)
	}

	i := 0

	for {
		// RECEIVE NEXT STATE
		var state game.GameStateResponse
		err := wsjson.Read(context.Background(), conn, &state)
		if err != nil {
			return fmt.Errorf("could not read from conn: %w", err)
		}

		action := strat.nextMove(state)
		slog.Info(action)

		// SEND ACTION
		err = wsjson.Write(context.Background(), conn, action)
		if err != nil {
			return fmt.Errorf("could not send next action: %w", err)
		}

		i++
	}
}
