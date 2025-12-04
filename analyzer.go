// public API for game analyzer
package gameanalyzer

import (
	"github.com/Chesstutis/GameAnalyzer/internal/config"
)

type Analyzer struct {
}

// creates a new game analyzer
func NewAnalyzer() (*Analyzer, error) {
	_, err := config.MustStockfishPath()
	if err != nil {
		return nil, err
	}
	return &Analyzer{}, nil
}



// analyzes the game pgn using internal/parser/pgn.go and returns a AnalysisResult
func (a *Analyzer) AnalyzeGame(pgnString string) (*AnalysisResult, error) {
	var allPuzzles []Puzzle
	var allMoveEvals []MoveStat


	result := &AnalysisResult{
		Accuracy: 0,
		Puzzles: allPuzzles,
		MoveStats: allMoveEvals,
	}
	return result, nil
}