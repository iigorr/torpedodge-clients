# Starter templates for Torpedodge Bots

Starter template to write a bot for Torpedodge.  
Connects to the game server at `wss://gameserver.resamvi.io` and sails in a circle.

## Golang
Start with
```
go run main.go
```

## JavaScript
Start with
```
npm install && npm run start
```

## Kotlin

Run with
```
./gradlew run
```

## Python

Run with
```
poetry install && poetry run python main.py
```

## Gamestate

The complete JSON that is sent on every turn:

```json
{
  // current players connected on the battlefield
  "players": [
    {
      "id": 28,
      "name": "Niko Garrison", // player name
      "x": 1, // position
      "y": 0, // position
      "rotation": "LEFT", // which direction he's looking at
      "charging": false, // if true is about to shoot a laser
      "team": "none", // determines the ship skin (e.g."golang", "python")
      "score": 17, // current score
      "health": 1, // remaining hit points
      "bombCount": 2, // amount of bombs in inventory
      "bombRespawn": 0
    }
  ],

  // random airstrikes that spawn out of the air
  "airstrikes": [
    {
      "id": 10088,
      "x": 3, // position
      "y": 10, // position
      "fuseCount": 2 // remaining turns until it explodes
    },
    {
      "id": 10063,
      "x": 8,
      "y": 9,
      "fuseCount": 1
    }
  ],

  // explosions in this turn that hurt ships when on this position
  "explosions": [
    {
      "id": 10091,
      "x": 6, // position
      "y": 1, // position
      "playerId": 1 // associated player that created the explosion
    },
  ],

  // bombs dropped by players
  "bombs": [
     {
      "id": 16139,
      "playerId": 29, // associated player
      "x": 9, // position
      "y": 7, // position
      "fuseCount": 1 // remaining turns until it explodes
    }
  ],

  // corpses are dead players that remain for a few more turns
  "corpses": [
    {
      "id": 17090,
      "name": "JulienBot", // name of player that died
      "x": 5, // position
      "y": 0, // position
      "rotation": "LEFT", // in which direction it is looking
      "DeathTimer": 1 // remaining turns visible
    }
  ],

  // items that can be picked up for score
  "loot": [
    {
      "id": 4,
      "type": "good", // (browser-only) determines the texture used 
      "value": 12, // how many points in grants on pickup
      "x": 5, // position
      "y": 10 // position
    },
    {
      "id": 1,
      "type": "mediocre",
      "value": 6,
      "x": 11,
      "y": 4
    }
  ],

  // scores of each player
  "leaderboard": [
    {
      "name": "JulienBot",
      "score": 9
    },
    {
      "name": "Duke Krueger",
      "score": 4
    }
  ],

  // highest achieved scores in the game
  "kings": [
    {
      "name": "Cadence Graham",
      "score": "55",
      "created_at": "2024-09-11T19:19:23.192214Z"
    },
    {
      "name": "Andre Montes",
      "score": "35",
      "created_at": "2024-09-11T18:49:06.795141Z"
    },
    {
      "name": "Skyla Holt",
      "score": "28",
      "created_at": "2024-09-11T19:18:18.097312Z"
    }
  ],

  // (for browser-client) mentions additional animations that should be rendered
  "animations": [],

  // recent events that occured
  "events": [
    "Kase Gregory died",
    "Kase Gregory took a hit",
    "Alaya Herman died",
    "Alaya Herman took a hit",
    "Alaya Herman took a hit"
  ],

  // the game's settings (they don't change between turns)
  "settings": {
    "turnDuration": 3000000000, // 3 seconds
    "gridSize": 12, // width and height of grid
    "inventorySize": 2, // amount of bombs a player can carry
    "bombRespawnTime": 3, // (unused)
    "startHealth": 3, // start health
    "airstrikeFuseLength": 3, // how many turns until an airstrike detonates
    "bombFuseLength": 3, // how many turns until a bomb detonates
    "deathTime": 3, // amount of turns that a corpse remains
    "locked": false, // if true no new players can join
    "paused": false // if true no turns will occur, but new players can join 
  }
}
```
