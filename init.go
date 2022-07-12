package main

import (
	"context"
	"embed"
	"flag"

	"github.com/google/subcommands"
)

//go:embed template
var template embed.FS

type initCmd struct {
	name    string
	version string
}

func (*initCmd) Name() string     { return "init" }
func (*initCmd) Synopsis() string { return "initialize a Chisel project" }
func (*initCmd) Usage() string {
	return `init [OPTS] [NAME]:
  Create a new Chisel project.
`
}

func (p *initCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.outname, "v", "", "Chisel version")
}

func (p *initCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
}
