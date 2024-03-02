package types

import (
	"os"
	"path/filepath"
)

var (
	Home, _  = os.UserHomeDir()
	PathDir  = filepath.Join(Home, ".config", "upgit")
	PathFile = filepath.Join(Home, ".config", "upgit", "paths.json")
)
