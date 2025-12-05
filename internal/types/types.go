package types

type ParsedMove struct {
    FEN  string
    Move string
}

type EvaluatedMove struct {
    Move       string
    Evaluation int
}

type MoveStat struct {
    Move           string
    CentipawnLoss  int
    Classification string
}

type EngineAnalysis struct {
    BestMove   EvaluatedMove
    PlayerMove EvaluatedMove
}

type Puzzle struct {
    FEN             string
    BestMove        string
    PlayerMove      string
    Classification  string
}

type AnalysisResult struct {
    Accuracy  float64
    Puzzles   []Puzzle
    MoveStats []MoveStat
}