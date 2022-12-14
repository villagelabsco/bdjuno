/*
 * Copyright 2022 LimeChain Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package database

import (
	"fmt"
	"github.com/forbole/bdjuno/v3/database/types"
	reputationtypes "github.com/villagelabs/villaged/x/reputation/types"
)

func (db *Db) FeedbackAggregate(index string) (*reputationtypes.Feedback, error) {
	q := `
	SELECT (rfa.index, rfa.last_change, rfa.cpt_positive, rfa.cpt_neutral, rfa.cpt_negative, rfa.positive, rfa.neutral, rfa.negative, rfa.feedbackers)
	FROM reputation_feedback_aggregate AS rfa
	WHERE rfa.index = $1;`

	var fb types.ReputationFeedbackAggregate
	err := db.Sqlx.Select(&fb, q, index)
	if err != nil {
		return nil, fmt.Errorf("error while querying reputation feedback aggregate: %s", err)
	}

	return fb.ToDto(), nil
}

func (db *Db) InsertFeedbackAggregate(fb *reputationtypes.Feedback) error {
	q := `
	INSERT INTO reputation_feedback_aggregate ("index", "cpt_positive", "cpt_negative", "cpt_neutral", "positive", "negative", "neutral", "feedbackers", "last_change")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	_, err := db.Sql.Exec(q, fb.Index, fb.CptPositive, fb.CptNegative, fb.CptNeutral, fb.Positive, fb.Negative, fb.Neutral, fb.Feedbackers, fb.LastChange)
	if err != nil {
		return fmt.Errorf("error while inserting reputation feedback aggregate: %s", err)
	}

	return nil
}

func (db *Db) UpdateFeedbackAggregate(fb *reputationtypes.Feedback) error {
	stmt := `
	UPDATE reputation_feedback_aggregate AS rfa
	SET last_change = $2,
		cpt_positive = $3,
		cpt_neutral = $4,
		cpt_negative = $5,
		positive = $6,
		neutral = $7,
		negative = $8,
		feedbackers = $9
    WHERE rfa.index = $1`

	_, err := db.Sql.Exec(stmt,
		fb.Index,
		fb.LastChange,
		fb.CptPositive,
		fb.CptNeutral,
		fb.CptNegative,
		fb.Positive,
		fb.Neutral,
		fb.Negative,
		fb.Feedbackers,
	)
	if err != nil {
		return fmt.Errorf("error while updating reputation feedback aggregate: %s", err)
	}

	return nil
}

func (db *Db) SavePostFeedback(msg *reputationtypes.MsgPostFeedback) error {
	stmt := `
	INSERT INTO reputation_feedback ("creator", "network", "fb_type", "dst_account", "tx_id", "ref")
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Sql.Exec(stmt, msg.Creator, msg.Network, msg.FbType, msg.DstAccount, msg.TxId, msg.Ref)
	if err != nil {
		return fmt.Errorf("error while saving reputation feedback: %s", err)
	}

	return nil
}
