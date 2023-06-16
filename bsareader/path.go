package bsareader

import (
	"os"
	"path/filepath"
)

// GetDaggerfallPath
// IDK if it's cool for me to redistribute BSA files with this repo
// so this is what we're going to do instead.
func GetDaggerfallPath() string {
	path := os.Getenv("DAGGERFALL_PATH")
	if path != "" {
		return path
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "/home/rowan/abandon/dfall/arena2"
	}
	return filepath.Join(home, "arena2")
}
