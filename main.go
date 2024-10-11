package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xplshn/apkg/cmds/build"
	"github.com/xplshn/apkg/cmds/get"
	"github.com/xplshn/apkg/cmds/info"
	"github.com/xplshn/apkg/cmds/install"
	"github.com/xplshn/apkg/cmds/list"
	"github.com/xplshn/apkg/cmds/remove"
	"github.com/xplshn/apkg/cmds/verify"
)

func main() {
	log.SetFlags(0)
	switch os.Args[1] {
	case "build":
		build.Run()
	case "verify":
		verify.Run()
	case "install":
		install.Run()
	case "get":
		get.Run()
	case "remove":
		remove.Run()
	case "list":
		list.Run()
	case "info":
		info.Run()
	default:
		fmt.Println("not a command")
	}
}
