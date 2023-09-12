package virtualbox

import (
	"context"
	"fmt"
)

func (m *Manager) Version(ctx context.Context) (string, error) {
	m.log.Println("get version")
	stdout, _, err := m.run(ctx, "--version")
	if err != nil {
		return "", fmt.Errorf("unable to get version: %w", err)
	}
	return stdout, nil
}
