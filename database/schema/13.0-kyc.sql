CREATE TABLE kyc_identity_provider (
   index TEXT NOT NULL,
   admin_accounts jsonb NOT NULL,
   provider_accounts jsonb NOT NULL,
   asset_minter_accounts jsonb NOT NULL,
   asset_burner_accounts jsonb NOT NULL,
   primary key (index)
);