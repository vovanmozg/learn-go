package main

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	freeFlowJobs := []job{
		job(func(in, out chan interface{}) {
			iVal := <-in
			val, ok := iVal.(int32)
			if !ok {
				fmt.Printf("Can't cast %v\n", val)
			}
			// to do something with val
			// we multiply it to 2
			out <- val * 2
		}),
		job(func(in, out chan interface{}) {
			iVal := <-in
			val, ok := iVal.(int32)
			if !ok {
				fmt.Printf("Can't cast %v\n", val)
			}
			out <- val * 3
		}),
		job(func(in, out chan interface{}) {
			val := <-in
			fmt.Println(val)
			out <- val
		}),
	}

	in, out := ExecutePipeline(freeFlowJobs...)
	input := []int32{1, 3, 4, 12}
	for _, val := range input {
		in <- val
		result := <-out
		if result != val*6 {
			t.Fail()
		}
	}
	close(in)
	close(out)
}
