package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed torb-func.ps1
var psfileContents string

func checkPowerShellCmd(cmd string) (string, error) {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return "", err
	}
	return cmd, nil
}

func prepareArgs(args []string) string {
	var preparedArgs []string
	for _, arg := range args {
		preparedArgs = append(preparedArgs, fmt.Sprintf("'%s'", arg))
	}
	return strings.Join(preparedArgs, ", ")
}

func createTempfile() (string, error) {
	tempfile, err := os.CreateTemp("", "torb-func-*.ps1")
	if err != nil {
		return "", err
	}
	if _, err := tempfile.Write([]byte(psfileContents)); err != nil {
		return "", err
	}
	if err := tempfile.Close(); err != nil {
		return "", err
	}
	return tempfile.Name(), nil
}

func moveToRecycleBin(args []string) error {
	pwshCmd, err := checkPowerShellCmd("pwsh")
	if err != nil {
		pwshCmd, err = checkPowerShellCmd("powershell")
		if err != nil {
			return fmt.Errorf("PowerShell not found: %w", err)
		}
	}

	tempfile, err := createTempfile()
	if err != nil {
		return fmt.Errorf("create temp file error: %w", err)
	}
	defer os.Remove(tempfile)

	pwshArgs := prepareArgs(args)
	pwshCommand := fmt.Sprintf("& { . '%s'; MoveTo-RecycleBin @(%s) }", tempfile, pwshArgs)

	cmd := exec.Command(pwshCmd, "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", pwshCommand)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: torb.exe <file or folder> [...]")
		os.Exit(1)
	}

	if err := moveToRecycleBin(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
