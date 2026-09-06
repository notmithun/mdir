package main

import (
	"os"

	"github.com/notmithun/mdir/internal/app"
)

func main() {
	application := app.New()
	application.Run(os.Args[1:])
}
