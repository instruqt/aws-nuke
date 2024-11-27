package cmd

import (
	"fmt"
	"strings"
)

type NukeParameters struct {
	ConfigPath string

	Targets      []string
	Excludes     []string
	CloudControl []string

	NoDryRun   bool
	Force      bool
	ForceSleep int
	Quiet      bool

	MaxWaitRetries int

	Metadata map[string]any
}

func (p *NukeParameters) Validate() error {
	if strings.TrimSpace(p.ConfigPath) == "" {
		return fmt.Errorf("You have to specify the --config flag.")
	}

	return nil
}
