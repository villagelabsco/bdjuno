CREATE TABLE marketplace_listings (
    network text not null references identity_networks (index),
    index text not null,
    reference text,
    product_class_id text references products_product_class_infos (full_class_id),
    product_nft_id text,
    attributes jsonb,
    creator text references account (address),
    active bool,
    primary key (network, index)
);

CREATE TABLE marketplace_orders (
    network text references identity_networks (index),
    index text not null,
    status numeric not null,
    timestamp timestamp not null,
    creator text references account (address),
    attributes jsonb not null,
    items jsonb not null,
    total jsonb not null,
    primary key (network, index)
)
