package sorter

func (p *Pipeline) readDir(dir string) (fileNamesChan chan string) {
	fileNamesChan = make(chan string)
	go func() {
		defer close(fileNamesChan)
		files := []string{
			"qwe", "asd", "zxc",
		}
		for _, f := range files {
			fileNamesChan <- f
		}
	}()
	return fileNamesChan
}
