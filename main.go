package main

import (
	"flag"
	"fmt"
	"log"

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

	// Define command flags
	buildFlag := flag.Bool("build", false, "Build the package")
	verifyFlag := flag.Bool("verify", false, "Verify the package")
	installFlag := flag.Bool("install", false, "Install the package")
	getFlag := flag.Bool("get", false, "Get the package")
	removeFlag := flag.Bool("remove", false, "Remove the package")
	listFlag := flag.Bool("list", false, "List the packages")
	infoFlag := flag.Bool("info", false, "Show package information")

	// Parse flags
	flag.Parse()

	// Handle the commands based on the flags
	switch {
	case *buildFlag:
		build.Run()
	case *verifyFlag:
		verify.Run()
	case *installFlag:
		install.Run()
	case *getFlag:
		get.Run()
	case *removeFlag:
		remove.Run()
	case *listFlag:
		list.Run()
	case *infoFlag:
		info.Run()
	default:
		fmt.Println("Please specify a valid command. Use --build, --verify, --install, --get, --remove, --list, or --info.")
	}
}
