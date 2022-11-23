package sorter

func (p *Pipeline) sortContent(content chan string) (res chan string) {
	res = make(chan string)
	go func() {
		defer close(res)
		var buffer = make([]string, 0, 1000)
		for line := range content {
			buffer = append(buffer, line)
		}
		content := p.processBufferContent(p.sortingFieldIndex, p.isReversedOrder, p.isNotIgnoreHeader, buffer)
		for _, i := range content {
			select {
			case res <- i:
				{
					continue
				}
			case <-p.done:
				{
					break
				}
			}
		}
	}()
	return res
}
