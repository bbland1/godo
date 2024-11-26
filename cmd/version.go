package cmd

import (
	"flag"
	"fmt"
	"io"
)

const VersionUsage = `print the app version what is currently install
usage:
	goDo version [options] 

options:
	-verbose	default: false. if true, print just the version and build info`

type VersionInfo struct {
	Build string
	Version string
	Verbose bool
}

var BuildInfo = VersionInfo{Build: "blank", Version: "blank", Verbose: false}

func versionPrintFunc(w io.Writer) {
	if BuildInfo.Verbose {
		fmt.Fprintf(w, "goDo v%s, build: %s\n", BuildInfo.Version, BuildInfo.Build)
		return
	}
	
	fmt.Fprintf(w, "goDo v%s\n", BuildInfo.Version)
}

func NewVersionCommand(w io.Writer) *Command {
	command := &Command{
		flags: flag.NewFlagSet("version", flag.ExitOnError),
		Execute: func(cmd *Command, args []string) {
			versionPrintFunc(w)
		},
	}

	command.flags.BoolVar(&BuildInfo.Verbose, "verbose", false, "print out the full version/build info")

	command.flags.Usage = func() {
		fmt.Fprintln(w, VersionUsage)
	}

	return command
}