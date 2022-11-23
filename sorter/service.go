package sorter

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func (p *Pipeline) ReadFromConsole(sortingFieldIndex int, isReversedOrder, isNotIgnoreHeader bool) string {
	scanner := bufio.NewScanner(os.Stdin)
	return p.processContent(sortingFieldIndex, isReversedOrder, isNotIgnoreHeader, scanner)
}

func (p *Pipeline) ReadFromFile(sortingFieldIndex int, isReversedOrder, isNotIgnoreHeader bool, inputFile string) string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	content := p.processContent(sortingFieldIndex, isReversedOrder, isNotIgnoreHeader, fileScanner)
	return content
}

func (p *Pipeline) processContent(sortingFieldIndex int, isReversedOrder, isNotIgnoreHeader bool, scanner *bufio.Scanner) string {
	var header string
	n := 0
	table := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, ",")
		if n == 0 {
			n = len(row)
			if isNotIgnoreHeader {
				header = line
				continue
			}
		}
		if line == "" {
			break
		}
		if n != len(row) {
			log.Fatalf("Error: row has %d columns, but must have %d\n", len(row), n)
		}
		table = append(table, row)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	sort.Slice(table, func(i, j int) bool {
		return p.Compare(table[i][sortingFieldIndex], table[j][sortingFieldIndex], isReversedOrder)
	})
	var result strings.Builder
	if header != "" {
		result.WriteString(header)
		result.WriteString("\n")
	}
	for _, row := range table {
		result.WriteString(strings.Join(row, ","))
		result.WriteString("\n")
	}
	return result.String()
}

func (p *Pipeline) processBufferContent(sortingFieldIndex int, isReversedOrder, isNotIgnoreHeader bool, buffer []string) []string {
	var header string
	n := 0
	table := [][]string{}
	for _, line := range buffer {
		row := strings.Split(line, ",")
		if n == 0 {
			n = len(row)
			if isNotIgnoreHeader {
				header = line
				continue
			}
		}
		if line == "" {
			break
		}
		if n != len(row) {
			log.Fatalf("Error: row has %d columns, but must have %d\n", len(row), n)
		}
		table = append(table, row)
	}
	sort.Slice(table, func(i, j int) bool {
		return p.Compare(table[i][sortingFieldIndex], table[j][sortingFieldIndex], isReversedOrder)
	})
	var result [] string
	if header != "" {
		result = append(result, header)
	}
	for _, row := range table {
		result = append(result, strings.Join(row, ","))
	}
	return result
}

func (p *Pipeline) WriteToFileIfPresent(content, fileName string) {
	if fileName != "" {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString(content)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (p *Pipeline) Compare(first, next string, isReversed bool) bool {
	return first < next != isReversed
}
