package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseEnv() error {
	if hostEnv := os.Getenv("host"); hostEnv != "" {
		host = hostEnv
	}

	if portEnv := os.Getenv("port"); portEnv != "" {
		value, err := strconv.Atoi(portEnv)
		if err != nil {
			return fmt.Errorf("unable to parse port: %w", err)
		}
		port = uint16(value)
	}
	return nil
}
