package domain

import "strconv"

type Key string

func (k Key) IsKey() bool {
	n, err := strconv.Atoi(string(k))
	if err != nil {
		return false
	}
	return n >= 1 && n <= 9
}

func (k Key) ToCell() Cell {
	return keyToCell()[k]
}
