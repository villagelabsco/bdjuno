CREATE TABLE products_product_class_infos (
    network text not null references identity_networks (index),
    class_id text not null,
    full_class_id text not null,
    class_type numeric not null,
    name text not null,
    description text not null,
    metadata jsonb not null,
    specific_metadata jsonb not null,
    primary key (full_class_id)
);