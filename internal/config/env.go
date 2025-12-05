package config

import (
    "fmt"
    "os"
    "path/filepath"
    "runtime"

    "github.com/joho/godotenv"
)

var EnginePath string
// vibe coding is so cooked
func init() {
    // find project root (where .env lives) by walking up from this file
    _, filename, _, _ := runtime.Caller(0)
    dir := filepath.Dir(filename)
    projectRoot := filepath.Join(dir, "..", "..") // go up to project root
    envPath := filepath.Join(projectRoot, ".env")

    // load .env from project root (ignore error if missing)
    _ = godotenv.Load(envPath)

    // read env var
    EnginePath = os.Getenv("ENGINE_PATH")
}

// MustStockfishPath returns the path or an error when not set.
func MustStockfishPath() (string, error) {
    if EnginePath == "" {
        return "", fmt.Errorf("ENGINE_PATH not set in .env or environment")
    }
    return EnginePath, nil
}