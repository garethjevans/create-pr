package cmd

import (
	"errors"
	"os/exec"
	"strings"
)

var _ Gitter = (*DefaultGitter)(nil)
var Runner = ShellCommand

type Gitter interface {
	GetOrigin() (string, error)
	HasLocalChanges() (bool, error)
}

type DefaultGitter struct {
}

func (d DefaultGitter) GetOrigin() (string, error) {
	remotes, err := Runner("git", "remote", "-v")
	if err != nil {
		return "", err
	}
	lines := strings.Split(remotes, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "origin") && strings.HasSuffix(line, "(fetch)") {
			return strings.TrimSuffix(strings.TrimSpace(strings.TrimPrefix(strings.TrimSuffix(line, "(fetch)"), "origin")), ".git"), nil
		}
	}
	return "", errors.New("unable to find origin")
}

func (d DefaultGitter) HasLocalChanges() (bool, error) {
	status, err := Runner("git", "status", "--porcelain")
	if err != nil {
		return false, err
	}
	return status != "", nil
}

func ShellCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}
