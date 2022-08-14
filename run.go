package main

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

func run(command string) error {
	logrus.Debugf("run: %s", command)
	cmd := exec.Command("/bin/sh", "-c", command)

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func runWithOutput(command string) ([]byte, error) {
	logrus.Debugf("runWithOutput: %s", command)
	cmd := exec.Command("/bin/sh", "-c", command)

	return cmd.Output()
}

func mustRun(command string) {
	logrus.Debugf("mustRun: %s", command)
	cmd := exec.Command("/bin/sh", "-c", command)

	if err := cmd.Run(); err != nil {
		logrus.WithError(err).Panicf("failed to run command: %s", command)
	}
}
