package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

type Repository struct {
	workdir string
	gitdir  string
	confdir string
	conf    *ini.File
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
	repo_create(path)
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
	rep.confdir = repo_file(true, "config")

	if _, err := os.Stat(rep.confdir); err == nil {
		cfg, err := ini.Load(rep.confdir + "/config")
		rep.conf = cfg
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cfg.Section("core").Key("repositoryformatversion").String())

	} else if !force {
		fmt.Println("No configuration file found")
	}

	if !force {
		vers, err := rep.conf.Section("core").Key("repositoryformatversion").Int()
		if err != nil {
			log.Fatal(err)
		}
		if vers != 0 {
			panic("Repository format version " + fmt.Sprint(vers) + " not supported")
		}
	}
}

func repo_path(path string) string {
	return filepath.Join(rep.gitdir, path)
}

func opt_file_slice(opt []string) (string, error) {
	if len(opt) > 0 {
		p := strings.Join(opt[:len(opt)-1], "")
		return p, nil
	} else {
		return "", fmt.Errorf("no file specified")
	}
}

func opt_file(opt []string) (string, error) {
	if len(opt) > 0 {
		return strings.Join(opt, ""), nil
	} else {
		return "", fmt.Errorf("no file specified")
	}
}

func repo_file(mkdir bool, opt ...string) string {
	path, err := opt_file_slice(opt)
	if err != nil {
		log.Fatal(err)
	}

	if repo_dir(mkdir, path) != "" {
		return repo_path(path)
	}
	return ""
}

func repo_dir(mkdir bool, opt ...string) string {
	path, err := opt_file(opt)
	if err != nil {
		log.Fatal(err)
	}

	path = repo_path(path)

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

func repo_create(path *string) {
	repo_init(path, true)
}
