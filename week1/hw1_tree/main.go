package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	//"strings"

)


func dirTree2(out io.Writer, path string, printFiles bool, level int) error {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return fmt.Errorf("error")
	}


	fileIndex := 0
	for _, f := range files {
		
		if f.IsDir() {
			for i := 0; i < level; i++ {
				fmt.Print("│ ")
			}

			if fileIndex < len(files) {
				fmt.Print("├─")	
			} else {
				fmt.Print("└─")	
			}

			
			// for i := 0; i < level; i++ {
			// 	fmt.Print("──")
			// }
			fmt.Println(f.Name() + "%v %v", fileIndex, len(files))

			dirTree2(out, filepath.Join(path, f.Name()), printFiles, level + 1)
		}
		
	}

	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTree2(out, path, printFiles, 0)
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
