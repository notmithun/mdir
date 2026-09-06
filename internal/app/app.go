package app

import (
	"fmt"

	"github.com/notmithun/mdir/internal/filesystem"
	"github.com/notmithun/mdir/internal/formatter"
)

const Version = "v1"

type App struct {
	Scanner *filesystem.Scanner
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

		default:
			a.list(args[0])
		}
	default:
		fmt.Println("mdir: too many arguments")
	}
}

func (a *App) list(path string) {
	entries, err := a.Scanner.Scan(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	formatter.PrintEntries(entries)
}

func (a *App) help() {
	fmt.Println("Mithun Dir Tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  mdir")
	fmt.Println("  mdir <path>")
	fmt.Println("  mdir -help")
	fmt.Println("  mdir -v")
	fmt.Println("  mdir -ver")
	fmt.Println("  mdir -version")
}

func (a *App) version() {
	fmt.Println("mdir", Version)
}
