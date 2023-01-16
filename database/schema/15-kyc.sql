CREATE TABLE kyc_humans (
   index TEXT PRIMARY KEY NOT NULL,
   vns_domain TEXT NOT NULL,
   accounts jsonb NOT NULL,
   network_primary_wallet jsonb NOT NULL
);

CREATE TABLE kyc_accounts (
    index TEXT PRIMARY KEY NOT NULL REFERENCES account (address),
    human_id TEXT UNIQUE REFERENCES kyc_humans (index),
    private_acc bool NOT NULL
);

CREATE TABLE kyc_invites (
    network TEXT NOT NULL,
    challenge TEXT PRIMARY KEY NOT NULL,
    registered bool NOT NULL,
    confirmed_account TEXT NOT NULL REFERENCES account (address),
    invite_creator TEXT NOT NULL REFERENCES account (address),
    human_id TEXT NOT NULL REFERENCES kyc_humans (index),
    given_roles text NOT NULL
);

CREATE TABLE kyc_status (
    provider_id TEXT NOT NULL REFERENCES kyc_identity_provider (index),
    human_id TEXT NOT NULL REFERENCES kyc_humans (index),
    data_hash TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE kyc_nb_invite_per_day (
    network TEXT NOT NULL REFERENCES village_networks (index),
    number TEXT NOT NULL
);

CREATE TABLE kyc_network_kyb (
    index TEXT PRIMARY KEY NOT NULL,
    status TEXT NOT NULL,
    data_hash TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    metadata TEXT NOT NULL
);

CREATE TABLE kyc_primary_wallet_transfer_proposals (
    index TEXT PRIMARY KEY NOT NULL,
    proposer_account TEXT NOT NULL REFERENCES account (address),
    human_id TEXT NOT NULL REFERENCES kyc_humans (index),
    set_as_primary_wallet_for_network TEXT NOT NULL REFERENCES village_networks (index),
    deposit jsonb NOT NULL
);

CREATE TABLE kyc_human_proposals (
    human_id TEXT NOT NULL REFERENCES kyc_humans (index),
    proposals jsonb NOT NULL
);

