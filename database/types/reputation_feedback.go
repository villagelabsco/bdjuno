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
	reputationtypes "github.com/villagelabs/villaged/x/reputation/types"
)

type DbReputationFeedback struct {
	Id         uint64                       `db:"id"`
	Creator    string                       `db:"creator"`
	Network    string                       `db:"network"`
	FbType     reputationtypes.FeedbackType `db:"fb_type"`
	DstAccount string                       `db:"dst_account"`
	TxId       string                       `db:"tx_id"`
	Ref        string                       `db:"ref"`
}

func (rf DbReputationFeedback) FromProto(f *reputationtypes.MsgPostFeedback) DbReputationFeedback {
	return DbReputationFeedback{
		Creator:    f.Creator,
		Network:    f.Network,
		FbType:     reputationtypes.FeedbackType(f.FbType),
		DstAccount: f.DstAccount,
		TxId:       f.TxId,
		Ref:        f.Ref,
	}
}
