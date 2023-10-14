package main

import (
	"github.com/esh2n/switchchecker"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(switchchecker.Analyzer) }
