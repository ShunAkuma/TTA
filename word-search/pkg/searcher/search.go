package searcher

import (
	"bufio"
	"io/fs"
	"os"
	"regexp"
	"strings"
	"sync"
	"word-search-in-files/pkg/internal/dir"
)

type Searcher struct {
	FS fs.FS
}

func (s *Searcher) Search(word string) (files []string, err error) {
	path := "Y:/Shun/TTA/word-search/examples/"

	fileNames, err := dir.FilesFS(s.FS, path)
	lenFile := len(fileNames)
	var resultList []string
	ch := make(chan string, lenFile)
	var wg sync.WaitGroup
	wg.Add(lenFile)
	if err != nil {
		return nil, err
	}
	for i := 0; i < lenFile; i++ {
		go func(fileName string) {
			defer wg.Done()
			indexFile(path, fileName, ch)
		}(fileNames[i])

	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		resultList = append(resultList, val) // Process the received value
	}

	return resultList, nil
}

func indexFile(filePath string, fileName string, ch chan string) error {
	path := filePath + fileName
	index := make(map[string][]int)
	reg := regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9]+`)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		cleanLine := reg.ReplaceAllString(line, " ")
		words := strings.Fields(cleanLine)

		for _, word := range words {
			// Приводим слово к нижнему регистру для учета регистра
			word = strings.ToLower(word)
			// Добавляем номер строки в индекс
			index[word] = append(index[word], lineNumber)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	result := searchSubstring(index, "пока")
	if result {
		ch <- fileName
	}
	return nil
}

func searchSubstring(index map[string][]int, substring string) bool {
	// Приводим подстроку к нижнему регистру для учета регистра
	substring = strings.ToLower(substring)
	exist := index[substring]
	return exist != nil
}
