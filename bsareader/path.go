package bsareader

import "os"

// GetDaggerfallPath
// IDK if it's cool for me to redistribute BSA files with this repo
// so this is what we're going to do instead.
func GetDaggerfallPath() string {
	path := os.Getenv("DAGGERFALL_PATH")
	if path != "" {
		return path
	}
	return "/home/rowan/games/abandon/dfall/arena2"
}
