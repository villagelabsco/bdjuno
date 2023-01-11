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

package marketplace

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	marketplacetypes "github.com/villagelabs/villaged/x/marketplace/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := msg.(type) {
	case *marketplacetypes.MsgCreateListing:
		return m.HandleMsgCreateListing(tx.Height, cosmosMsg)
	case *marketplacetypes.MsgUpdateListing:
		return m.HandleMsgUpdateListing(tx.Height, cosmosMsg)
	case *marketplacetypes.MsgDeleteListing:
		return m.HandleMsgDeleteListing(tx.Height, cosmosMsg)
	case *marketplacetypes.MsgCreateOrder:
		return m.HandleMsgCreateOrder(tx.Height, cosmosMsg)
	case *marketplacetypes.MsgUpdateOrder:
		return m.HandleMsgUpdateOrder(tx.Height, cosmosMsg)
	case *marketplacetypes.MsgDeleteOrder:
		return m.HandleMsgDeleteOrder(tx.Height, cosmosMsg)
	default:
		return fmt.Errorf("unrecognized marketplace message type: %T", msg)
	}
}

func (m *Module) HandleMsgCreateListing(height int64, msg *marketplacetypes.MsgCreateListing) error {
	lst, err := m.src.GetListing(height, marketplacetypes.QueryGetListingRequest{
		Network: msg.Network,
		Index:   msg.Index,
	})
	if err != nil {
		return fmt.Errorf("error while handling create listing msg: %s", err)
	}
	listing := lst.Listing

	if err := m.db.SaveListing(&listing); err != nil {
		return fmt.Errorf("error while handling create listing msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgUpdateListing(height int64, msg *marketplacetypes.MsgUpdateListing) error {
	lst, err := m.src.GetListing(height, marketplacetypes.QueryGetListingRequest{
		Network: msg.Network,
		Index:   msg.Index,
	})
	if err != nil {
		return fmt.Errorf("error while handling update listing msg: %s", err)
	}
	listing := lst.Listing

	if err := m.db.UpdateListing(&listing); err != nil {
		return fmt.Errorf("error while handling update listing msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgDeleteListing(height int64, msg *marketplacetypes.MsgDeleteListing) error {
	if err := m.db.UpdateListingActive(msg.Network, msg.Index, false); err != nil {
		return fmt.Errorf("error while handling delete listing msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgCreateOrder(height int64, msg *marketplacetypes.MsgCreateOrder) error {
	ord, err := m.src.GetOrder(height, marketplacetypes.QueryGetOrderRequest{
		Network: msg.Network,
		Index:   msg.Index,
	})
	if err != nil {
		return fmt.Errorf("error while handling create order msg: %s", err)
	}
	order := ord.Order

	if err := m.db.InsertOrder(&order); err != nil {
		return fmt.Errorf("error while handling create order msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgUpdateOrder(height int64, msg *marketplacetypes.MsgUpdateOrder) error {
	ord, err := m.src.GetOrder(height, marketplacetypes.QueryGetOrderRequest{
		Network: msg.Network,
		Index:   msg.Index,
	})
	if err != nil {
		return fmt.Errorf("error while handling update order msg: %s", err)
	}
	order := ord.Order

	if err := m.db.UpdateOrder(&order); err != nil {
		return fmt.Errorf("error while handling update order msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgDeleteOrder(height int64, msg *marketplacetypes.MsgDeleteOrder) error {
	if err := m.db.DeleteOrder(msg.Network, msg.Index); err != nil {
		return fmt.Errorf("error while handling delete order msg: %s", err)
	}

	return nil
}
