package types

import (
	"github.com/notnil/chess"
)
type ParsedMove struct {
    Position  chess.Position
    Move      chess.Move
}

type EvaluatedMove struct {
    Move       chess.Move
    Evaluation int
}

type MoveStat struct {
    Move           chess.Move
    CentipawnLoss  int
    Classification MoveClass
}

type EngineAnalysis struct {
    BestMove   EvaluatedMove
    PlayerMove EvaluatedMove
}

type Puzzle struct {
    Position        chess.Position
    BestMove        chess.Move
    PlayerMove      chess.Move
    Classification  MoveClass
}

type AnalysisResult struct {
    Accuracy  float64
    Puzzles   []Puzzle
    MoveStats []MoveStat
}

type MoveClass string

const (
	Best       MoveClass = "Best"
	Excellent  MoveClass = "Excellent"
	Good       MoveClass = "Good"
	Innacuracy MoveClass = "Innacuracy"
	Mistake    MoveClass = "Mistake"
	Blunder    MoveClass = "Blunder"
	Miss       MoveClass = "Miss"
)