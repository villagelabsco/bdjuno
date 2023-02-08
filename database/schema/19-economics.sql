CREATE TABLE economics_transaction_hooks (
    network text not null references identity_networks (index),
    index numeric not null,
    trigger numeric not null,
    type numeric not null,
    name_id text not null,
    description text,
    params jsonb not null,
    primary key (network, index)
);

CREATE TABLE economics_scheduled_hooks (
    network text not null references identity_networks (index),
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

CREATE TABLE economics_network_enabled (
    network text not null references identity_networks (index),
    active boolean not null,
    primary key (network)
);

CREATE TABLE economics_scheduled_hooks_manual_triggers (
    id serial not null,
    creator text not null references account (address),
    network text not null references identity_networks (index),
    hook_idxs jsonb not null,
    primary key (id)
);