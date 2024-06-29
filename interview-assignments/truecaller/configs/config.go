package configs

import (
	_ "embed"
)

var (
	//go:embed data/prefixes.txt
	Data []byte
)
