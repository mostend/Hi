//go:build darwin
// +build darwin

package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-cmd/cmd"
)

func (a *App) Start() {
	// cmdOptions Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	if strings.Contains(a.Core, "brook") {
		var brook Brook
		b, err := os.ReadFile(a.Conf)
		if err != nil {
			log.Fatalln(err)
		}
		_ = json.Unmarshal(b, &brook)
		a.Cmd = cmd.NewCmdOptions(cmdOptions, a.Core,
			brook.Server,
			brook.ServerType,
			brook.Addr,
			`--password`,
			brook.Password,
			`--http`,
			brook.Http,
			`--socks5`,
			brook.Socks5,
			brook.NoBrook,
		)
	} else {
		a.Cmd = cmd.NewCmdOptions(cmdOptions, a.Core, "-config", a.Conf)
	}

	//doneChan := make(chan struct{})
	//go func() {
	//	defer close(doneChan)
	//	// Done when both channels have been closed
	//	for a.Cmd.Stdout != nil || a.Cmd.Stderr != nil {
	//		select {
	//		case line, open := <-a.Cmd.Stdout:
	//			if !open {
	//				a.Cmd.Stdout = nil
	//				continue
	//			}
	//			fmt.Println(line)
	//		case line, open := <-a.Cmd.Stderr:
	//			if !open {
	//				a.Cmd.Stderr = nil
	//				continue
	//			}
	//			fmt.Println(line)
	//		}
	//	}
	//}()

	// Run and wait for Cmd to return, discard Status
	<-a.Cmd.Start()
}

func (a *App) Stop() {
	err := a.Cmd.Stop()
	fmt.Println(err)
}

func (a *App) ReStart() {
	a.Stop()
	a.Start()
}
