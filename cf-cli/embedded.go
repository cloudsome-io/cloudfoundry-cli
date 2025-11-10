// internal/cfcli/embedded.go
package cfcli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type embeddedBinary struct {
	data []byte
	name string
}

var platformBinary embeddedBinary

type EmbeddedCLI struct {
	binaryPath string
}

func New() (*EmbeddedCLI, error) {
	if len(platformBinary.data) == 0 {
		return nil, fmt.Errorf("embedded CF CLI binary not available for this platform")
	}

	// Create a cache directory for the CLI binary
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, fmt.Errorf("cannot get cache dir: %w", err)
	}

	cliDir := filepath.Join(cacheDir, "your-cf-cli")
	if err := os.MkdirAll(cliDir, 0755); err != nil {
		return nil, fmt.Errorf("cannot create cli dir: %w", err)
	}

	binaryPath := filepath.Join(cliDir, platformBinary.name)

	// Write the binary if it is missing or outdated
	if needsUpdate(binaryPath, platformBinary.data) {
		if err := os.WriteFile(binaryPath, platformBinary.data, 0755); err != nil {
			return nil, fmt.Errorf("cannot write cf binary: %w", err)
		}
	}

	return &EmbeddedCLI{binaryPath: binaryPath}, nil
}

func needsUpdate(path string, data []byte) bool {
	existing, err := os.ReadFile(path)
	if err != nil {
		return true
	}
	// Compare size (or checksum in the future) to detect changes
	return len(existing) != len(data)
}

func (e *EmbeddedCLI) Execute(args ...string) error {
	cmd := exec.Command(e.binaryPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func (e *EmbeddedCLI) ExecuteWithOutput(args ...string) (string, error) {
	cmd := exec.Command(e.binaryPath, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
