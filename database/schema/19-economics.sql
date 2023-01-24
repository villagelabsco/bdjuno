CREATE TABLE economics_transaction_hooks (
    network text not null references village_networks (index),
    index numeric not null,
    trigger numeric not null,
    type numeric not null,
    name_id text not null,
    description text,
    params jsonb not null,
    primary key (network, index)
);

CREATE TABLE economics_scheduled_hooks (
    network text not null references village_networks (index),
    index numeric not null,
    type numeric not null,
    name_id text not null,
    description text,
    cron_rule text,
    dependencies jsonb,
    auto_trigger boolean not null,
    params jsonb not null,
    last_executed_timestamp timestamp not null,
    last_executed_block numeric not null,
    primary key (network, index)
);

CREATE TABLE economics_nb_tx_per_day (
    network text not null references village_networks (index),
    number numeric not null,
    primary key (network)
);

CREATE TABLE economics_network_enabled (
    network text not null references village_networks (index),
    active boolean not null,
    primary key (network)
);

CREATE TABLE economics_pending_tasks (

)