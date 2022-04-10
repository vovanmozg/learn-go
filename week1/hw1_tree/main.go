package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	//"strings"
)

/*

├───project                       d 0 1 1 <>
│ ├───file.txt (19b)              f 0 0 2 <│  >
│ └───gopher.png (70372b)         f 1 0 2 <│  >
├───static                        d 0 1 1 <>
│ ├───a_lorem                     d 0 0 2 <│  >
│ │ ├───dolor.txt (empty)         f 0 0 3 <│  │  >
│ │ ├───gopher.png (70372b)       f 0 0 3 <│  │  >
│ │ └───ipsum                     d 1 0 3 <│  │  >
│ │   └───gopher.png (70372b)     f 1 1 4 <│  │    >
│ ├───css													d 0 0 2 <│  >
│ │ └───body.css (28b)						f 1 0 3 <│  │  >
│ ├───empty.txt (empty)						f 0 0 2	<│  >
│ ├───html										  	d 0 0 2 <│  >
│ │ └───index.html (57b)					f 1 0 3 <│  │  >
│ ├───js													d 0 0 2 <│  >
│ │ └───site.js (10b)							f 1 0 3 <│  │  >
│ └───z_lorem                     d 1 0 2 <│  >
│   ├───dolor.txt (empty)         f 0 1 3 <│    >
│   ├───gopher.png (70372b)       f 0 1 3 <│    >
│   └───ipsum                     d 1 1 3 <│    >
│     ├───dolor                   d 0 1 4 <│      >
│     │ └───amet                  d 1 0 5 <│      │ >
│     ├───sit                     d 0 1 4
│     └───gopher.png (70372b)     f 1 1 4
├───zline
│ ├───empty.txt (empty)
│ └───lorem
│   ├───dolor.txt (empty)
│   ├───gopher.png (70372b)
│   └───ipsum
│     └───gopher.png (70372b)
└───zzfile.txt (empty)						f 1 1 1
																	│ │ │ └── level
																	│ │ └──── is last
																	│ └────── is parent last
																	└──────── dir / file


*/

func dirTree2(out io.Writer, path string, printFiles bool, parentPrefix string, isParentLast bool) error {
	allFiles, err := ioutil.ReadDir(path)

	if err != nil {
		return fmt.Errorf("error")
	}

	// Remove .DS_Store from files
	var files []fs.FileInfo
	for _, f := range allFiles {
		if f.Name() != ".DS_Store" {
			files = append(files, f)
		}
	}

	var dirs []fs.FileInfo
	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, f)
		}
	}

	if printFiles {
		fileIndex := 0
		for _, f := range files {
			prefix := ""

			// Print top-level prefixes
			fmt.Fprint(out, parentPrefix)

			if fileIndex != len(files)-1 { // If not last item
				fmt.Fprint(out, "├───")

				if isParentLast {
					prefix = parentPrefix + "	"
				} else {
					prefix = parentPrefix + "│	"
				}

			} else {
				fmt.Fprint(out, "└───")
				prefix = parentPrefix + "	"
			}

			fmt.Fprint(out, f.Name())
			if !f.IsDir() {
				size := "empty"
				if f.Size() != 0 {
					size = strconv.FormatInt(f.Size(), 10) + "b"
				}

				fmt.Fprint(out, " (", size, ")")
			}
			fmt.Fprintln(out, "")

			if f.IsDir() {
				dirTree2(out, filepath.Join(path, f.Name()), printFiles, prefix, isParentLast)
			}

			fileIndex++
		}
	} else {
		fileIndex := 0
		for _, f := range dirs {
			prefix := ""

			// Print top-level prefixes
			fmt.Fprint(out, parentPrefix)

			if fileIndex != len(dirs)-1 { // If not last item
				fmt.Fprint(out, "├───")

				if isParentLast {
					prefix = parentPrefix + "	"
				} else {
					prefix = parentPrefix + "│	"
				}

			} else {
				fmt.Fprint(out, "└───")
				prefix = parentPrefix + "	"
			}

			fmt.Fprint(out, f.Name())
			fmt.Fprintln(out, "")

			dirTree2(out, filepath.Join(path, f.Name()), printFiles, prefix, isParentLast)

			fileIndex++
		}
	}

	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTree2(out, path, printFiles, "", false)
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
