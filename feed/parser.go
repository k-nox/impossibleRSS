package feed

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
)

type Parser struct {
	parser    *gofeed.Parser
	sanitizer *bluemonday.Policy
}

func NewParser() *Parser {
	return &Parser{
		parser:    gofeed.NewParser(),
		sanitizer: bluemonday.UGCPolicy(),
	}
}

func (p *Parser) ParseURL(url string) (*Feed, error) {
	feed, err := p.parser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	items := make([]*Item, len(feed.Items))
	for index, item := range feed.Items {
		content := item.Content
		if content == "" {
			content = item.Custom["content"]
		}
		items[index] = &Item{
			Content: p.sanitizer.Sanitize(content),
		}
	}

	return &Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		Items:       items,
	}, nil
}
