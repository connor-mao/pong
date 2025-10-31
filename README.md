# Pong Game with Ebiten

## Overview

This is a simple Pong game implemented in Go using the Ebiten library. It features a moving ball, played-controlled paddles, scoring system, and various gameplay enhancements.

![Demo](https://media.giphy.com/media/7BZ7GROeJgHRhJ3L90/giphy.gif)


## Features

* **Scoreboard**: Keeps track of player scores and displays them on the screen.
* **Color-Changing Ball**: The ball changes to a random color each time it hits a paddle.
* **Pause Mechanic**: The game pauses briefly whenever the ball hits the wall, adding a subtle delay to ensure both players are ready to continue playing.

## Controls

* **Player Paddle**: Player 1 uses W/S, and Player 2 uses Up/Down Arrow to move the paddles. 

## Installation

1. Make sure you have [Go](https://go.dev/doc/install) installed.
2. Install Ebiten:

```bash
go get github.com/hajimehoshi/ebiten/v2
go get github.com/hajimehoshi/ebiten/v2/ebitenutil
```

3. Clone this repository and navigate to the project folder.

## Running the Game

```bash
go run main.go
```

