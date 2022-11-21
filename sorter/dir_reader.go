package sorter

func (p *Pipeline) readDir(dir string) (fileNamesChan chan string) {
	fileNamesChan = make(chan string)
	go func() {
		defer close(fileNamesChan)
		files := []string{
			"qwe", "asd", "zxc",
		}
		for _, f := range files {
			select {
			case fileNamesChan <- f:
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
	return fileNamesChan
}
