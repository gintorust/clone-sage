package detect

import (
	"bufio"
	"os"
	"strings"
)

// looks for the filename
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// returns the env keys from the filename
func extractEnvKeys(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
      return nil
	}

	defer file.Close()

	var keys []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) > 0 {
			keys = append(keys, strings.TrimSpace(parts[0]))
		}
	}

	return keys
}