CREATE TABLE marketplace_listings (
    network text references village_networks (index),
    index text primary key,
    reference text,
    product_class_id text references products_product_class_infos (full_class_id),
    product_nft_id text,
    attributes jsonb,
    creator text references account (address),
    active bool
);

CREATE TABLE marketplace_orders (
    index text primary key,
    network text references village_networks (index),
    status numeric not null,
    timestamp numeric not null,
    creator text references account (address),
    attributes jsonb not null,
    items jsonb not null,
    total jsonb not null
)
