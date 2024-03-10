package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Repository struct {
	workdir string
	gitdir  string
	confdir string
	conf    string
}

var (
	rep    Repository
	forced *bool
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mydir" + dir)
	forced = flag.Bool("forced", false, "Use Forced mode")

	flag.Parse()
	fmt.Println("forced:", *forced)
	fmt.Println("tail:", flag.Args())
	fmt.Println("------------------")

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: sgit <command> [<args>]")
		os.Exit(1)
	}

	fmt.Println(args)

	switch args[0] {
	case "init":
		cmd_init(&dir)
	default:
		fmt.Println("Invalid command")
	}
}

func cmd_init(path *string) {
	*forced = true

	fmt.Println("path:", *path)
	fmt.Println("Initializing sgit")
	repo_init(path, *forced)
}

func repo_init(path *string, force bool) {
	rep.workdir = *path
	rep.gitdir = filepath.Join(rep.workdir, ".sgit")

	file, err := os.Stat(rep.workdir)
	if err != nil {
		log.Fatal(err)
	}

	// Check if Git repository exists
	if !force || !file.IsDir() {
		mess := "Not a Git repository: " + rep.workdir
		panic(mess)
	}

	// Read configuration file
	rep.confdir = repo_file("config/")
	fmt.Println(rep)

	fmt.Println("repo_init")
}

func repo_path(path string) string {
	return filepath.Join(rep.gitdir, path)
}

func opt_bool(opt []bool) bool {
	if len(opt) > 0 {
		return true
	} else {
		return false
	}
}

func repo_file(path string, opt ...bool) string {
	mkdir := opt_bool(opt)

	fmt.Println("path:", path)
	p := (path)[:len(path)-1]

	if repo_dir(p, mkdir) != "" {
		return repo_path(path)
	}
	return ""
}

func repo_dir(path string, opt ...bool) string {
	path = repo_path(path)
	mkdir := opt_bool(opt)

	file, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	if file != nil {
		if file.IsDir() {
			return path
		} else {
			message := "Not a directory: " + path
			panic(message)
		}
	}

	if mkdir {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Create directory")
		return path
	}
	return ""
}

func repo_create() {
	fmt.Println("Creating repository")
}
