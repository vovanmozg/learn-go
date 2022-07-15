package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

func wrapper(job job, wg *sync.WaitGroup, in chan interface{}, out chan interface{}) {
	fmt.Println("wrapper: START")
	defer wg.Done()
	defer close(out)

	job(in, out)
	fmt.Println("wrapper: Prepare to close out: " + spew.Sdump(job))
	fmt.Println("wrapper: FINISH")
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}

	var chans = []chan interface{}{
		make(chan interface{}),
		make(chan interface{}),
	}

	for job_index, job := range jobs {
		wg.Add(1)
		go wrapper(job, wg, chans[job_index], chans[job_index+1])
		chans = append(chans, make(chan interface{}))
	}

	wg.Wait()
}

// var SingleHash = func(in chan interface{}, out chan interface{}) {
// 	fmt.Println("SingleHash: START")

// 	mu := &sync.Mutex{}
// 	var worker = func() {
// 		fmt.Println("worker: START")
// 		for val := range in {
// 			val_int := val.(int)
// 			val_string := strconv.Itoa(val_int)
// 			mu.Lock()
// 			md5 := DataSignerMd5(val_string)
// 			mu.Unlock()

// 			crc_val := DataSignerCrc32(val_string)
// 			crc_md5 := DataSignerCrc32(md5)
// 			joined := crc_val + "~" + crc_md5
// 			spew.Dump("worker: preparing to write 'out': "+joined, out)
// 			out <- joined
// 			spew.Dump("worker: wrote to 'out': "+joined, out)
// 		}
// 		fmt.Println("worker: FINISH")
// 	}

// 	for w := 1; w <= 3; w++ {
// 		go worker()
// 	}
// 	fmt.Println("SingleHash: FINISH")
// }

var WorkerSingleHash = func(out chan interface{}, val_string string, md5 string, wg *sync.WaitGroup) {

	async_result1 := make(chan interface{})
	async_result2 := make(chan interface{})

	go func() {
		async_result1 <- DataSignerCrc32(val_string)
	}()

	go func() {
		async_result2 <- DataSignerCrc32(md5)
	}()

	crc_val := <-async_result1
	crc_md5 := <-async_result2
	joined := crc_val.(string) + "~" + crc_md5.(string)
	fmt.Println("SingleHash: preparing to write 'out': " + joined)
	out <- joined
	fmt.Println("SingleHash: wrote to 'out': " + joined)

	wg.Done()
}

var SingleHash = func(in chan interface{}, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for val := range in {
		fmt.Println("SingleHash: LOOP START: " + strconv.Itoa(val.(int)))
		val_int := val.(int)
		val_string := strconv.Itoa(val_int)
		md5 := DataSignerMd5(val_string)

		wg.Add(1)
		go WorkerSingleHash(out, val_string, md5, wg)
	}
	wg.Wait()
}

var CalcMultiHash = func(wg *sync.WaitGroup, i int, val string, result []string) {
	result[i] = DataSignerCrc32(strconv.Itoa(i) + val)
	wg.Done()
}

var WorkerMultiHash = func(val string, out chan interface{}, wg *sync.WaitGroup) {
	wg2 := &sync.WaitGroup{}
	result := make([]string, 6)

	for i := 0; i < 6; i++ {
		wg2.Add(1)
		go CalcMultiHash(wg2, i, val, result)
	}

	wg2.Wait()

	buffer := strings.Join(result, "")
	fmt.Println("MultiHash: preparing to write buffer to 'out': " + buffer)
	out <- buffer
	fmt.Println("MultiHash: Wrote buffer to 'out': " + buffer)
	wg.Done()
}

var MultiHash = func(in chan interface{}, out chan interface{}) {
	fmt.Println("MultiHash: START")
	var wg sync.WaitGroup
	for val := range in {
		fmt.Println("MultiHash: Readed val from channel: " + val.(string))
		valStr := val.(string)

		wg.Add(1)
		go WorkerMultiHash(valStr, out, &wg)
	}
	wg.Wait()
}

var CombineResults = func(in chan interface{}, out chan interface{}) {

	var values []string
	for val := range in {
		valString := val.(string)
		values = append(values, valString)
	}

	sort.Strings(values)
	out <- strings.Join(values, "_")
}
