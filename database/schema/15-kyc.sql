CREATE TABLE kyc_human (
   index TEXT PRIMARY KEY NOT NULL,
   vns_domain TEXT NOT NULL,
   accounts JSON NOT NULL,
   network_primary_wallet JSON NOT NULL
);

CREATE TABLE kyc_account (
    index TEXT PRIMARY KEY NOT NULL REFERENCES account (address),
    human_id TEXT UNIQUE REFERENCES kyc_human (index),
    private_acc BIT NOT NULL DEFAULT 0::BIT
);

CREATE TABLE kyc_invite (
    network TEXT NOT NULL,
    challenge TEXT PRIMARY KEY NOT NULL,
    registered BIT NOT NULL DEFAULT 0::BIT,
    confirmed_account TEXT NOT NULL REFERENCES account (address),
    invite_creator TEXT NOT NULL REFERENCES account (address),
    human_id TEXT NOT NULL REFERENCES kyc_human (index),
    given_roles text NOT NULL
);

CREATE TABLE kyc_status (
    provider_id TEXT NOT NULL REFERENCES kyc_identity_provider (index),
    human_id TEXT NOT NULL REFERENCES kyc_human (index),
    data_hash TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE kyc_nb_invite_per_day (
    network TEXT NOT NULL REFERENCES village_network (index),
    number TEXT NOT NULL
);

CREATE TABLE kyc_network_kyb (
    index TEXT PRIMARY KEY NOT NULL,
    status TEXT NOT NULL,
    data_hash TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    metadata TEXT NOT NULL
);

CREATE TABLE kyc_primary_wallet_transfer_proposal (
    index TEXT PRIMARY KEY NOT NULL,
    proposer_account TEXT NOT NULL REFERENCES account (address),
    human_id TEXT NOT NULL REFERENCES kyc_human (index),
    set_as_primary_wallet_for_network TEXT NOT NULL REFERENCES village_network (index),
    deposit JSON NOT NULL
);

CREATE TABLE kyc_human_proposals (
    human_id TEXT NOT NULL REFERENCES kyc_human (index),
    proposals JSON NOT NULL
);

