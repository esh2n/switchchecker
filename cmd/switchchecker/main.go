package main

import (
	"switchchecker"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(switchchecker.Analyzer) }
