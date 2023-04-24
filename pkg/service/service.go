package service

import (
	"sync"
)

var (
	Mu    sync.Mutex
	Codes = make(map[string]int)
)
