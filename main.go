package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PhilippReinke/godiff/dir"

	"github.com/alexflint/go-arg"
)

var (
	GitCommit = "unknown"
)

type args struct {
	Path1 string `arg:"positional,required" help:"first path for comparison"`
	Path2 string `arg:"positional,required" help:"second path for comparison"`
	// NumOfWorkers int    `arg:"-n,--numOfWorkers" default:"1" help:"number of workers"`
}

func (args) Version() string {
	return fmt.Sprintf("GitCommit: %v", GitCommit)
}

func main() {
	var args args
	arg.MustParse(&args)

	// determine absoulte paths
	path1Abs, err := filepath.Abs(args.Path1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	path2Abs, err := filepath.Abs(args.Path2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// check if input is sane
	if path1Abs == path2Abs {
		fmt.Println("Paths coincide. Nothing to do.")
		return
	}
	dir.CheckExistenceWithExit(args.Path1, args.Path2)

	// recursively read all files and directories
	dir1, err := dir.NewDir(args.Path1)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}
	dir2, err := dir.NewDir(args.Path2)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	// compare
	result := dir.Compare(dir1, dir2)
	fmt.Println(result)
}
