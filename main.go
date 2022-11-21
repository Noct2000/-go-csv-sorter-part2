package main

import (
	"task2/sorter"
)

func main() {
	done := make(chan struct{})
	sorter := sorter.NewPipeline(3, done)
	res := sorter.Run("fake-dir")
	for i := range res {
		println(i)
	}
}
