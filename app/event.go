package app

type Event string

const (
	RefreshError Event = "RefreshError"
	NewItem      Event = "NewItem"
)

var Events = []struct {
	Value  Event
	TSName string
}{
	{RefreshError, "REFRESH_ERROR"},
	{NewItem, "NEW_ITEM"},
}
