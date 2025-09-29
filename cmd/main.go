package main

import (
	"github.com/Komilov31/distributed-cut/pkg/cut"
	"github.com/Komilov31/distributed-cut/pkg/flags"
)

func main() {
	flags := flags.Parse()
	cut := cut.New(flags)

	cut.ProcessProgram()
}
