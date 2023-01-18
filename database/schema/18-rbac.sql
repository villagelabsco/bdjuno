CREATE TABLE rbac_authorizations (
    index text primary key,
    messages jsonb,
    metadata text,
    group_id numeric,
    role_admins jsonb,
    role_delegates jsonb
)