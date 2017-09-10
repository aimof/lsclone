package main

import "flag"

var (
	aOpt   = flag.Bool("a", false, "show all hidden files")
	allOpt = flag.Bool("all", false, "show all hidden files")
	pOpt   = flag.Bool("p", true, "append / indicator to directories")
)
