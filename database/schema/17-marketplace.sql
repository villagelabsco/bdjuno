CREATE TABLE marketplace_listings (
    network text references village_network (index),
    index text primary key,
    product text references products_products (index),
    nft text,
    attributes json,
    creator text references account (address),
    active bit
);

CREATE TABLE marketplace_orders (
    index text primary key,
    network text references village_network (index),
    status numeric not null,
    timestamp numeric not null,
    creator text references account (address),
    attributes json not null,
    items json[] not null,
    total json not null
)
