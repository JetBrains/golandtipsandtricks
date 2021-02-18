package main

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {
	// Step 1. Observe the warning in Errorf missing %s
	err := errors.Errorf("Hello!", "World")

	// Step 2. Observe the warning in Infof missing %s
	logrus.WithError(err).Infof("Hello!", "World")

	sugar := zap.NewExample().Sugar()
	defer func() { _ = sugar.Sync() }()

	// Step 3. Observe the warning in Infof missing %s
	sugar.Infof("Hello!", "World")
}
