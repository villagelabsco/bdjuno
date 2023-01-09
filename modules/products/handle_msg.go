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
	default:
		return fmt.Errorf("unrecognized products message type: %T", msg)
	}
}

func (m *Module) HandleMsgCreateProduct(height int64, msg *productstypes.MsgCreateProduct) error {
	pr, err := m.src.GetProduct(height, productstypes.QueryGetProductRequest{
		Network: msg.Network,
		Index:   msg.Index,
	})
	if err != nil {
		return fmt.Errorf("error while handling create product msg: %s", err)
	}
	product := pr.Product

	if err := m.db.InsertProduct(&product); err != nil {
		return fmt.Errorf("error while handling create product msg: %s", err)
	}

	if product.Parent != "" {
		if err := m.db.UpdateProductHasChildren(product.Network, product.Parent, true); err != nil {
			return fmt.Errorf("error while handling create product msg: %s", err)
		}
	}

	return nil
}

func (m *Module) HandleMsgUpdateProduct(height int64, msg *productstypes.MsgUpdateProduct) error {
	pr, err := m.src.GetProduct(height, productstypes.QueryGetProductRequest{
		Network: msg.Network,
		Index:   msg.Index,
	})
	if err != nil {
		return fmt.Errorf("error while handling update product msg: %s", err)
	}
	product := pr.Product

	if err := m.db.UpdateProduct(&product); err != nil {
		return fmt.Errorf("error while handling update product msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgDeleteProduct(height int64, msg *productstypes.MsgDeleteProduct) error {
	if err := m.db.UpdateProductActive(msg.Network, msg.Index, false); err != nil {
		return fmt.Errorf("error while handling delete product msg: %s", err)
	}

	// Set parent's has_children flag if needed
	pr, err := m.db.Product(msg.Network, msg.Index)
	if err != nil {
		return fmt.Errorf("error while handling delete product msg: %s", err)
	}

	if pr.Parent != "" {
		parent, err := m.src.GetProduct(height, productstypes.QueryGetProductRequest{
			Network: msg.Network,
			Index:   pr.Parent,
		})
		if err != nil {
			return fmt.Errorf("error while handling delete product msg: %s", err)
		}
		if !parent.Product.HasChildren {
			if err := m.db.UpdateProductHasChildren(msg.Network, pr.Parent, false); err != nil {
				return fmt.Errorf("error while handling delete product msg: %s", err)
			}
		}
	}

	return nil
}
