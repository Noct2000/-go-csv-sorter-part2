package main

import (
	"task2/sorter"
)

func main() {
	done := make(chan struct{})
	pipeline := sorter.NewPipeline(3, done)
	res := pipeline.Run("fake-dir")

	for line := range res {
		println(line)
	}
}
