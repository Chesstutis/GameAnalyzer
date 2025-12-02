package gameanalyzer

import (
	"fmt"

	"github.com/notnil/chess/uci"
)

func newUCIEngine(stockfishPath string) (*uci.Engine, error) {
	eng, err := uci.New(stockfishPath)
	if err != nil {
		panic(err)
	}
	return eng, nil
}
