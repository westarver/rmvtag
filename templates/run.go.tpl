package main

import (
	"errors"

	"github.com/westarver/boa"
	msg "github.com/westarver/messenger"
)

func run(writer *msg.Messenger) int {
	var (
		cfg      string
		exitCode int
	)

	cli := boa.FromHelp(getUsage())

	help, hlp := cli.Items["help"].(boa.CmdLineItem[string])
	if hlp {
		topic := help.Value()
		item, exist := cli.AllHelp[topic]
		if exist {
			ShowHelp(writer, item)
		} else {
			ShowHelp(writer)
		}
		return 0
	}

	if !quiet {
		writer.SetLogoutStr(log)
	} else {
		_ = writer.LoggingOff() //turn off logging (-q | --quiet)
	}
	// run the program
	
	
	//writer.InfoMsg(writer.Logout(), msg.MESSAGE, "Exiting with exit code %d", exitCode)
	return exitCode
}
