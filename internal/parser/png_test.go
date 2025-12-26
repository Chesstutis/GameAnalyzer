package parser

import (
	"testing"
)

func TestParseSimpleGame(t *testing.T) {
	pgn := `1. e4 e5 2. Nf3 Nc6 3. Bb5 a6`

	moves, err := ParsePGN(pgn)
	if err != nil {
		t.Fatalf("error-jamin")
	}

	if len(moves) == 0 {
		t.Fatalf("expected moves, got 0")
	}

	// Check first move
	first := moves[0]
	if first.Move.String() != "e2e4" {
		t.Errorf("expected first move 'e4', got '%s'", first.Move.String())
	}
	if first.Position.String() == "" {
		t.Errorf("expected FEN for first move, got empty string")
	}

	// Check second move
	second := moves[1]
	if second.Move.String() != "e7e5" {
		t.Errorf("expected second move 'e5', got '%s'", second.Move.String())
	}

	thrid := moves[2]
	if thrid.Move.String() !=  "g1f3" {
		t.Errorf("expected third move 'g1f3', got '%s'", thrid.Move.String())

	}

	fourth := moves[3]
	if fourth.Move.String() != "b8c6"{
		t.Errorf("expected third move 'b8c6', got '%s'", fourth.Move.String())
	}
}

func TestParsePGN_Empty(t *testing.T) {
    pgn := ""

    moves, err := ParsePGN(pgn)
    
    // empty pgn should return an error (not implemented)
    if err == nil {
        t.Fatalf("expected error for empty pgn, got nil")
    }
    
    // empty pgn should return no moves
    if len(moves) > 0 {
        t.Fatalf("expected 0 moves for empty pgn, got %d", len(moves))
    }
}

