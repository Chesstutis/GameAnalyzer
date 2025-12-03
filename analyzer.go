// public API for game analyzer
package gameanalyzer

type Analyzer struct { 
	// path to stockfish executable
	stockfishPath string
}

// creates a new game analyzer
func NewAnalyzer(stockfishPath string) (*Analyzer, error) {
	// check for stockfish path then 
	return nil, nil
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