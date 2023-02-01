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
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	"github.com/pkg/errors"
	reputationtypes "github.com/villagelabsco/villaged/x/reputation/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch cosmosMsg := msg.(type) {
	case *reputationtypes.MsgPostFeedback:
		return m.HandleMsgPostFeedback(index, tx, cosmosMsg)
	}

	return nil
}

func (m *Module) HandleMsgPostFeedback(index int, tx *juno.Tx, msg *reputationtypes.MsgPostFeedback) error {
	err := m.db.SavePostFeedback(msg)
	if err != nil {
		return errors.Wrap(err, "error while saving reputation post feedback")
	}

	fb, err := m.s.GetFeedback(tx.Height, reputationtypes.QueryGetFeedbackRequest{
		Network: msg.Network,
		Index:   msg.DstAccount,
	})
	if err != nil {
		return fmt.Errorf("error while getting feedback: %s", err)
	}

	existing, err := m.db.FeedbackAggregate(msg.DstAccount)
	if err != nil {
		return fmt.Errorf("error while getting feedback aggregate: %s", err)
	}
	if existing == nil {
		if err := m.db.SaveFeedbackAggregate(&fb.Feedback); err != nil {
			return fmt.Errorf("error while inserting feedback aggregate: %s", err)
		}
	} else {
		if err := m.db.UpdateFeedbackAggregate(&fb.Feedback); err != nil {
			return fmt.Errorf("error while updating feedback aggregate: %s", err)
		}
	}

	return nil
}
