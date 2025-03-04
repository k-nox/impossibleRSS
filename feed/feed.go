package feed

type Feed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type FeedList struct {
	urls []string
}

func NewFeedList() *FeedList {
	return &FeedList{
		urls: []string{},
	}
}

func (fl *FeedList) Add(url string) {
	fl.urls = append(fl.urls, url)
}

func (fl *FeedList) Urls() []string {
	return fl.urls
}
