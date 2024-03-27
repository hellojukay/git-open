package main

import (
	"flag"
	"log"
	"os"
	"runtime/debug"

	tui "github.com/manifoldco/promptui"
)

var (
	pipeline bool
	version  bool
)

func init() {
	flag.BoolVar(&pipeline, "p", false, "open github action page or gitlab pipeline page")
	flag.BoolVar(&version, "v", false, "show program version")
	flag.Parse()
	if version {
		printVersion()
		os.Exit(0)
	}
}
func main() {
	origins, err := GitRemotes()
	if err != nil {
		log.Fatal(err)
	}
	if err := OpenWithSelect(origins); err != nil {
		log.Fatal(err)
	}
}

func OpenWithSelect(origins []string) error {
	if len(origins) > 1 {
		prompt := tui.Select{
			Label: "Select a remote address",
			Items: origins,
		}
		_, origin, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
			return err
		}
		if pipeline {
			return Open(WithPipeline(origin))
		}
		return Open(origin)
	} else {
		if pipeline {
			return Open(WithPipeline(origins[0]))
		}
		return Open(origins[0])
	}
}

func printVersion() {
	info, ok := debug.ReadBuildInfo()
	if ok {
		println(info.Main.Version)
	}
}
