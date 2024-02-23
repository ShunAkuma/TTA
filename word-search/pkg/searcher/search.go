package searcher

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

type Searcher struct {
	FS fs.FS
}

func (s *Searcher) Search(word string) (files []string, err error) {
	// fileNames, err := dir.FilesFS(s.FS, "/Users/ashuramaru/TTA/word-search/examples")
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *Searcher) Test() {
	data, err := ioutil.ReadFile("/Users/ashuramaru/TTA/word-search/examples")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
func indexFile(filePath string) (map[string][]int, error) {
	index := make(map[string][]int)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		for _, word := range words {
			// Приводим слово к нижнему регистру для учета регистра
			word = strings.ToLower(word)
			// Добавляем номер строки в индекс
			index[word] = append(index[word], lineNumber)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return index, nil
}
