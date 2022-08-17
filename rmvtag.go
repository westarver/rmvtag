package rmvtag

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/westarver/boa"
	"github.com/westarver/helper"
	msg "github.com/westarver/messenger"
)

const (
	CommentString   = "//"
	LineTag         = "<rmv/>"
	OpenTag         = "<rmv>"
	CloseTag        = "</rmv>"
	RegionTagPat    = `<rgn[[:blank:]]+_*[a-zA-Z0-9]+/?>`
	DefaultRangeLen = 20
)

//─────────────┤ Run ├─────────────

func Run(m *msg.Messenger) int {
	cli, action := prepArgs()
	var err error
	var ins, saves []string

	if action == actionHelp { //help is unique in that the command causes an early exit
		hlp, b := cli.Items["help"].(boa.CmdLineItem[string])
		if b {
			topic := hlp.Value()
			item, exist := cli.AllHelp[topic]
			if exist {
				ShowHelp(m, item)
			} else {
				ShowHelp(m)
			}
			return int(action)
		}
		ShowHelp(m)
		return int(actionHelp)
	}

	src, b1 := cli.Items["--source"].(boa.CmdLineItem[[]string])
	if b1 {
		ins = src.Value()
	}
	// no input file names passed on command line
	if len(ins) == 0 {
		// user enters file name at command line
		o, _ := os.Stdin.Stat()
		if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice { //Terminal
			for {
				fmt.Print("Enter a file path > ")
				inb := make([]byte, 256)
				n, _ := os.Stdin.Read(inb)
				if n > 1 {
					in := string(inb)
					in = strings.Trim(in, "\000\t \n")
					ins = append(ins, in)
				} else {
					if len(ins) == 0 {
						m.InfoMsg(m.Logout(), msg.MESSAGE, "No source file name was given")
						m.InfoMsg(m.Logout(), msg.MESSAGE, "Exiting with exit code %d", int(actionError))
						return int(actionError)
					} else {
						break
					}
				}
			}
		} else {
			m.InfoMsg(m.Logout(), msg.MESSAGE, "No source file name was given")
			m.InfoMsg(m.Logout(), msg.MESSAGE, "Exiting with exit code %d", int(actionError))
			return int(actionError)
		}
	}

	cpy, b2 := cli.Items["--copy"].(boa.CmdLineItem[[]string])
	if b2 {
		saves = cpy.Value()
	}

	var rng []int
	r, b3 := cli.Items["--range"].(boa.CmdLineItem[[]int])
	if b3 {
		rng = r.Value()
	}

	var rgn []string
	rg, b4 := cli.Items["--region"].(boa.CmdLineItem[[]string])
	if b4 {
		rgn = rg.Value()
		rng = getRangeFromRegion(ins[0], rgn...)
	}

	err = performCommand(ins, saves, rng, action)

	if m.Catch(msg.LOG, err) != nil {
		action = actionError
	}

	m.InfoMsg(m.Logout(), msg.MESSAGE, "Exiting with exit code %d", int(action))
	return int(action)
}

//─────────────┤ performCommand ├─────────────

func performCommand(ins, saves []string, rng []int, act actionType) error {
	matched := helper.Matchio(ins, saves)
	// use of ranges only makes sense for a single input
	if len(rng) > 0 {
		return process(act, matched[0].In, matched[0].Out, rng)
	}

	for _, m := range matched {
		if err := process(act, m.In, m.Out, rng); err != nil {
			return err
		}
	}
	return nil
}

//─────────────┤ process ├─────────────

func process(a actionType, in, out string, rng []int) error {
	var err error
	mock := false
	if a&actionMock == actionMock {
		mock = true
		a ^= actionMock
	}

	switch a {
	case actionRemove:
		err = doRemove(in, out, rng, mock)
	case actionComment:
		err = doComment(in, out, rng, mock)
	case actionUncomment:
		err = doUncomment(in, out, rng, mock)
	case actionList:
		err = doList(in)
	default:
		return errors.New("invalid command")
	}
	return err
}
