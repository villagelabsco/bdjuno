CREATE TABLE marketplace_listings (
    network text references village_network (index),
    index text primary key,
    product text references products_products (index),
    nft text,
    attributes json,
    creator text references account (address),
    active bit
)
