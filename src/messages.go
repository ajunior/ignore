package main

import (
	"fmt"
	"runtime"
)

func printUsageMessage() {
	fmt.Printf("Usage: %s --create <path> <template1[, template2, template3, ...]>\n", prgName)
	fmt.Printf("Run 'ignore --help' for more information.\n")
}

func printHelpMessage() {
	path := "/path/to/project"

	if runtime.GOOS == "windows" {
		path = "C:\\path\\to\\project"
	}

	fmt.Print("Ignore helps you add GitHub-based .gitignore file to your local git repositories.\n\n")
	fmt.Print("Usage examples:\n\n")
	fmt.Println("  - 'ignore --version' to see the program version.")
	fmt.Println("  - 'ignore --templates' to list the available templates.")
	fmt.Printf("  - 'ignore --create %s java' to create a .gitignore from one template.\n", path)
	fmt.Printf("  - 'ignore --create %s java python go' to create a .gitignore from multiple templates.\n\n", path)
	fmt.Println("  * You can use one-letter shortcut for any command: -c, -h, -t, -v.")
}

func printVersionMessage() {
	fmt.Printf("%s v%s - %s\n", prgName, prgVersion, prgRepo)
}
