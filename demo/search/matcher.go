package search

import "log"

type Result struct {
	Field   string
	Content string
}

//定义 Matcher 接口 有一个Search 方法
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResult, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, result := range searchResult {
		results <- result
	}

}

func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
