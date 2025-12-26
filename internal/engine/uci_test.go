package engine

import (
	"fmt"
	"os"
	"testing"

	"github.com/Chesstutis/GameAnalyzer/internal/config"
	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)

func TestMain(m *testing.M) {
	_, err := config.MustStockfishPath()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Using stockfish at:", config.EnginePath)
	os.Exit(m.Run())
}

func TestParseMove_Found(t *testing.T) {
	game := chess.NewGame()
	validMoves := game.Position().ValidMoves()
	if len(validMoves) == 0 {
		t.Fatal("no valid moves")
	}

	m, err := parseMove(game.Position(), *validMoves[0])
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Move Found: %v", m)
}

func TestCheckSolvedPos(t *testing.T) {
	t.Run("mateReported", func(t *testing.T) {
		var res uci.SearchResults
		res.Info.Score.Mate = 2

		game := chess.NewGame()
		validMoves := game.Position().ValidMoves()
		if len(validMoves) == 0 {
			t.Fatal("no valid moves")
		}

		em := CheckSolvedPos(res, *validMoves[0])
		if em == nil {
			t.Fatal("expected non-nil EvaluatedMove for mate")
		}
		if em.Evaluation != 1000000 {
			t.Fatalf("expected Evaluation=1000000 for mate; got %d", em.Evaluation)
		}
	})

	t.Run("centipawnScore", func(t *testing.T) {
		var res uci.SearchResults
		res.Info.Score.CP = 123

		game := chess.NewGame()
		validMoves := game.Position().ValidMoves()
		if len(validMoves) == 0 {
			t.Fatal("no valid moves")
		}

		em := CheckSolvedPos(res, *validMoves[0])
		if em == nil {
			t.Fatal("expected non-nil EvaluatedMove for cp score")
		}
		if em.Evaluation != 123 {
			t.Fatalf("expected Evaluation=123; got %d", em.Evaluation)
		}
	})
}