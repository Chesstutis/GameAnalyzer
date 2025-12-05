package parser

import (
	"testing"
)

func TestParseSimpleGame(t *testing.T) {
	pgn := `1. e4 e5 2. Nf3 Nc6 3. Bb5 a6`

	moves, err := Parse(pgn)
	if err == nil {
		t.Fatalf("expected not implemented error, got nil")
	}

	if len(moves) == 0 {
		t.Fatalf("expected moves, got 0")
	}

	// Check first move
	first := moves[0]
	if first.Move != "e2e4" {
		t.Errorf("expected first move 'e4', got '%s'", first.Move)
	}
	if first.FEN == "" {
		t.Errorf("expected FEN for first move, got empty string")
	}

	// Check second move
	second := moves[1]
	if second.Move != "e7e5" {
		t.Errorf("expected second move 'e5', got '%s'", second.Move)
	}

	thrid := moves[2]
	if thrid.Move !=  "g1f3"{
		t.Errorf("expected third move 'g1f3', got '%s'", thrid.Move)

	}

	fourth := moves[3]
	if fourth.Move != "b8c6"{
		t.Errorf("expected third move 'b8c6', got '%s'", fourth.Move)
	}


	
}


