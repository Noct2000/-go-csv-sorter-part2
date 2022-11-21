package sorter

import "sort"

func (p *Pipeline) sortContent(content chan string) (res chan string) {
	res = make(chan string)
	defer close(res)
	go func() {
		var buffer = make([]string, 0, 1000)
		for line := range content {
			buffer = append(buffer, line)
		}
		sort.Slice(buffer, func(i, j int) bool { return buffer[i] < buffer[j] })
		for i := range buffer {
			println(buffer[i])
		}

		for i := range buffer {
			res <- buffer[i]
		}
	}()
	return res
}
