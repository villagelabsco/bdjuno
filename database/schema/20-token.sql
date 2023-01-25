CREATE TABLE token_tokens (
    network text not null references identity_networks (index),
    denom text not null,
    ticker text not null,
    description text not null,
    nb_decimals numeric not null,
    transferable boolean not null,
    backing_asset text not null,
    admin text not null,
    name_id text not null,
    incentive_type text not null,
    icon_path text not null,
    referenced_denom text not null,
    offramp_enabled boolean not null,
    clawback_enabled boolean not null,
    clawback_period_sec numeric not null,
    primary key (denom)
);

CREATE TABLE token_offramp_operations (
    id serial not null,
    account text not null references account (address),
    human_id text not null,
    executed boolean not null,
    amount jsonb not null,
    creation_block numeric not null,
    execution_block numeric not null,
    funds_transfer_method_pseudo_id text not null,
    id_provider text not null,
    primary key (id)
);

CREATE TABLE token_onramp_operations (
    id serial not null,
    payment_ref text not null,
    amount jsonb not null,
    account text not null references account (address),
    primary key (id)
);

CREATE TABLE token_nb_token_creation_per_day (
    network text not null references identity_networks (index),
    number numeric not null,
    primary key (network)
);

CREATE TABLE token_last_input_activities (
    denom text not null references token_tokens (denom),
    account text not null references account (address),
    timestamp timestamp not null,
    block_height numeric not null
);

CREATE TABLE token_known_accounts (
    account text not null references account (address)
);

CREATE TABLE token_immobilized_funds (
    account text not null references account (address),
    amount jsonb not null
);

CREATE TABLE token_pending_balances (
    account text not null references account (address),
    amount jsonb not null
);

CREATE TABLE token_pending_clawbackable_operations (
    id serial not null,
    "from" text not null references account (address),
    "to" text not null references account (address),
    amount jsonb not null,
    clearing_timestamp timestamp not null,
    primary key (id)
);

CREATE TABLE token_pending_clawbackable_multi_operations (
   id serial not null,
   inputs jsonb not null,
   outputs jsonb not null,
   clearing_timestamp timestamp not null,
   primary key (id)
);

CREATE TABLE token_id_provider_manager_accounts (
    account text not null references account (address)
);
