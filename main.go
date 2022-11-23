package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"task2/sorter"
)

func main() {
	fmt.Println("== Started ==")
	done := make(chan struct{})
	defer close(done)
	inputFileName := flag.String("i", "", "Use a file with the name file-name as an input.")
	outputFileName := flag.String("o", "", "Use a file with the name file-name as an output.")
	sortingFilesIndex := flag.Int("f", 0, "Sort input lines by value number N.")
	isNotIgnoreHeader := flag.Bool("h", false, "The first line is a header that must be ignored during sorting but included in the output.")
	isReversedOrder := flag.Bool("r", false, "Sort input lines in reverse order.")
	dirName := flag.String("d", "", "dir-name that specifies a directory where it must read input files from.")
	flag.Parse()
	sorter := sorter.NewPipeline(3, done, *sortingFilesIndex, *isReversedOrder, *isNotIgnoreHeader)
	sorter.WaitSignal(done)
	var content string
	if *inputFileName == "" && *dirName == "" {
		content = sorter.ReadFromConsole(*sortingFilesIndex, *isReversedOrder, *isNotIgnoreHeader)
	} else if *inputFileName != "" && *dirName != "" {
		log.Fatal("Invalid options. Don't use -i and -d together")
	} else if *inputFileName != "" {
		content = sorter.ReadFromFile(*sortingFilesIndex, *isReversedOrder, *isNotIgnoreHeader, *inputFileName)
	} else if *dirName != "" {
		res := sorter.Run(*dirName)
		var stringRes strings.Builder
		for i := range res {
			stringRes.WriteString(i + "\n")
		}
		content = stringRes.String()
	}
	if content != "" {
		fmt.Println("sorted result:\n" + content)
		sorter.WriteToFileIfPresent(content, *outputFileName)
	}
	fmt.Println("== Finished ==")
}
