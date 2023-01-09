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
)

/*
   string network = 1;
  string index = 2;
  string parent = 3;
  string parentChain = 4;
  bool hasChildren = 5;
  string name = 6;
  string description = 7;
  string attributes = 8;
  repeated string images = 9;
  repeated string tags = 10;
  ProductType ptype = 11;
  string classID = 12;
  string creator = 13;
  bool active = 14;
 */