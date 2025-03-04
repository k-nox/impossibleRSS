package app

import (
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Title       string
	Description string
	Link        string
}

type Parser struct {
	parser gofeed.Parser
}

func NewParser() *Parser {
	return &Parser{
		parser: *gofeed.NewParser(),
	}
}

func (p *Parser) ParseURL(url string) (*Feed, error) {
	feed, err := p.parser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return &Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
	}, nil
}
