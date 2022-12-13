CREATE TABLE kyc_identity_provider (
   index TEXT PRIMARY KEY NOT NULL,
   admin_accounts JSON[] NOT NULL,
   provider_accounts JSON[] NOT NULL
);