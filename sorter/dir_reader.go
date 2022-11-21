package sorter

func (p *Pipeline) readDir(dir string) (fileNamesChan chan string) {
	fileNamesChan = make(chan string)
	go func() {
		defer close(fileNamesChan)
		files := []string{
			"qwe", "asd", "zxc",
		}
		for _, f := range files {
			println(f)
			select {
			case fileNamesChan <- f:
				{
					continue
				}
			case <-p.done:
				{
					println("stop")
					return
				}
			}

		}
	}()
	return fileNamesChan
}
