package feed

type Feed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Items       []*Item `json:"items"`
}

type Item struct {
	Content string `json:"content"`
}

type FeedList struct {
	urls []string
}

func NewFeedList() *FeedList {
	return &FeedList{
		urls: []string{
			"https://blog.luxatweb.dev/index.xml",
		},
	}
}

func (fl *FeedList) Add(url string) {
	fl.urls = append(fl.urls, url)
}

func (fl *FeedList) Urls() []string {
	return fl.urls
}
