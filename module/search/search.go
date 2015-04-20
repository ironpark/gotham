package search

import (
	"github.com/blevesearch/bleve"
)

func AddIndex(id, username, text, url string) {
	message := struct {
		ID   string
		User string
		Text string
		URL  string
	}{
		ID:   id,
		User: username,
		Text: text,
		URL:  url,
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("index.bleve", mapping)
	if err != nil {
		panic(err)
	}
	index.Index(message.ID, message)
}

func Search(text string) {
	index, _ := bleve.Open("index.bleve")
	query := bleve.NewQueryStringQuery(text)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)
	searchResult.String()
}
