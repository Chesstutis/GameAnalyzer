package parser

import (
	"fmt"
	"strings"

	"github.com/notnil/chess"

	"github.com/Chesstutis/GameAnalyzer/internal/types"
)

func ParsePGN(pgn string) ([]types.ParsedMove, error) {
	if pgn == "" {
		return nil, fmt.Errorf("empty PGN")
	}

	reader := strings.NewReader(pgn)
	parsedPgn, err := chess.PGN(reader)

	if err != nil {
		return nil, fmt.Errorf("problem reading FEN")
	}

	game := chess.NewGame(parsedPgn)
	game_positions := game.Positions()
	game_moves := game.Moves()
	ret := make([]types.ParsedMove, len(game_moves))

	for i, j := range game_moves {
		ret[i] = types.ParsedMove{Position: *game_positions[i], Move: *j}
	}

	return ret, nil
}
