package player

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/position"
)

// Player is a participant in a match
type Player struct {
	color    sdl.Color
	strategy Strategy
	moves    []position.Position
}