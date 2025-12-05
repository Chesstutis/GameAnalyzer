package parser

import (
	"fmt"
	"github.com/notnil/chess"
	"strings"
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
	game_postions := game.Positions()
	game_moves := game.Moves()
	ret := make([]ParsedMove, len(game_moves)+1)

	for i , j := range game_moves{
		ret[i] = ParsedMove{FEN: game_postions[i + 1].String(), Move: j.String()}
	}
	
	return ret, nil
}
