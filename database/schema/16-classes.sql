CREATE TABLE classes_classes (
    network text not null references village_network ("index"),
    index text primary key,
    name text not null,
    description text,
    parent text references classes_classes ("index"),
    parent_chain text,
    has_children bit,
    creator text references account ("address")
)

/*
 string networkID = 1;
  string index = 2;
  string name = 3;
  string description = 4;
  string parent = 5;
  string parentChain = 6;
  bool hasChildren = 7;
  string creator = 8;
 */