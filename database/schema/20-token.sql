CREATE TABLE token_tokens (
    network text not null,
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
    id numeric not null,
    account text not null
        references account (address)
        references identity_accounts (index),
    human_id text not null,
    executed boolean not null,
    amount jsonb not null,
    creation_block numeric not null,
    execution_block numeric not null,
    funds_transfer_method_pseudo_id text not null,
    id_provider text not null references identity_providers (index),
    primary key (id)
);

CREATE TABLE token_onramp_operations (
    payment_ref text not null,
    amount jsonb not null,
    account text not null
        references identity_accounts (index),
    primary key (payment_ref)
);

CREATE TABLE token_immobilized_funds (
    account text not null
        references account (address)
        references identity_accounts (index),
    amount jsonb not null,
    primary key (account)
);
