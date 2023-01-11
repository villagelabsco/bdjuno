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

package products

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v3/types"
	productstypes "github.com/villagelabs/villaged/x/products/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := (msg).(type) {
	case *productstypes.MsgCreateProduct:
		return m.HandleMsgCreateProduct(tx.Height, cosmosMsg)
	case *productstypes.MsgUpdateProduct:
		return m.HandleMsgUpdateProduct(tx.Height, cosmosMsg)
	case *productstypes.MsgDeleteProduct:
		return m.HandleMsgDeleteProduct(tx.Height, cosmosMsg)
	case *productstypes.MsgCreateProductClass:
		return m.HandleMsgCreateProductClass(tx.Height, cosmosMsg)
	case *productstypes.MsgCreateTaskClass:
		return m.HandleMsgCreateTaskClass(tx.Height, cosmosMsg)
	case *productstypes.MsgFreezeClass:
		return m.HandleMsgFreezeClass(tx.Height, cosmosMsg)
	case *productstypes.MsgUpdateClass:
		return m.HandleMsgUpdateClass(tx.Height, cosmosMsg)
	default:
		return fmt.Errorf("unrecognized products message type: %T", msg)
	}
}

func (m *Module) HandleMsgCreateProduct(height int64, msg *productstypes.MsgCreateProduct) error {
	// TODO: How do we handle with new nft-first product model?
	return nil
}

func (m *Module) HandleMsgUpdateProduct(height int64, msg *productstypes.MsgUpdateProduct) error {
	// TODO: How do we handle with new nft-first product model?
	return nil
}

func (m *Module) HandleMsgDeleteProduct(height int64, msg *productstypes.MsgDeleteProduct) error {
	// TODO: How do we handle with new nft-first product model?
	return nil
}

func (m *Module) HandleMsgCreateProductClass(height int64, msg *productstypes.MsgCreateProductClass) error {
	return nil
}

func (m *Module) HandleMsgCreateTaskClass(height int64, msg *productstypes.MsgCreateTaskClass) error {
	return nil
}

func (m *Module) HandleMsgFreezeClass(height int64, msg *productstypes.MsgFreezeClass) error {
	return nil
}

func (m *Module) HandleMsgUpdateClass(height int64, msg *productstypes.MsgUpdateClass) error {
	return nil
}
