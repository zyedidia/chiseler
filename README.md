# Chiseler

A tool for quickly setting up Chisel projects. In the future this tool might
also include commands for managing Chisel projects.

# Install

```
go install github.com/zyedidia/chiseler@latest
```

or install a prebuilt binary from the releases page.

```
eget zyedidia/chiseler
```

# Usage

Create a new project by running `chiseler init`. This will set up the directory
structure with a blank Chisel project.


```
Usage: chiseler <flags> <subcommand> <subcommand args>

Subcommands:
	help             describe subcommands and their syntax
	init             initialize a Chisel project

init [OPTS]:
  Create a new Chisel project.
  -chisel string
    	Chisel version (default "3.5.1")
  -chiselTest string
    	ChiselTest version (default "0.5.1")
  -package string
    	Package name (default "main")
  -scala string
    	Scala version (default "2.12.13")
  -top string
    	Top module name (default "Top")
```
