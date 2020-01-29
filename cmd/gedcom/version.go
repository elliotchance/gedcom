package main

import "fmt"

// Version is replaced by CI before building in .travis.yml.
const Version = "unknown version"

func runVersionCommand() {
	fmt.Println(Version)
}
