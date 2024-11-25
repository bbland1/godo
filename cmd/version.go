package cmd

import (
	"flag"
	"fmt"
	"io"
)

const VersionUsage = `print the app version and build info of what is currently install
usage:
	goDo version [options] 

options:
	-verbose	default: false. if true, print just the version info`

type VersionInfo struct {
	Build string
	Version string
	Verbose bool
}

var buildInfo = VersionInfo{Build: "blank", Version: "blank", Verbose: false}

func versionPrintFunc(w io.Writer) {
	if buildInfo.Verbose {
		fmt.Fprintf(w, "goDo v%s, build: %s", buildInfo.Version, buildInfo.Build)
		return
	}
	
	fmt.Fprintf(w, "goDo v%s", buildInfo.Version)
}

func NewVersionCommand(w io.Writer) *Command {
	command := &Command{
		flags: flag.NewFlagSet("version", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			versionPrintFunc(w)
		},
	}

	command.flags.Usage = func() {
		fmt.Fprintln(w, VersionUsage)
	}

	return command
}