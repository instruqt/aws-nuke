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
}

func (p *NukeParameters) Validate() error {
	if strings.TrimSpace(p.ConfigPath) == "" {
		return fmt.Errorf("You have to specify the --config flag.\n")
	}
	if p.ForceSleep < 3 && p.NoDryRun {
		return fmt.Errorf("Value for --force-sleep cannot be less than 3 seconds if --no-dry-run is set. This is for your own protection.\n")
	}

	return nil
}
