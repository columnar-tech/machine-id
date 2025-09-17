package machineid

import (
	"fmt"
	"os/exec"
	"strings"
)

// machineID returns the uuid returned by `ioreg -rd1 -c IOPlatformExpertDevice`.
// If there is an error running the commad an empty string is returned.
func machineID() (string, error) {
	buf, err := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()
	if err != nil {
		return "", err
	}
	id, err := extractID(string(buf))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(id), nil
}

func extractID(lines string) (string, error) {
	for _, line := range strings.Split(lines, "\n") {
		if strings.Contains(line, "IOPlatformUUID") {
			parts := strings.SplitAfter(line, `" = "`)
			if len(parts) == 2 {
				return strings.TrimRight(parts[1], `"`), nil
			}
		}
	}
	return "", fmt.Errorf("failed to extract 'IOPlatformUUID' value from `ioreg` output: %s", lines)
}
