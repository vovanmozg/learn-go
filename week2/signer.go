package main

import (
	"sync"
)

func wrapper(job job, wg *sync.WaitGroup, in chan interface{}, out chan interface{}) {
	defer wg.Done()
	job(in, out)
	close(out)
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	wg.Add(len(jobs))

	var chans = []chan interface{}{
		make(chan interface{}),
		make(chan interface{}),
	}

	for job_index, job := range jobs {
		go wrapper(job, wg, chans[job_index], chans[job_index+1])
		chans = append(chans, make(chan interface{}))
	}

	wg.Wait()
}

var SingleHash = func(in chan interface{}, out chan interface{}) {
	for val := range in {
		out <- DataSignerCrc32(val.(string)) + "~" + DataSignerCrc32(DataSignerMd5(val.(string)))
	}
}

var MultiHash = func(in chan interface{}, out chan interface{}) {

}

var CombineResults = func(in chan interface{}, out chan interface{}) {

}
