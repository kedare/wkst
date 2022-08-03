package tools

import (
	"bufio"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

func streamToLog(scanner bufio.Scanner, fields log.Fields) {
	log.WithFields(fields).Debug("Starting Streaming")
	for scanner.Scan() {
		var message map[string]interface{}
		rawMessage := scanner.Text()
		err := json.Unmarshal([]byte(rawMessage), &message)
		if err != nil {
			log.WithFields(fields).Info(rawMessage)
		} else {
			log.WithFields(fields).WithFields(log.Fields{"json": message}).Info()
		}
	}
	log.WithFields(fields).Debug("Ending Streaming")
}

// ExecAndStream will execute the given command and stream stdout + stderr in the logger
func ExecAndStream(cmdCall string, parameters ...string) (*exec.Cmd, error) {
	log.WithFields(log.Fields{"cmd": cmdCall, "cmd_args": parameters}).Debug("Executing")
	cmd := exec.Command(cmdCall, parameters...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderrScanner := bufio.NewScanner(stderr)
	stdoutScanner := bufio.NewScanner(stdout)

	err = cmd.Start()
	go streamToLog(*stderrScanner, log.Fields{"stream": "stderr", "process": cmdCall, "pid": cmd.Process.Pid})

	go streamToLog(*stdoutScanner, log.Fields{"stream": "stdout", "process": cmdCall, "pid": cmd.Process.Pid})

	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
