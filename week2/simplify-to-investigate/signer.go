package main

// сюда писать код
func ExecutePipeline(jobs ...job) (firstIn, lastOut chan interface{}) {
	firstIn = make(chan interface{})
	lastOut = make(chan interface{})
	prevOut := make(chan interface{})
	for i, j := range jobs {
		var in, out chan interface{}
		out = make(chan interface{})
		if i == 0 {
			in = firstIn
		}
		if i > 0 {
			in = prevOut
		}
		if i == len(jobs)-1 {
			out = lastOut
		}
		prevOut = out
		go j(in, out)
	}
	return
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
