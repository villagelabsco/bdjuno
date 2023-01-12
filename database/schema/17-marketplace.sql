CREATE TABLE marketplace_listings (
    network text references village_network (index),
    index text primary key,
    nft text,
    attributes jsonb,
    creator text references account (address),
    active bit
);

CREATE TABLE marketplace_orders (
    index text primary key,
    network text references village_network (index),
    status numeric not null,
    timestamp numeric not null,
    creator text references account (address),
    attributes jsonb not null,
    items jsonb not null,
    total jsonb not null
)
