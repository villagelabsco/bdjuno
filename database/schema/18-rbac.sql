CREATE TABLE rbac_authorizations (
    index text primary key,
    messages json[],
    metadata text,
    group_id text,
    role_admins json,
    role_delegates json
)