package game

import (
	. "tictactoeweb/internal"
	. "tictactoeweb/internal/domain/game"
)

type CliKey Key

// "3" IsIn {"1", "3", "5"}
func (k CliKey) IsIn(kk []CliKey) bool {
	for _, v := range kk {
		if v == k {
			return Yes
		}
	}
	return No
}
