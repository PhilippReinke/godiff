package dir

import (
	"fmt"
	"os"
)

func CheckExistenceWithExit(paths ...string) {
	for _, path := range paths {
		exists(path)
	}
}

func exists(path string) {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%v' does not exist\n", path)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Directory '%v' caueses unknown error\n", path)
		os.Exit(1)
	}
	if !fileInfo.IsDir() {
		fmt.Printf("'%v' is not a directory\n", path)
		os.Exit(1)
	}
}
