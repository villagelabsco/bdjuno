CREATE TABLE products_product_class_infos (
    network text not null references village_network (index),
    class_id text not null,
    full_class_id text not null primary key,
    class_type numeric not null,
    name text not null,
    description text not null,
    metadata jsonb not null,
    specific_metadata jsonb not null
);