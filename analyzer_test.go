package gameanalyzer

import (
	"fmt"
	"log"
	"testing"
)
const (
	scholarsPgn = "1. e4 e5 2. Bc4 Nc6 3. Qh5 Nf6 4. Qxf7# 1-0"
)

func TestAnalyzeGame_ScholarsMate(t *testing.T) {
	a, err := NewAnalyzer()
	if err != nil {
		log.Fatalf("problem creating analyzer %v", err)
	}
	gameResult, err := a.AnalyzeGame(scholarsPgn)

	if err != nil {
		log.Fatalf("problem analyzing game %v", err)
	}

	fmt.Printf("Puzzles: %v\n\n", gameResult.Puzzles)
	fmt.Printf("MoveStats: %v\n\n", gameResult.MoveStats)
	
}