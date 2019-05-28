package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func inputOut(logDir, commandName string) (stdout, stderr io.Writer, err error) {
	if logDir == "" {
		stdout = os.Stdout
		stderr = os.Stderr
	} else {
		ts := time.Now().Unix()

		stdoutFileName := fmt.Sprintf("%s-%v-status.log", commandName, ts)
		stdoutFile, err := os.Create(filepath.Join(logDir, stdoutFileName))
		if err != nil {
			return nil, nil, err
		}
		stdout = io.MultiWriter(os.Stdout, stdoutFile)

		stderrFileName := fmt.Sprintf("%s-%v-error.log", commandName, ts)
		stderrFile, err := os.Create(filepath.Join(logDir, stderrFileName))
		if err != nil {
			return nil, nil, err
		}
		stderr = io.MultiWriter(os.Stderr, stderrFile)
	}
	return
}
