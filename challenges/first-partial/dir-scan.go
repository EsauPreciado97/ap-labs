package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./dir-scan <directory>")
		os.Exit(1)
	}
	scanDir(os.Args[1])
}

func scanDir(dir string) error {

	// Thanks to the walk function we are going to see every file and directory in the following tree
	// Then, all we have to do is to classify the files/directories according to corresponding the type

	directories := 0 // Directories
	symLinks := 0    // Symbolic Links
	devices := 0     // Devices
	sockets := 0     // Sockets
	other := 0       // Not defined

	err := filepath.Walk(dir, func(dir string, fileI os.FileInfo, err error) error {

		fileI, err = os.Lstat(dir)
		if fileI.Mode()&os.ModeDir != 0 {
			directories = directories + 1
		} else if fileI.Mode()&os.ModeSymlink != 0 {
			symLinks = symLinks + 1
		} else if fileI.Mode()&os.ModeDevice != 0 {
			devices = devices + 1
		} else if fileI.Mode()&os.ModeSocket != 0 {
			sockets = sockets + 1
		} else {
			other = other + 1
		}
		return err

	})

	// Printing the result in here.
	// I had trouble aligning the right vertical line of the table.

	fmt.Println("Directory Scanner Tool Results")
	fmt.Println("+-------------------------+-------------------+")
	fmt.Println("| Path                    |", dir, "         |")
	fmt.Println("+-------------------------+-------------------+")
	fmt.Println("| Directories             | ", directories, " |")
	fmt.Println("| Symbolic Links          | ", symLinks, "    |")
	fmt.Println("| Devices                 | ", devices, "     |")
	fmt.Println("| Sockets                 | ", sockets, "     |")
	fmt.Println("| Other files             | ", other, "       |")
	fmt.Println("+-------------------------+------------------+")
	return err
}
