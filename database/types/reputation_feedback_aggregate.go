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

package types

import (
	"encoding/json"
	"fmt"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	reputationtypes "github.com/villagelabsco/villaged/x/reputation/types"
)

type DbReputationFeedbackAggregate struct {
	Index       string             `db:"index"`
	CptPositive uint64             `db:"cpt_positive"`
	CptNegative uint64             `db:"cpt_negative"`
	CptNeutral  uint64             `db:"cpt_neutral"`
	Positive    sqlxtypes.JSONText `db:"positive"`
	Negative    sqlxtypes.JSONText `db:"negative"`
	Neutral     sqlxtypes.JSONText `db:"neutral"`
	Feedbackers sqlxtypes.JSONText `db:"feedbackers"`
	LastChange  uint64             `db:"last_change"`
}

func (fa DbReputationFeedbackAggregate) FromProto(f *reputationtypes.Feedback) (DbReputationFeedbackAggregate, error) {
	pos, err := json.Marshal(f.Positive)
	if err != nil {
		return DbReputationFeedbackAggregate{}, fmt.Errorf("error marshalling positive: %v", err)
	}
	neg, err := json.Marshal(f.Negative)
	if err != nil {
		return DbReputationFeedbackAggregate{}, fmt.Errorf("error marshalling negative: %v", err)
	}
	neu, err := json.Marshal(f.Neutral)
	if err != nil {
		return DbReputationFeedbackAggregate{}, fmt.Errorf("error marshalling neutral: %v", err)
	}
	feedbackers, err := json.Marshal(f.Feedbackers)
	if err != nil {
		return DbReputationFeedbackAggregate{}, fmt.Errorf("error marshalling feedbackers: %v", err)
	}

	return DbReputationFeedbackAggregate{
		Index:       f.Index,
		CptPositive: f.CptPositive,
		CptNegative: f.CptNegative,
		CptNeutral:  f.CptNeutral,
		Positive:    pos,
		Negative:    neg,
		Neutral:     neu,
		Feedbackers: feedbackers,
		LastChange:  f.LastChange,
	}, nil
}

func (fa DbReputationFeedbackAggregate) ToProto() (*reputationtypes.Feedback, error) {
	var pos, neg, neu []*reputationtypes.FeedbackItem
	var feedbackers map[string]uint64

	if err := json.Unmarshal(fa.Positive, &pos); err != nil {
		return nil, fmt.Errorf("error unmarshalling positive: %v", err)
	}
	if err := json.Unmarshal(fa.Negative, &neg); err != nil {
		return nil, fmt.Errorf("error unmarshalling negative: %v", err)
	}
	if err := json.Unmarshal(fa.Neutral, &neu); err != nil {
		return nil, fmt.Errorf("error unmarshalling neutral: %v", err)
	}
	if err := json.Unmarshal(fa.Feedbackers, &feedbackers); err != nil {
		return nil, fmt.Errorf("error unmarshalling feedbackers: %v", err)
	}

	return &reputationtypes.Feedback{
		Index:       fa.Index,
		LastChange:  fa.LastChange,
		CptPositive: fa.CptPositive,
		CptNeutral:  fa.CptNeutral,
		CptNegative: fa.CptNegative,
		Positive:    pos,
		Neutral:     neu,
		Negative:    neg,
		Feedbackers: feedbackers,
	}, nil
}
