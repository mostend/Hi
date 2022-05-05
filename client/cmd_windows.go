//go:build windows
// +build windows

package client

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/go-cmd/cmd"
)

func (a *App) Start() {
	// cmdOptions Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
		BeforeExec: []func(cmd *exec.Cmd){
			func(cmd *exec.Cmd) {
				if runtime.GOOS == "windows" {
					cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				}
			},
		},
	}
	a.Cmd = cmd.NewCmdOptions(cmdOptions, a.Core, "-c", a.Conf)

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
