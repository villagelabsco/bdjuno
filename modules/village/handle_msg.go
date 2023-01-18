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

package village

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	villagetypes "github.com/villagelabs/villaged/x/village/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *villagetypes.MsgCreateNetwork:
		return m.HandleMsgCreateNetwork(index, tx, cosmosMsg)
	}

	return nil
}

func (m *Module) HandleMsgCreateNetwork(index int, tx *juno.Tx, msg *villagetypes.MsgCreateNetwork) error {
	return m.db.SaveNetwork(&villagetypes.Network{
		Index:            msg.ShortName,
		Active:           true,
		FullName:         msg.FullName,
		IdentityProvider: msg.IdentityProvider,
		InviteOnly:       msg.InviteOnly,
	})
}
