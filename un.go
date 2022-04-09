package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	in := bufio.NewScanner(os.Stdin)
	var prev strin
	for in.Scan() {
		txt := in.Text()

		
		if txt == prev {
			continue
		}

		alreadySeen[txt] = true
		fmt.Println(txt)
	}
}