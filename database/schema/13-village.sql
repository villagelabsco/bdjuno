CREATE TABLE village_networks (
    index TEXT PRIMARY KEY NOT NULL,
    active bool NOT NULL,
    full_name TEXT NOT NULL,
    identity_provider TEXT NOT NULL REFERENCES kyc_identity_provider (index),
    invite_only bool NOT NULL
);

CREATE TABLE village_user_networks (
    index TEXT PRIMARY KEY NOT NULL REFERENCES account (address),
    networks jsonb NOT NULL
);

