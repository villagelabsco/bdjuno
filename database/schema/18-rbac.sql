CREATE TABLE rbac_authorizations (
    index text not null,
    messages jsonb not null,
    metadata text not null,
    group_id numeric not null,
    role_admins jsonb not null,
    role_delegates jsonb not null,
    primary key (index)
)