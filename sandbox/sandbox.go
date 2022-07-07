package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type job func(in, out chan interface{})

func main() {
	inputData := []int{3, 1, 1, 2, 3, 5, 8}

	jobs := []job{
		job(func(in, out chan interface{}) {
			for _, fibNum := range inputData {
				out <- fibNum
			}
		}),
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			//data, ok := dataRaw.(string)
			// if !ok {
			// 	fmt.Println("cant convert result data to string")
			// }
			spew.Dump(dataRaw)

		}),
	}
	ExecutePipeline(jobs...)
	time.Sleep(5 * time.Second)
}

func ExecutePipeline(jobs ...job) {
	job1 := jobs[0]
	job2 := jobs[1]

	in := make(chan interface{})
	out := make(chan interface{})

	go job1(in, out)
	go job2(out, out)

	// for _, oneJob := range jobs {
	// 	oneJob(in, out)
	// }
}
