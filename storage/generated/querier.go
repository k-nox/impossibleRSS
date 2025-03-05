// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package generated

import (
	"context"
)

type Querier interface {
	CreateFeed(ctx context.Context, arg CreateFeedParams) error
	CreateItem(ctx context.Context, arg CreateItemParams) error
	GetFeeds(ctx context.Context) ([]Feed, error)
	GetItemsForFeed(ctx context.Context, feedUrl string) ([]GetItemsForFeedRow, error)
}

var _ Querier = (*Queries)(nil)
