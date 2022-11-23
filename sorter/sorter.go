package sorter

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Pipeline struct {
	n                 int
	done              chan struct{}
	sortingFieldIndex int
	isReversedOrder   bool
	isNotIgnoreHeader bool
}

func NewPipeline(n int, done chan struct{}, sortingFieldIndex int, isReversedOrder, isNotIgnoreHeader bool) *Pipeline {
	return &Pipeline{
		n:                 n,
		done:              done,
		sortingFieldIndex: sortingFieldIndex,
		isReversedOrder:   isReversedOrder,
		isNotIgnoreHeader: isNotIgnoreHeader,
	}
}

func (p *Pipeline) Run(dirName string) (sortedContent chan string) {
	fnChan := p.readDir(dirName)
	contChan := p.fileReadingStage(fnChan, 3)
	return p.sortContent(contChan)
}

func (p *Pipeline) WaitSignal(done chan struct{}) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		close(done)
	}()
}
