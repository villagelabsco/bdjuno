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

import reputationtypes "github.com/villagelabs/villaged/x/reputation/types"

func (receiver *ReputationFeedbackAggregate) ToDto() *reputationtypes.Feedback {
	pos := make([]*reputationtypes.FeedbackItem, len(receiver.Positive))
	for i, v := range receiver.Positive {
		pos[i] = v.ToDto()
	}
	neg := make([]*reputationtypes.FeedbackItem, len(receiver.Negative))
	for i, v := range receiver.Negative {
		neg[i] = v.ToDto()
	}
	neu := make([]*reputationtypes.FeedbackItem, len(receiver.Neutral))
	for i, v := range receiver.Neutral {
		neu[i] = v.ToDto()
	}

	return &reputationtypes.Feedback{
		Index:       receiver.Index,
		LastChange:  receiver.LastChange,
		CptPositive: receiver.CptPositive,
		CptNeutral:  receiver.CptNeutral,
		CptNegative: receiver.CptNegative,
		Positive:    pos,
		Neutral:     neu,
		Negative:    neg,
		Feedbackers: receiver.Feedbackers,
	}
}

func (receiver *ReputationFeedbackItem) ToDto() *reputationtypes.FeedbackItem {
	return &reputationtypes.FeedbackItem{
		CreatorAcc: receiver.CreatorAcc,
		DestAcc:    receiver.DestAcc,
		Txid:       receiver.TxId,
		Ref:        receiver.Ref,
		Timestamp:  uint64(receiver.Timestamp.UnixNano()),
	}
}

func (receiver *ReputationPostFeedback) ToDto() *reputationtypes.MsgPostFeedback {
	return &reputationtypes.MsgPostFeedback{
		Creator:    receiver.Creator,
		Network:    receiver.Network,
		FbType:     receiver.FbType,
		DstAccount: receiver.DstAccount,
		Ref:        receiver.Ref,
	}
}
