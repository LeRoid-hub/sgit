package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"path/filepath"
	)

type Repository struct {
	workdir string
	gitdir string
	confdir string

}

var rep Repository
var forced *bool

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(dir)
	forced = flag.Bool("forced", false, "Use Forced mode")


	flag.Parse()
	fmt.Println("forced:", *forced)
	fmt.Println("tail:", flag.Args())

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: sgit <command> [<args>]")
		os.Exit(1)
	}


	fmt.Println(args)

	switch args[0] {
		case "init": cmd_init()
		default: fmt.Println("Invalid command")
	}
}

func cmd_init() {
	*forced = true

	fmt.Println("Initializing sgit")
}

func repo_create() {
	fmt.Println("Creating repository")
}

func repo_path(repo Repository, path *string) {
	fmt.Println("repo_path")
}

func repo_file(repo Repository, path *string) {
	fmt.Println("repo_file")

}

func repo_dir(repo Repository, path *string, mkdir bool) {
	fmt.Println("repo_dir")
}

