package sorter

type Pipeline struct {
	n    int
	done chan struct{}
}

func NewPipeline(n int, done chan struct{}) *Pipeline {
	return &Pipeline{
		n:    n,
		done: done,
	}
}

func (p *Pipeline) Run(dirName string) (sortedContent chan string) {
	fnChan := p.readDir(dirName)
	contChan := p.fileReadingStage(fnChan, 3)
	return p.sortContent(contChan)
}
