package game

import (
	. "tictactoeweb/internal"
	"tictactoeweb/internal/domain"
)

type Key domain.Key

// "3" IsIn {"1", "3", "5"}
func (k Key) IsIn(kk []Key) bool {
	for _, v := range kk {
		if v == k {
			return Yes
		}
	}
	return No
}
