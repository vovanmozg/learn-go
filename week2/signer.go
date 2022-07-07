package main

import (
	"fmt"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

// func ExecutePipeline(jobs ...job) {
// 	ch1 := make(chan interface{}, 2)

// 	ch1 <- 1
// 	fmt.Println(1)
// 	ch1 <- 2
// 	fmt.Println(2)
// 	ch1 <- 3
// 	fmt.Println(3)
// 	close(ch1)
// 	for i := range ch1 {
// 		fmt.Println(i)
// 	}

// }

func wrapper1(job job, wg *sync.WaitGroup, in chan interface{}, out chan interface{}) {
	defer wg.Done()
	// notused := make(chan interface{})
	fmt.Println("Start job 1")
	job(in, out)
	fmt.Println("Finish job 1")
	close(out)
}

func ExecutePipeline(jobs ...job) {
	//notused := make(chan interface{})
	//ch2 := make(chan interface{})
	//ch3 := make(chan interface{})
	//job1done :=make(chan interface{})
	// ch4 := make(chan interface{})
	// ch5 := make(chan interface{})

	wg := &sync.WaitGroup{}
	wg.Add(len(jobs))

	ch_in := make(chan interface{})
	for _, job := range jobs {
		ch_out := make(chan interface{})
		go wrapper1(job, wg, ch_in, ch_out)
		ch_in = ch_out
	}
	// go wrapper1(jobs[0], wg, notused, ch2)
	// go wrapper1(jobs[1], wg, ch2, ch3)
	// go wrapper1(jobs[2], wg, ch3, notused)

	// go func() {
	// 	jobs[0](notused, ch2)
	// 	wg.Done()
	// }()
	// go func() {
	// 	fmt.Println("Start job 2")
	// 	defer wg.Done()
	// 	jobs[1](ch2, ch3)
	// 	fmt.Println("Finish job 2")
	// 	close(ch3)
	// }()
	// go func() {
	// 	fmt.Println("Start job 3")
	// 	defer wg.Done()
	// 	jobs[2](ch3, notused)
	// 	fmt.Println("Finish job 3")
	// 	close(notused)

	// }()
	wg.Wait()
	fmt.Println(">>>>>>>>>>")

	// LOOP:
	// 	for {
	// 		select {
	// 		case val := <-ch1:
	// 			fmt.Println(val)
	// 		default:
	// 			fmt.Println("default")
	// 			break LOOP
	// 		}
	// 	}

}

// сюда писать код
func ExecutePipeline2(jobs ...job) {
	fmt.Println("--------------------------------------")
	//ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})
	ch4 := make(chan interface{})
	ch5 := make(chan interface{})

	go jobs[1](ch3, ch4)
	go jobs[2](ch5, ch5)

	func() {
		jobs[0](ch2, ch2)
	LOOP1:
		for {
			select {
			case val := <-ch2:
				fmt.Println("S: ch3 <- ", val)
				ch3 <- val
				fmt.Println("F: ch3 <- ", val)
			default:
				fmt.Println("default")
				break LOOP1
			}

		}
		fmt.Println("func 1 finished")
	}()

LOOP:
	for {
		select {
		case val := <-ch4:
			fmt.Println("S: ch5 <- ", val)
			ch5 <- val
			fmt.Println("S: ch5 <- ", val)
			//spew.Dump(val)
		default:
			spew.Dump("default")
			break LOOP
		}

	}

	//spew.Dump(<-ch2)
	//spew.Dump(<-ch2)
	// for i := range ch2 {
	// 	spew.Dump(i)
	// 	//	ch3 <- i
	// }
}

// var SingleHash = func(in, out chan interface{}) {
// 	for i := range in {
// 		fmt.Println(i)
// 		// out <- DataSignerCrc32(i) + "~" + DataSignerCrc32(DataSignerMd5(i))
// 	}
// }

// var MultiHash = func(in, out chan interface{}) {

// }

// var CombineResults = func(in, out chan interface{}) {

// }
