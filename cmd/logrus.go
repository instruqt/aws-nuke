package cmd

import (
	"github.com/sirupsen/logrus"
)

// metadataHook implementation to add default fields
type metadataHook struct {
	fields logrus.Fields
}

func (h *metadataHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *metadataHook) Fire(entry *logrus.Entry) error {
	for key, value := range h.fields {
		entry.Data[key] = value
	}

	return nil
}
