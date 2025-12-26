// take raw data from the engine and outputs meaningful results
// ex: move classifications, should a position be a puzzle, overall accuracy
package logic

import (
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

	classification := ClassifyMove(cpl)

	return &types.MoveStat{
		Move:           playerMove.Move,
		CentipawnLoss:  cpl,
		Classification: classification,
	}
}

func ClassifyMove(cpl int) types.MoveClass {
	switch {
	case cpl <= 15: return  types.Best
	case cpl <= 50: return  types.Excellent
	case cpl <= 100: return types.Innacuracy
	case cpl <= 250: return types.Mistake
	default: return types.Blunder
	}
}

func CalculateAccuracy() int {
	return 0
}

func IsPuzzleWorthy(classification types.MoveClass) bool {
	return classification == types.Mistake || 
	       classification == types.Blunder || 
		   classification == types.Miss
}
