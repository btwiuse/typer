package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/vyrx-dev/toofan/internal/tui"
)

var version = "dev"

func main() {
	v := flag.Bool("version", false, "Print current version")
	flag.Parse()

	if *v {
		fmt.Println("toofan " + version)
		os.Exit(0)
	}

	p := tea.NewProgram(tui.New())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
