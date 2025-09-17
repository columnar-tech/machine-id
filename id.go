// Package machineid provides a way to get a unique identifier for the current machine
// without using any internal hardware IDs. As a result, it doesn't require any special
// permissions to run and works in unprivileged containers as well.
//
// The returned ID is generally stable for the OS installation and usually stays the same
// after updates or hardware changes.
//
// Caveat: Image-based environments usually have the same machine-id (perfect clones).
// Linux users can generate a new id with `dbus-uuidgen` and put the id into
// `/var/lib/dbus/machine-id` and `/etc/machine-id`.
//
// Windows users can use the `sysprep` toolchain to create images, which produce valid images
// ready for distribution.
package machineid

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// ID returns the platform specific machine id of the current host OS.
// Consider the returned id as "confidential" information and consider using
// ProtectedID() instead for privacy preserving usage.
func ID() (string, error) {
	id, err := machineID()
	if err != nil {
		return "", fmt.Errorf("machineid: %w", err)
	}
	return id, nil
}

// ProtectedID returns an HMAC-SHA256 hashed version of the machine ID using a fixed key.
// This is a privacy preserving way to use the machine ID as it is not reversible
// and the same on every call on the same machine.
func ProtectedID() (string, error) {
	id, err := ID()
	if err != nil {
		return "", err
	}

	h := hmac.New(sha256.New, []byte("machine-id"))
	if _, err = h.Write([]byte(id)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
