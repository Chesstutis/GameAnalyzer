// public API for game analyzer
package gameanalyzer

import (
	"log"

	"github.com/Chesstutis/GameAnalyzer/internal/config"
	"github.com/Chesstutis/GameAnalyzer/internal/engine"
	"github.com/Chesstutis/GameAnalyzer/internal/logic"
	"github.com/Chesstutis/GameAnalyzer/internal/parser"
	"github.com/Chesstutis/GameAnalyzer/internal/types"
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
func (a *Analyzer) AnalyzeGame(pgnString string) (*types.AnalysisResult, error) {
	var allPuzzles []types.Puzzle
	var allMoveEvals []types.MoveStat // used for game review

	parsedMoves, err := parser.ParsePGN(pgnString)

	if err != nil {
		log.Fatalf("failed to parse PGN %v", err)
	}

	for i, m := range parsedMoves {
		analysis, err := engine.AnalyzePosition(m.FEN, m.Move)
		if err != nil {
			log.Fatalf("error analyzing position %v", err)
		}

		stat := logic.ProcessMove(analysis.BestMove, analysis.PlayerMove, i % 2 == 0)

		allMoveEvals = append(allMoveEvals, *stat)

		if logic.IsPuzzleWorthy(stat.Classification) {
			allPuzzles = append(allPuzzles, types.Puzzle{
				FEN: m.FEN,
				BestMove: analysis.BestMove.Move,
				PlayerMove: m.Move,
				Classification: stat.Classification,
			})
		}
	}


	result := &types.AnalysisResult{
		Accuracy: 0,
		Puzzles: allPuzzles,
		MoveStats: allMoveEvals,
	}
	return result, nil
}