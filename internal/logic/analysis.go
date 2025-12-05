// take raw data from the engine and outputs meaningful results
// ex: move classifications, should a position be a puzzle, overall accuracy
package logic

import (
	"fmt"

	"github.com/Chesstutis/GameAnalyzer/internal/types"
)

// processes user move by turning it into a MoveStat
func ProcessMove(bestMove types.EvaluatedMove, playerMove types.EvaluatedMove, isWhiteToMove bool) *types.MoveStat {
	cpl := bestMove.Evaluation - playerMove.Evaluation
	if isWhiteToMove {
		cpl = max(bestMove.Evaluation - playerMove.Evaluation, 0)
	} else {
		cpl = playerMove.Evaluation - bestMove.Evaluation
        if cpl < 0 {
            cpl = -cpl // make it positive
        }
	}
	fmt.Printf("CPL: %v\n\n", cpl)

	classification := ClassifyMove(cpl)

	return &types.MoveStat{
		Move:           playerMove.Move,
		CentipawnLoss:  cpl,
		Classification: classification,
	}
}

func ClassifyMove(cpl int) string {
	switch {
	case cpl <= 15: return "Best"
	case cpl <= 50: return "Excellent"
	case cpl <= 100: return "Inaccuracy"
	case cpl <= 250: return "Mistake"
	default: return "Blunder"
	}
}

func CalculateAccuracy() int {
	return 0
}

func IsPuzzleWorthy(classification string) bool {                                    //TODO Miss not yet implemented
	return classification == "Mistake" || classification == "Blunder" || classification == "Miss"
}
