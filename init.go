package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"strings"

	"github.com/google/subcommands"
)

//go:embed template
var templateFS embed.FS

type initCmd struct {
	Top               string
	Package           string
	ChiselVersion     string
	ScalaVersion      string
	ChiselTestVersion string
}

func (*initCmd) Name() string     { return "init" }
func (*initCmd) Synopsis() string { return "initialize a Chisel project" }
func (*initCmd) Usage() string {
	return `init [OPTS]:
  Create a new Chisel project.
`
}

func (p *initCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.ChiselVersion, "chisel", "3.5.1", "Chisel version")
	f.StringVar(&p.ScalaVersion, "scala", "2.12.13", "Scala version")
	f.StringVar(&p.ChiselTestVersion, "chiselTest", "0.5.1", "ChiselTest version")
	f.StringVar(&p.Top, "top", "Top", "Top module name")
	f.StringVar(&p.Package, "package", "main", "Package name")
}

func (p *initCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	err := fs.WalkDir(templateFS, "template", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasPrefix(path, "template/") {
			return nil
		}
		newpath := path[len("template/"):]

		if strings.Contains(newpath, "Top.scala") {
			newpath = strings.Replace(newpath, "Top.scala", p.Top+".scala", -1)
		}

		fmt.Println(newpath)

		if d.IsDir() {
			return os.MkdirAll(newpath, os.ModePerm)
		}

		data, err := fs.ReadFile(templateFS, path)
		if err != nil {
			return err
		}
		t, err := template.New("chisel").Parse(string(data))
		if err != nil {
			return err
		}
		f, err := os.Create(newpath)
		if err != nil {
			return err
		}
		err = t.Execute(f, p)
		if err != nil {
			return err
		}
		return f.Close()
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
