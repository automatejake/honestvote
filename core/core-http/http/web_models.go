package http

import "github.com/jneubaum/honestvote/core/core-database/database"

type ElectionInfo struct {
}

func Translate(e database.Election) ElectionInfo {
	return ElectionInfo{}
}
