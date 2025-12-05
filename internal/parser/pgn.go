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

func Parse(pgn string) ([]ParsedMove, error) {
	// js like parse type scriiiipt
	reader := strings.NewReader(pgn)
	PGN, err := chess.PGN(reader)
	
	if err != nil {
		fmt.Println(err)
	}
	
	game := chess.NewGame(PGN)
	game_postions := game.Positions()
	game_moves := game.Moves()
	ret := make([]ParsedMove, len(game_moves)+1)

	for i , j := range game_moves{
		ret[i] = ParsedMove{FEN: game_postions[i + 1].String(), Move: j.String()}
	}

	
	return ret, fmt.Errorf("not Implemented")
}
