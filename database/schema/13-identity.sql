CREATE TABLE identity_providers (
    index TEXT NOT NULL,
    admin_accounts jsonb NOT NULL,
    provider_accounts jsonb NOT NULL,
    asset_minter_accounts jsonb NOT NULL,
    asset_burner_accounts jsonb NOT NULL,
    primary key (index)
);

CREATE TABLE identity_humans (
    index TEXT NOT NULL,
    vns_domain TEXT NOT NULL,
    accounts jsonb NOT NULL,
    networks jsonb NOT NULL,
    network_primary_wallet jsonb NOT NULL,
    primary key (index)
);

CREATE TABLE identity_accounts (
   index TEXT NOT NULL REFERENCES account (address),
   human_id TEXT UNIQUE REFERENCES identity_humans (index),
   private_acc bool NOT NULL,
   primary key (index)
);

CREATE TABLE identity_networks (
    index TEXT NOT NULL,
    active bool NOT NULL,
    full_name TEXT NOT NULL,
    identity_provider TEXT NOT NULL REFERENCES identity_providers (index),
    invite_only bool NOT NULL,
    primary key (index)
);

CREATE TABLE identity_account_networks (
    index TEXT NOT NULL
        REFERENCES account (address)
        REFERENCES identity_accounts (index),
    networks jsonb NOT NULL,
    primary key (index)
);

CREATE TABLE identity_invites (
    network TEXT NOT NULL,
    challenge TEXT NOT NULL,
    registered bool NOT NULL,
    confirmed_account TEXT,
    invite_creator TEXT NOT NULL
        REFERENCES account (address)
        REFERENCES identity_accounts (index),
    human_id TEXT NOT NULL REFERENCES identity_humans (index),
    given_roles text NOT NULL,
    primary key (challenge)
);

CREATE TABLE identity_kyc_statuses (
    human_id TEXT NOT NULL REFERENCES identity_humans (index),
    identity_provider TEXT NOT NULL REFERENCES identity_providers (index),
    status numeric NOT NULL,
    data_hash TEXT NOT NULL,
    timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    primary key (human_id, identity_provider)
);

CREATE TABLE identity_network_kyb (
    index TEXT NOT NULL,
    status numeric NOT NULL,
    data_hash TEXT NOT NULL,
    timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    metadata TEXT NOT NULL,
    primary key (index)
);

CREATE TABLE identity_account_link_proposals (
    index TEXT NOT NULL,
    proposer_account TEXT NOT NULL
        REFERENCES account (address)
        REFERENCES identity_accounts (index),
    human_id TEXT NOT NULL REFERENCES identity_humans (index),
    set_as_primary_wallet_for_network TEXT NOT NULL REFERENCES identity_networks (index),
    deposit jsonb NOT NULL,
    primary key (index)
);

CREATE TABLE identity_params (
    one_row_id     BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    granter_account text NOT NULL,
    granted_denom text NOT NULL,
    granted_amount bigint NOT NULL,
    spam_deposit_denom text NOT NULL,
    spam_deposit_amount bigint NOT NULL,
    height         BIGINT  NOT NULL,
    CHECK (one_row_id)
)