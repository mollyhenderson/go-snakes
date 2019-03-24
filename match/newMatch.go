package match

import (
	"math/rand"
	"time"

	"github.com/wmurray8989/go-snakes/position"
)

// NewMatch creates a Match
func NewMatch(player1 Strategy, player2 Strategy) Match {
	var match = Match{}

	// random starting player
	startingPlayer := rand.Intn(1)
	if startingPlayer == 0 {
		match.playerTurn = false
	} else {
		match.playerTurn = true
	}

	// setup strategies
	match.player1 = player1
	match.player2 = player2

	// setup starting positions
	rand.Seed(time.Now().UTC().UnixNano())
	match.player1History = append(match.player1History, position.Position{X: rand.Intn(49), Y: rand.Intn(49)})
	match.player2History = append(match.player2History, position.Position{X: rand.Intn(49), Y: rand.Intn(49)})

	// setup colors
	match.p1Color.R = 255
	match.p1Color.A = 255

	match.p2Color.G = 255
	match.p2Color.A = 255

	match.gridColor.R = 100
	match.gridColor.G = 100
	match.gridColor.B = 100
	match.gridColor.A = 100

	return match
}
