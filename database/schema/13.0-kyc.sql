CREATE TABLE kyc_identity_provider (
   index TEXT PRIMARY KEY NOT NULL,
   admin_accounts JSON[] NOT NULL,
   provider_accounts JSON[] NOT NULL,
   asset_minter_accounts JSON[] NOT NULL,
   asset_burner_accounts JSON[] NOT NULL
);