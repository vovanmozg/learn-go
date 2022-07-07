package main

import (
	"fmt"
	"time"
)



func main() {
	//ch := make(chan int)


	// при 1 выполнится таймаут, при 3 - выполнится операция
	ticker := time.NewTicker(1 * time.Second)


	for {
		x := <-ticker.C
		fmt.Println(x)
	}

}
