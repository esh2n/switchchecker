package main

import (
	"github.com/esh2n/switchchecker"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(switchchecker.Analyzer) }
