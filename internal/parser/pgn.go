package parser

import (
	"fmt"
	"strings"


	"github.com/notnil/chess"
)

type ParsedMove struct {
	FEN  string
	Move string
}

func ParsePGN(pgn string) ([]ParsedMove, error) {
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
	ret := make([]ParsedMove, len(game_moves))

	for i, j := range game_moves {
		ret[i] = ParsedMove{FEN: game_positions[i].String(), Move: j.String()}
	}

	return ret, nil
}
