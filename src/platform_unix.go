//go:build !windows

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func knownExecutablePaths(names ...string) []string {
	paths := []string{}
	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		for _, name := range names {
			paths = append(paths, filepath.Join(exeDir, name))
		}
	}
	if home, err := os.UserHomeDir(); err == nil {
		for _, name := range names {
			paths = append(paths, filepath.Join(home, ".local", "bin", name))
		}
	}
	return paths
}

func removeShellPathEntries(rcNames []string, markers []string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	var errs []error
	for _, rcName := range rcNames {
		path := filepath.Join(home, rcName)
		data, err := os.ReadFile(path)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", path, err))
			continue
		}
		lines := strings.Split(string(data), "\n")
		kept := make([]string, 0, len(lines))
		changed := false
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			if line == "# ZPM (Zen Process Manager)" {
				if i+1 < len(lines) && containsAny(lines[i+1], markers) {
					i++
					changed = true
					continue
				}
			}
			if strings.HasPrefix(strings.TrimSpace(line), "export PATH=") && containsAny(line, markers) {
				changed = true
				continue
			}
			kept = append(kept, line)
		}
		if !changed {
			continue
		}
		err = os.WriteFile(path, []byte(strings.Join(kept, "\n")), 0644)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", path, err))
		}
	}
	return errors.Join(errs...)
}

func containsAny(value string, needles []string) bool {
	for _, needle := range needles {
		if needle != "" && strings.Contains(value, needle) {
			return true
		}
	}
	return false
}
