CREATE TABLE products_products (
    network text not null references village_network (index),
    index text primary key,
    parent text references products_products (index),
    parent_chain text,
    has_children bit,
    name text not null,
    description text,
    attributes json,
    images json[],
    tags json[],
    p_type numeric not null,
    class text not null references classes_classes (index),
    creator text not null references account (address),
    active bit not null
);

CREATE TABLE products_product_class_infos (
    network text not null references village_network (index),
    type text not null,
    parent text,
    parent_chain text,
    has_children bool,
    disabled bool,
    is_soul_bond bool,
    base_image_uri text,
    metadata json
);