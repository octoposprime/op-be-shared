package tshell

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

type ShellExecutor struct {
	ManagerCtx context.Context
	ResultCh   chan string
	CmdName    string
	CmdParams  []string
	FinishCB   func()
	ErrorCB    func(err error)
}

func NewShellExecutor(managerCtx context.Context,
	resultCh chan string,
	cmdName string,
	cmdParams []string,
	finishCB func(),
	errorCB func(err error)) *ShellExecutor {
	return &ShellExecutor{
		ManagerCtx: managerCtx,
		ResultCh:   resultCh,
		CmdName:    cmdName,
		CmdParams:  cmdParams,
		FinishCB:   finishCB,
		ErrorCB:    errorCB,
	}
}

func (s *ShellExecutor) ExecCommand() {
	cmd := exec.Command(s.CmdName, s.CmdParams...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		s.ErrorCB(err)
		return
	}
	defer cmdReader.Close()

	done := make(chan bool)
	/*scanner := bufio.NewScanner(cmdReader)
	go func(scanner *bufio.Scanner, out chan string, done chan bool) {
		for scanner.Scan() {
			val := scanner.Text()
			out <- val
		}
		done <- true
		if scanner.Err() != nil {
			s.ErrorCB(scanner.Err())
		}
	}(scanner, s.ResultCh, done)*/

	go func(cmd *exec.Cmd) {
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			s.ErrorCB(fmt.Errorf(stderr.String()+" error: %w", err))
			return
		}
	}(cmd)

	go func(done chan bool) {
		time.Sleep(5 * time.Second)
		cmd.Wait()
		done <- true
	}(done)

	for {
		select {
		case <-done:
			s.FinishCB()
			close(s.ResultCh)
		case <-s.ManagerCtx.Done():
			err := cmd.Process.Kill()
			if err != nil {
				s.ErrorCB(err)
			}
			return
		}
	}
}
