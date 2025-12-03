// take raw data from the engine and outputs meaningful results
// ex: move classifications, should a position be a puzzle, overall accuracy
package logic

import gameanalyzer "github.com/Chesstutis/GameAnalyzer"

// processes user move by turning it into a MoveStat
func ProcessMove(bestMove gameanalyzer.EvaluatedMove, playerMove gameanalyzer.EvaluatedMove) *gameanalyzer.MoveStat {
	cpl := bestMove.Evaluation - playerMove.Evaluation

	classification := ClassifyMove(cpl)

	return &gameanalyzer.MoveStat{
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

func isPuzzleWorthy(classification string) bool {                                        // Miss - not yet implemented
	return classification == "Mistake" || classification == "Blunder" || classification == "Miss"
}
