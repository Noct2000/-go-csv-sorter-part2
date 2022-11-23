package sorter

import (
	"log"
	"os"
	"path/filepath"
)

func (p *Pipeline) readDir(dir string) (fileNamesChan chan string) {
	fileNamesChan = make(chan string)
	go func() {
		defer close(fileNamesChan)
		files := p.iterate(dir)
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

func (p *Pipeline) iterate(path string) (files []string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files
}
