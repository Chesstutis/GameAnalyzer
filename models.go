package gameanalyzer

// ---
// Public API for Analysis Results
// ---

// AnalysisResult is the top-level struct returned after a full game analysis.
type AnalysisResult struct {
	// accuracy of overall game 0-100
	//! wont need this yet for out project
	Accuracy float64 `json:"accuracy"`

	// list of puzzles generated from games blunders/misses/mistakes
	Puzzles []Puzzle `json:"puzzles"`

	// A list of stats for every move in the game.
	// This is the "full report".
	MoveStats []MoveStat `json:"moveStats"`
}

// Puzzle represents a single, solvable puzzle position.
type Puzzle struct {
	// starting fen string of the puzzle (before the blunder was made)
	FEN string `json:"fen"`

	// incorrect move that the player made
	BlunderMove string `json:"blunderMove"`

	// best move (solution to the puzzle)
	BestMove string `json:"bestMove"`
}

// MoveStats holds the analysis for a single move. (This was your 'MoveEvaluation')
// Your AnalysisResult will contain one of these for each move in the game.
type MoveStat struct {
	// move the player made in long algebraic notation
	// ex: e2e4
	Move string `json:"move"`

	// difference, in centipawns, between the best move and the players move
	// 0 is best
	CentipawnLoss int `json:"centipawnLoss"`

	// move classification
	// ex: "Best", "Blunder"....
	Classification string `json:"classification"`
}

// ---
// Internal Helper Structs
// ---
// These are structs used to pass data between your internal packages.
// ---

// EvaluatedMove holds a single move and its score from the engine.
// This is the raw data from Stockfish.
type EvaluatedMove struct {
	// move in long algebraic notation
	// ex: e2e4
	Move string

	// engines score in centipawns
	Evaluation int
}

