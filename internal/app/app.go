package app

import (
	"fmt"

	"github.com/notmithun/mdir/internal/filesystem"
	"github.com/notmithun/mdir/internal/formatter"
)

const Version = "v1.5"

type App struct {
	Scanner *filesystem.Scanner
	All     bool
}

func New() *App {
	return &App{
		Scanner: filesystem.NewScanner(),
	}
}

func (a *App) Run(args []string) {
	switch {
	case len(args) == 0:
		a.list(".")
	case len(args) == 1:
		switch args[0] {
		case "-help":
			a.help()
		case "-v", "-ver", "-version":
			a.version()
		case "-a", "-all":
			a.All = true
			a.list(".")
		case "-r", "--recursive":
			a.listRecursive(".")
		default:
			a.list(args[0])
		}
	default:
		fmt.Println("mdir: too many arguments")
	}
}

func (a *App) list(path string) {
	options := filesystem.ScanOptions{
		All: a.All,
	}

	entries, err := a.Scanner.Scan(path, options)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	formatter.PrintEntries(entries)
}

func (a *App) listRecursive(path string) {
	options := filesystem.ScanOptions{
		All:       a.All,
		Recursive: true,
	}

	entries, err := a.Scanner.Scan(path, options)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	formatter.PrintTree(entries)
}

func (a *App) help() {
	fmt.Println("mdir")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  mdir")
	fmt.Println("  mdir <path>")
	fmt.Println("  mdir -a")
	fmt.Println("  mdir -all")
	fmt.Println("  mdir -r")
	fmt.Println("  mdir --recursive")
	fmt.Println("  mdir -help")
	fmt.Println("  mdir -v")
	fmt.Println("  mdir -ver")
	fmt.Println("  mdir -version")
}
func (a *App) version() {
	fmt.Println("mdir version:", Version)
	fmt.Println("Created by Mithun, MIT License, github.com/notmithun/mdir")
}
