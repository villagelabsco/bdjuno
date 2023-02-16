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
	"time"
)

type ReputationFeedbackAggregate struct {
	Index       string                   `db:"index"`
	LastChange  uint64                   `db:"last_change"`
	CptPositive uint64                   `db:"cpt_positive"`
	CptNeutral  uint64                   `db:"cpt_neutral"`
	CptNegative uint64                   `db:"cpt_negative"`
	Positive    []ReputationFeedbackItem `db:"positive"`
	Neutral     []ReputationFeedbackItem `db:"neutral"`
	Negative    []ReputationFeedbackItem `db:"negative"`
	Feedbackers map[string]uint64        `db:"feedbackers"`
}

type ReputationFeedbackItem struct {
	CreatorAcc string    `db:"creator_acc"`
	DestAcc    string    `db:"dest_acc"`
	TxId       string    `db:"tx_id"`
	Ref        string    `db:"ref"`
	Timestamp  time.Time `db:"timestamp"`
}

type ReputationPostFeedback struct {
	Creator    string    `db:"creator"`
	Network    string    `db:"network"`
	FbType     uint64    `db:"fb_type"`
	DstAccount string    `db:"dst_account"`
	TxId       string    `db:"tx_id"`
	Ref        string    `db:"ref"`
	Timestamp  time.Time `db:"timestamp"`
}
