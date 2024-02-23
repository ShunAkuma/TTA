package main

import (
	"fmt"
	"word-search-in-files/pkg/searcher"
)

func main() {
	ss := searcher.Searcher{}

	list, err := ss.Search("s")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)

}
