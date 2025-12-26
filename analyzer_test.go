package gameanalyzer

import (
	"fmt"
	"testing"
)
const (
	scholarsPgn = "1. e4 e5 2. Bc4 Nc6 3. Qh5 Nf6 4. Qxf7# 1-0"
	whiteBlunderQueen = "1. e4 e5 2. Qh5 Nf6 3. Bc4 Nxh5 *"
	chesstutisVsAlex = "1. e4 e5 2. Nf3 Nc6 3. Bc4 Bc5 4. c3 d6 5. d4 Bg4 6. dxc5 dxc5 7. Qxd8+ Rxd8 8. Nbd2 f6 9. h3 Bh5 10. O-O Nge7 11. a3 a6 12. b4 b5 13. Be6 cxb4 14. axb4 Bf7 15. Bxf7+ Kxf7 16. Nb3 Rhe8 17. Be3 Rd3 18. Bd2 Ng6 19. Rad1 Red8 20. Nc5 R3d6 21. Nb7 a5 22. Nxd6+ Rxd6 23. bxa5 Nxa5 24. Rb1 Nc4 25. Rfd1 Nf4 26. Bxf4 exf4 27. Rxd6 cxd6 28. Rxb5 f5 29. exf5 g6 30. fxg6+ hxg6 31. Rb7+ Kf6 32. Nd4 d5 33. Rc7 Nd2 34. Rc6+ Ke5 35. Re6# 1-0"
)

func TestAnalyzeGame_ScholarsMate(t *testing.T) {
	a, err := NewAnalyzer()
	if err != nil {
		t.Fatalf("problem creating analyzer %v", err)
	}
	gameResult, err := a.AnalyzeGame(scholarsPgn)

	if err != nil {
		t.Fatalf("problem analyzing game %v", err)
	}

	fmt.Printf("Puzzles: %v\n\n", gameResult.Puzzles)
	fmt.Printf("MoveStats: %v\n\n", gameResult.MoveStats)
	
}

func TestWhiteBlundersQueen(t *testing.T) {
	a, err := NewAnalyzer()
	if err != nil {
		t.Fatalf("problem creating analyzer %v", err)
	}
	gameResult, err := a.AnalyzeGame(whiteBlunderQueen)

	if err != nil {
		t.Fatalf("problem analyzing game %v", err)
	}
	
	fmt.Printf("Puzzles: %v\n\n", gameResult.Puzzles[0].Position)
	fmt.Printf("MoveStats: %v\n\n", gameResult.MoveStats)
}

func TestAlexBlunders(t *testing.T) {
	a, err := NewAnalyzer()
	if err != nil {
		t.Fatalf("problem creating analyzer %v", err)
	}
	gameResult, err := a.AnalyzeGame(chesstutisVsAlex)

	if err != nil {
		t.Fatalf("problem analyzing game %v", err)
	}
	
	fmt.Printf("Puzzles: %v\n\n", gameResult.Puzzles)
	fmt.Printf("MoveStats: %v\n\n", gameResult.MoveStats)
}