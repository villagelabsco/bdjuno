CREATE TABLE reputation_feedback (
    id SERIAL NOT NULL PRIMARY KEY,
    creator TEXT NOT NULL REFERENCES account (address),
    network TEXT NOT NULL,
    fb_type INTEGER NOT NULL,
    dst_account TEXT NOT NULL REFERENCES account (address),
    tx_id TEXT NOT NULL,
    ref TEXT NOT NULL,
    primary key (id)
);

CREATE TABLE reputation_feedback_aggregate (
    index TEXT NOT NULL,
    cpt_positive INTEGER NOT NULL,
    cpt_negative INTEGER NOT NULL,
    cpt_neutral INTEGER NOT NULL,
    positive jsonb NOT NULL,
    negative jsonb NOT NULL,
    neutral jsonb NOT NULL,
    feedbackers jsonb NOT NULL,
    last_change TIMESTAMP NOT NULL,
    primary key (index)
);
