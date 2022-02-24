//go:build windows
// +build windows

package main

func Open(url string) error {
	bin, err := exec.LookPath("start")
	if err != nil {
		return fmt.Errorf("can not find start, %s", err.Error())
	}
	return exec.Command(bin, url).Run()
}
