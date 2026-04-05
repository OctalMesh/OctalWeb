// Package version holds the build-time version information for the ow CLI.
// The Version variable is overridden at link time via -ldflags.
//
// Example build:
//
//	go build -ldflags="-X github.com/OctalMesh/OctalWeb/cli/pkg/version.Version=1.2.3" .
package version

// Version is the canonical SemVer string for this binary.
// Overridden at build time; defaults to "dev" for local builds.
var Version = "dev"
