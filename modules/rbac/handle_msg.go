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

package rbac

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v4/types"
	rbactypes "github.com/villagelabsco/villaged/x/rbac/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *rbactypes.MsgDeclareRole:
		return m.HandleMsgDeclareRole(index, tx, cosmosMsg)
	case *rbactypes.MsgUpdateGroupMembers:
		return m.HandleMsgUpdateGroupMembers(index, tx, cosmosMsg)
	case *rbactypes.MsgTransferRoleOwnership:
		return m.HandleMsgTransferRoleOwnership(index, tx, cosmosMsg)
	case *rbactypes.MsgUpdateRole:
		return m.HandleMsgUpdateRole(index, tx, cosmosMsg)
	case *rbactypes.MsgSetRoleDelegates:
		return m.HandleMsgSetRoleDelegates(index, tx, cosmosMsg)
	default:
		return fmt.Errorf("unrecognized rbac message type: %T", msg)
	}
}

func (m *Module) HandleMsgDeclareRole(index int, height *juno.Tx, msg *rbactypes.MsgDeclareRole) error {
	return nil
}

func (m *Module) HandleMsgUpdateGroupMembers(index int, height *juno.Tx, msg *rbactypes.MsgUpdateGroupMembers) error {
	return nil
}

func (m *Module) HandleMsgTransferRoleOwnership(index int, height *juno.Tx, msg *rbactypes.MsgTransferRoleOwnership) error {
	return nil
}

func (m *Module) HandleMsgUpdateRole(index int, height *juno.Tx, msg *rbactypes.MsgUpdateRole) error {
	return nil
}

func (m *Module) HandleMsgSetRoleDelegates(index int, height *juno.Tx, msg *rbactypes.MsgSetRoleDelegates) error {
	return nil
}
