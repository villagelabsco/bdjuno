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

package reputation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/bdjuno/v3/database/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/pkg/errors"
	reputationtypes "github.com/villagelabs/villaged/x/reputation/types"
	"time"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *reputationtypes.MsgPostFeedback:
		return m.HandleMsgPostFeedback(cosmosMsg)
	}

	return nil
}

func (m *Module) HandleMsgPostFeedback(msg *reputationtypes.MsgPostFeedback) error {
	err := m.db.SavePostFeedback(msg)
	if err != nil {
		return errors.Wrap(err, "error while saving reputation post feedback")
	}

	fbItem := &types.ReputationFeedbackItem{
		CreatorAcc: msg.Creator,
		DestAcc:    msg.DstAccount,
		TxId:       msg.TxId,
		Ref:        msg.Ref,
		// TODO: Any way to get a block time?
		Timestamp: time.Now(),
	}

	fb, err := m.db.FeedbackAggregate(msg.DstAccount)
	if err != nil {
		return errors.Wrap(err, "error while getting reputation feedback aggregate")
	}

	switch reputationtypes.FeedbackType(msg.FbType) {
	case reputationtypes.PositiveFeedback:
		fb.CptPositive += 1
		fb.Positive = append(fb.Positive, fbItem.ToDto())
	case reputationtypes.NeutralFeedback:
		fb.CptNeutral += 1
		fb.Neutral = append(fb.Neutral, fbItem.ToDto())
	case reputationtypes.NegativeFeedback:
		fb.CptNegative += 1
		fb.Negative = append(fb.Negative, fbItem.ToDto())
	}

	if fb.Index == "" {
		return m.db.InsertFeedbackAggregate(fb)
	} else {
		return m.db.UpdateFeedbackAggregate(fb)
	}
}
