package engine

import (
	"fmt"
	gameanalyzer "github.com/Chesstutis/GameAnalyzer"
)


type EngineAnalysis struct {
	BestMove   gameanalyzer.EvaluatedMove
	PlayerMove gameanalyzer.EvaluatedMove
}

// finds the best move in the position as well as the analysis for the players move
func AnalyzePosition(fen string, playerMove string, stockfishPath string) (*EngineAnalysis, error) {
	return nil, fmt.Errorf("not implemented")
}

// finds the chess.Move object corresponding to the players move
func parseMove(pos *chess.Position, moveStr string) (chess.Move, error) {

	return nil, fmt.Errorf("not implemented")
}