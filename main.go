package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	fnChan := ReadDir("fake-dir")
	contChan := FileReadingStage(fnChan, 300)
	SortContent(contChan)
}

func SortContent(content chan string) {

	var buffer = make([]string, 0, 1000)
	for line := range content {
		buffer = append(buffer, line)
	}
	sort.Slice(buffer, func(i, j int) bool { return buffer[i] < buffer[j] })
	for i := range buffer {
		println(buffer[i])
	}

	for i := range buffer {
		println(buffer[i])
	}

}

func ReadDir(dir string) (fileNamesChan chan string) {
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

func FileReadingStage(fnames chan string, n int) (allLines chan string) {
	lines := make([]chan string, n)
	allLines = make(chan string)
	for i := 0; i < n; i++ {
		lines[i] = make(chan string)
		ReadFiles(fnames, lines[i])
	}
	go func() {
		defer close(allLines)
		wg := &sync.WaitGroup{}
		for i := range lines {
			wg.Add(1)
			go func(ch chan string) {
				defer wg.Done()
				for line := range ch {
					allLines <- line
				}
			}(lines[i])
		}
		wg.Wait()
	}()

	return allLines
}

func ReadFiles(fnames, lines chan string) {
	go func() {
		defer close(lines)
		for fname := range fnames {
			for i := 0; i < 10; i++ {
				lines <- fmt.Sprintf("%q content# %d\n", fname, i+1)
			}
		}
	}()
}
