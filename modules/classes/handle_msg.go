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

package classes

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	classestypes "github.com/villagelabs/villaged/x/classes/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *classestypes.MsgCreateClass:
		return m.HandleMsgCreateClass(tx.Height, cosmosMsg)
	case *classestypes.MsgDeleteClass:
		return m.HandleMsgDeleteClass(tx.Height, cosmosMsg)
	case *classestypes.MsgUpdateClass:
		return m.HandleMsgUpdateClass(tx.Height, cosmosMsg)
	default:
		return fmt.Errorf("unrecognized classes message type: %T", msg)
	}
}

func (m *Module) HandleMsgCreateClass(height int64, msg *classestypes.MsgCreateClass) error {
	cl, err := m.src.GetClass(height, classestypes.QueryGetClassRequest{
		NetworkID: msg.NetworkID,
		Index:     msg.Id,
	})
	if err != nil {
		return fmt.Errorf("error while getting class: %s", err)
	}
	class := cl.Class

	if err := m.db.InsertClass(&class); err != nil {
		return fmt.Errorf("error while inserting class: %s", err)
	}

	if cl.Class.Parent != "" {
		if err := m.db.UpdateClassHasChildren(class.NetworkID, class.Parent, true); err != nil {
			return fmt.Errorf("error while setting has children on parent: %s", err)
		}
	}

	return nil
}

func (m *Module) HandleMsgDeleteClass(height int64, msg *classestypes.MsgDeleteClass) error {
	if err := m.db.DeleteClass(msg.NetworkID, msg.Index); err != nil {
		return fmt.Errorf("error while deleting class: %s", err)
	}

	cl, err := m.db.Class(msg.NetworkID, msg.Index)
	if err != nil {
		return fmt.Errorf("error while getting class: %s", err)
	}

	if cl.Parent != "" {
		parent, err := m.src.GetClass(height, classestypes.QueryGetClassRequest{
			NetworkID: msg.NetworkID,
			Index:     cl.Parent,
		})
		if err != nil {
			return fmt.Errorf("error while getting parent class: %s", err)
		}
		if !parent.Class.HasChildren {
			if err := m.db.UpdateClassHasChildren(msg.NetworkID, parent.Class.Index, false); err != nil {
				return fmt.Errorf("error while setting has children on parent: %s", err)
			}
		}
	}

	return nil
}

func (m *Module) HandleMsgUpdateClass(height int64, msg *classestypes.MsgUpdateClass) error {
	cl, err := m.src.GetClass(height, classestypes.QueryGetClassRequest{
		NetworkID: msg.NetworkID,
		Index:     msg.Index,
	})
	if err != nil {
		return fmt.Errorf("error while handling msg update class: %s", err)
	}

	if err := m.db.UpdateClass(&cl.Class); err != nil {
		return fmt.Errorf("error while handling msg update class: %s", err)
	}

	return nil
}
