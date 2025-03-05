-- +migrate Up
create table feeds (
    url text primary key,
    title text,
    last_fetched_date timestamp not null default current_timestamp,
    description text
);

create table items (
    guid text primary key,
    title text,
    authors text,
    content text,
    description text,
    published_date timestamp,
    feed_url text not null,
    foreign key (feed_url) references feeds (url)
);

-- +migrate Down
drop table items;
drop table feeds;
