CREATE TABLE reputation_feedback (
    id SERIAL NOT NULL PRIMARY KEY,
    creator TEXT NOT NULL REFERENCES account (address),
    network TEXT NOT NULL,
    fb_type INTEGER NOT NULL,
    dst_account TEXT NOT NULL REFERENCES account (address),
    tx_id TEXT NOT NULL,
    ref TEXT NOT NULL
);

CREATE TABLE reputation_feedback_aggregate (
    index TEXT PRIMARY KEY NOT NULL,
    cpt_positive INTEGER NOT NULL,
    cpt_negative INTEGER NOT NULL,
    cpt_neutral INTEGER NOT NULL,
    positive JSON[] NOT NULL,
    negative JSON[] NOT NULL,
    neutral JSON[] NOT NULL,
    feedbackers JSON NOT NULL,
    last_change TIMESTAMP NOT NULL
);
