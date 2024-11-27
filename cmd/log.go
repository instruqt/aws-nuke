package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/instruqt/aws-nuke/v2/resources"
	"github.com/sirupsen/logrus"
)

var (
	ReasonSkip            = *color.New(color.FgYellow)
	ReasonError           = *color.New(color.FgRed)
	ReasonRemoveTriggered = *color.New(color.FgGreen)
	ReasonWaitPending     = *color.New(color.FgBlue)
	ReasonSuccess         = *color.New(color.FgGreen)
)

var (
	ColorRegion             = *color.New(color.Bold)
	ColorResourceType       = *color.New()
	ColorResourceID         = *color.New(color.Bold)
	ColorResourceProperties = *color.New(color.Italic)
)

// Format the resource properties in sorted order ready for printing.
// This ensures that multiple runs of aws-nuke produce stable output so
// that they can be compared with each other.
func Sorted(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sorted := make([]string, 0, len(m))
	for k := range keys {
		sorted = append(sorted, fmt.Sprintf("%s: '%s'", keys[k], m[keys[k]]))
	}
	return fmt.Sprintf("[%s]", strings.Join(sorted, ", "))
}

func Log(region *Region, resourceType string, r resources.Resource, c color.Color, msg string) {
	logMsg := fmt.Sprintf("%s - %s - ", region.Name, resourceType)

	rString, ok := r.(resources.LegacyStringer)
	if ok {
		logMsg += fmt.Sprintf("%s -", rString.String())
	}

	rProp, ok := r.(resources.ResourcePropertyGetter)
	if ok {
		logMsg += fmt.Sprintf("%s -", Sorted(rProp.Properties()))
	}

	logMsg += fmt.Sprintf("%s", msg)

	if c.Equals(&ReasonSuccess) || c.Equals(&ReasonWaitPending) {
		logrus.Info(logMsg)
	} else if c.Equals(&ReasonSkip) {
		logrus.Warn(logMsg)
	} else if c.Equals(&ReasonError) {
		logrus.Error(logMsg)
	}
}
