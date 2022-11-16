// Package core provides an entry point to use Aiko core functionalities.
//
// Aiko makes it possible to accept incoming network connections with certain
// protocol, process the data, and send them through another connection with
// the same or a difference protocol on demand.
//
// It may be configured to work with multiple protocols at the same time, and
// uses the internal router to tunnel through different inbound and outbound
// connections.
package core

//go:generate go run github.com/Github-Aiko/Aiko-Core/common/errors/errorgen

import (
	"runtime"

	"github.com/Github-Aiko/Aiko-Core/common/serial"
)

var (
	version  = "1.6.6"
	build    = "Custom"
	codename = "Aiko, Penetrates Everything."
	intro    = "A unified platform for anti-censorship."
)

// Version returns Aiko's version as a string, in the form of "x.y.z" where x, y and z are numbers.
// ".z" part may be omitted in regular releases.
func Version() string {
	return version
}

// VersionStatement returns a list of strings representing the full version info.
func VersionStatement() []string {
	return []string{
		serial.Concat("Aiko ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")"),
		intro,
	}
}
