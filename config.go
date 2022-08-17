package rmvtag

import (
	"github.com/westarver/boa"
)

type actionType int

const (
	actionHelp actionType = (1 << iota)
	actionList
	actionComment
	actionUncomment
	actionRemove
	actionMock
	actionError
)

//────────────────────┤ prepArgs ├────────────────────

func prepArgs() (*boa.CLI, actionType) {
	cli := boa.FromHelp(getUsage())
	//fmt.Println("cli", cli)
	var action actionType

	_, b1 := cli.Items["comment"].(boa.CmdLineItem[bool])
	if b1 {
		action = actionComment
	}

	_, b2 := cli.Items["uncomment"].(boa.CmdLineItem[bool])
	if b2 {
		action = actionUncomment
	}

	_, b3 := cli.Items["remove"].(boa.CmdLineItem[bool])
	if b3 {
		action = actionRemove
	}

	_, b4 := cli.Items["list"].(boa.CmdLineItem[bool])
	if b4 {
		action = actionList
	}
	_, b5 := cli.Items["help"].(boa.CmdLineItem[string])
	if b5 { //help is unique in that the command causes an early exit
		return cli, actionHelp
	}
	if !b1 && !b2 && !b3 && !b4 && !b5 { // default command is comment
		return cli, actionHelp
	}

	_, b6 := cli.Items["--mock"].(boa.CmdLineItem[bool])
	if b6 {
		if action&actionList == 0 { //list and --mock are mutually exclusive
			action |= actionMock
		}
	}

	return cli, action
}
