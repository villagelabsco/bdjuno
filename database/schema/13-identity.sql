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
    index TEXT NOT NULL REFERENCES identity_accounts (index),
    networks jsonb NOT NULL,
    primary key (index)
);

CREATE TABLE identity_invites (
    network TEXT NOT NULL,
    challenge TEXT NOT NULL,
    registered bool NOT NULL,
    confirmed_account TEXT NOT NULL REFERENCES account (address),
    invite_creator TEXT NOT NULL REFERENCES account (address),
    human_id TEXT NOT NULL REFERENCES identity_humans (index),
    given_roles text NOT NULL,
    primary key (challenge)
);

CREATE TABLE identity_statuses (
    provider_id TEXT NOT NULL REFERENCES identity_providers (index),
    human_id TEXT NOT NULL REFERENCES identity_humans (index),
    data_hash TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    primary key (human_id)
);

CREATE TABLE identity_nb_invite_per_day (
   network TEXT NOT NULL REFERENCES identity_networks (index),
   number TEXT NOT NULL,
   primary key (network)
);

CREATE TABLE identity_network_kyb (
    index TEXT NOT NULL,
    status numeric NOT NULL,
    data_hash TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    metadata TEXT NOT NULL,
    primary key (index)
);

CREATE TABLE identity_primary_wallet_transfer_proposals (
   index TEXT NOT NULL,
   proposer_account TEXT NOT NULL REFERENCES account (address),
   human_id TEXT NOT NULL REFERENCES identity_humans (index),
   set_as_primary_wallet_for_network TEXT NOT NULL REFERENCES identity_networks (index),
   deposit jsonb NOT NULL,
   primary key (index)
);

CREATE TABLE identity_human_proposals (
     human_id TEXT NOT NULL REFERENCES identity_humans (index),
     proposals jsonb NOT NULL,
     primary key (human_id)
);
