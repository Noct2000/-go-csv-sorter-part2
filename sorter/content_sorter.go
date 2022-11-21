package sorter

import "sort"

func (p *Pipeline) sortContent(content chan string) (res chan string) {
	res = make(chan string)
	go func() {
		defer close(res)
		var buffer = make([]string, 0, 1000)
		for line := range content {
			buffer = append(buffer, line)
		}
		sort.Slice(buffer, func(i, j int) bool { return buffer[i] < buffer[j] })
		for _, i := range buffer {
			res <- i
		}
	}()
	return res
}
