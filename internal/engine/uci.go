package engine

import (
	"fmt"
	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"

	gameanalyzer "github.com/Chesstutis/GameAnalyzer"
)


type EngineAnalysis struct {
	BestMove   gameanalyzer.EvaluatedMove
	PlayerMove gameanalyzer.EvaluatedMove
}

// finds the best move in the position as well as the analysis for the players move
func AnalyzePosition(fen string, playerMove string, stockfishPath string) (*EngineAnalysis, error) {
	eng, err := uci.New(stockfishPath)
	if err != nil {
		return nil, fmt.Errorf("failed to start engine: %w", err)
	}
	defer eng.Close()

	// 1. analyze engine move

	// 2. analyve player move

	return nil, fmt.Errorf("not implemented")
}

// finds the chess.Move object corresponding to the players move
func parseMove(pos *chess.Position, moveStr string) (chess.Move, error) {

	return chess.Move{}, fmt.Errorf("not implemented")
}