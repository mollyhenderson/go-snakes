package player

import (
	"errors"
	"time"

	"github.com/wmurray8989/go-snakes/position"
)

const sideLength = 50
const cellSize = 10

// NextMove returns the player's next move
func (p *Player) NextMove(opponent Player) (err error) {
	// Check if no moves are available
	currentPosition := p.moves[len(p.moves)-1]
	validMoves := currentPosition.GetValidMoves(p.moves, opponent.moves)
	if len(validMoves) == 0 {
		return errors.New("No Valid Moves")
	}

	// timeout after 1 second
	timeout := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	// get next move
	ch := make(chan position.Position)
	go func() {
		// Handle panics
		defer func() {
			if r := recover(); r != nil {
				err = errors.New("Panicked")
				ch <- position.Null
			}
		}()

		move := p.strategy(p.moves, opponent.moves)
		if !(move.IsValidMove(p.moves, opponent.moves)) {
			err = errors.New("Invalid Move")
		}
		ch <- move
	}()

	select {
	case move := <-ch:
		// a read from ch has occurred
		p.moves = append(p.moves, move)
	case <-timeout:
		err = errors.New("Timeout")
	}

	return err
}
