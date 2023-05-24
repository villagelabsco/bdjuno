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
	"github.com/villagelabsco/bdjuno/v4/database/types"
	reputationtypes "github.com/villagelabsco/village/x/reputation/types"
)

func (db *Db) SaveOrUpdateReputationFeedback(fb *reputationtypes.Feedback) error {
	stmt := `
		INSERT INTO reputation_feedback ("index", "cpt_positive", "cpt_negative", "cpt_neutral", "positive", "negative", "neutral", "feedbackers", "last_change")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT ("index") DO UPDATE
		SET
		    last_change = $2,
			cpt_positive = $3,
			cpt_neutral = $4,
			cpt_negative = $5,
			positive = $6,
			neutral = $7,
			negative = $8,
			feedbackers = $9;
	`

	dbRep, err := types.DbReputationFeedback{}.FromProto(fb)
	if err != nil {
		return fmt.Errorf("error while converting reputation feedback aggregate to db model: %s", err)
	}

	_, err = db.SQL.Exec(stmt,
		dbRep.Index,
		dbRep.LastChange,
		dbRep.CptPositive,
		dbRep.CptNeutral,
		dbRep.CptNegative,
		dbRep.Positive,
		dbRep.Neutral,
		dbRep.Negative,
		dbRep.Feedbackers,
	)
	if err != nil {
		return fmt.Errorf("error while saving reputation feedback aggregate: %s", err)
	}

	return nil
}
