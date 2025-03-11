-- name: CreateFeed :exec
insert into feeds (url, title, description) values (?, ?, ?);

-- name: CreateItem :exec
insert into items (
    guid, title, authors, content, description, published_date, feed_url
) values (?, ?, ?, ?, ?, ?, ?);

-- name: GetFeeds :many
select
    url,
    title,
    last_fetched_date,
    description
from feeds;

-- name: GetItemsForFeed :many
select
    guid,
    title,
    authors,
    content,
    description,
    published_date
from items
where feed_url = ?;
