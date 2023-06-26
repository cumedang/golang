package main

import (
	"github.com/cumedang/go/cli"
	"github.com/cumedang/go/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
