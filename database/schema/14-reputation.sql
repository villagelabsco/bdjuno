CREATE TABLE reputation_feedback (
    index TEXT NOT NULL,
    cpt_positive INTEGER NOT NULL,
    cpt_negative INTEGER NOT NULL,
    cpt_neutral INTEGER NOT NULL,
    positive jsonb NOT NULL,
    negative jsonb NOT NULL,
    neutral jsonb NOT NULL,
    feedbackers jsonb NOT NULL,
    last_change TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    primary key (index)
);
