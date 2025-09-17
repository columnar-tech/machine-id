//go:build freebsd || netbsd || openbsd || dragonfly || solaris

package machineid

import (
	"os"
	"os/exec"
	"strings"
)

const hostidPath = "/etc/hostid"

// machineID returns the uuid specified at `/etc/hostid`.
// If the returned value is empty, the uuid from a call to `kenv -q smbios.system.uuid` is returned.
// If there is an error an empty string is returned.
func machineID() (string, error) {
	id, err := readHostid()
	if err != nil {
		// try fallback
		id, err = readKenv()
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

func readHostid() (string, error) {
	buf, err := os.ReadFile(hostidPath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buf)), nil
}

func readKenv() (string, error) {
	buf, err := exec.Command("kenv", "-q", "smbios.system.uuid").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buf)), nil
}
