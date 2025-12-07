package engine

import (
	"fmt"
	"time"

	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"

	"github.com/Chesstutis/GameAnalyzer/internal/config"
	"github.com/Chesstutis/GameAnalyzer/internal/types"
)

// finds the best move in the position as well as the analysis for the players move
func AnalyzePosition(fen string, playerMove string) (_ *types.EngineAnalysis, err error) {
	// make new UCI engine
	eng, err := uci.New(config.EnginePath)
	if err != nil {
		err = fmt.Errorf("failed to start engine: %w", err)
		return
	}
	defer eng.Close()

	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}

	// make new chess game
	pos, _ := chess.FEN(fen)
	game := chess.NewGame(pos)

	// ------------------------------------
	// --- 1. ---  Analyze Engine Move ----
	// ------------------------------------
	cmdPos := uci.CmdPosition{Position: game.Position()}
	cmdGoBest := uci.CmdGo{MoveTime: time.Second / 10} //! probably should change later

	if err = eng.Run(cmdPos, cmdGoBest); err != nil {
		return nil, fmt.Errorf("engine run failed: %w", err)
	}
	bestSearchResults := eng.SearchResults()

	bestMoveEval := CheckSolvedPos(bestSearchResults, bestSearchResults.BestMove.String())
	
	// ------------------------------------
	// --- 2. --- Analyze Player Move -----
	// ------------------------------------
	playerMoveObj, err := parseMove(game.Position(), playerMove)
	if err != nil {
		return nil, fmt.Errorf("invalid player move: %w", err)
	}

	cmdGoPlayer := uci.CmdGo{
		MoveTime: time.Second,
		SearchMoves: []*chess.Move{playerMoveObj}, // forces eval on player move
	}

	if err := eng.Run(cmdPos, cmdGoPlayer); err != nil {
		return nil, fmt.Errorf("player move analysis failed: %w", err)
	} 
	playerMoveEval := CheckSolvedPos(eng.SearchResults(), playerMove)

	return &types.EngineAnalysis{
		BestMove: *bestMoveEval,
		PlayerMove: *playerMoveEval,
	}, nil
}


// finds the chess.Move object corresponding to the players move
func parseMove(pos *chess.Position, moveStr string) (*chess.Move, error) {
	validMoves := pos.ValidMoves()
	for _, m := range validMoves {
		if m.String() == moveStr {
			return m, nil
		}
	}
	return nil, fmt.Errorf("no valid move matching '%s' found in FEN: %s", moveStr, pos.Board().Draw())
}


func CheckSolvedPos(results uci.SearchResults, move string) *types.EvaluatedMove {
	
	if results.Info.Score.Mate != 0 { // position is solved (forced mate/draw)
		return &types.EvaluatedMove{
			Move: move,
			Evaluation: 1000000,
		}
	} else {
		return &types.EvaluatedMove{
			Move: move,
			Evaluation: results.Info.Score.CP,
		}
	}
}