//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Open(url string) error {
	bin, err := exec.LookPath("open")
	if err != nil {
		return fmt.Errorf("can not find open, %s", err.Error())
	}
	c := exec.Command(bin, url)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
