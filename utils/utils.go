package utils

import (
	"strconv"
	"strings"
)

func ValidNetwork(n string) bool {
	n = strings.TrimSpace(n)
	return (n == "tcp4") || (n == "udp4") || (n == "tcp6") || (n == "udp6")
}

func ValidPort(p string) bool {
	portNumber, err := strconv.Atoi(strings.TrimSpace(p))
	if err != nil {
		return false
	}
	if (portNumber >= 0) && (portNumber <= 65535) {
		return true
	} else {
		return false
	}
}
