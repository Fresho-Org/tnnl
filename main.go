package main

import (
	_ "embed"
	"strings"

	"github.com/fresho-org/tnnl/cmd"
	_ "github.com/fresho-org/tnnl/cmd/exec"
	_ "github.com/fresho-org/tnnl/cmd/portforward"
	_ "github.com/fresho-org/tnnl/cmd/remoteportforward"
	_ "github.com/fresho-org/tnnl/cmd/update"
)

//go:embed .version
var version string

func main() {
	cmd.Version = strings.TrimSpace(version)
	cmd.Execute()
}
