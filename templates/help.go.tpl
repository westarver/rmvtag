package main

import (
	"fmt"
	"os"
)

//────────────────────┤ getUsage ├────────────────────

func getUsage() string {
	var help = `{{file help.txt}}`
	return help
}

//────────────────────┤ ShowHelp ├────────────────────

func ShowHelp(err error, help string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", getUsage())
		os.Exit(0)
	} else {
		fmt.Println(getUsage())
	}
}

//────────────────────┤ showHelp ├────────────────────

func showHelp() {
	fmt.Println(getUsage())
}
