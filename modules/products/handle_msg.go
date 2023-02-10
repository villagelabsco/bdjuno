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
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/gogo/protobuf/proto"
	juno "github.com/villagelabsco/juno/v4/types"
	productstypes "github.com/villagelabsco/village/x/products/types"
)

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch cosmosMsg := (msg).(type) {
	case *productstypes.MsgCreateProductClass:
		return m.HandleMsgCreateProductClass(index, tx, cosmosMsg)
	case *productstypes.MsgCreateTaskClass:
		return m.HandleMsgCreateTaskClass(index, tx, cosmosMsg)
	case *productstypes.MsgFreezeClass:
		return m.HandleMsgFreezeClass(index, tx, cosmosMsg)
	case *productstypes.MsgUpdateClass:
		return m.HandleMsgUpdateClass(index, tx, cosmosMsg)
	}

	return nil
}

func (m *Module) HandleMsgCreateProductClass(index int, tx *juno.Tx, msg *productstypes.MsgCreateProductClass) error {
	cl, err := m.src.GetProductClassInfo(tx.Height, productstypes.QueryGetProductClassInfoRequest{
		Network: msg.Network,
		Idx:     msg.Id,
	})
	if err != nil {
		return fmt.Errorf("error while handling create product class info msg: %s", err)
	}
	class := cl.ProductClassInfo

	nftCl, err := m.nftSrc.Class(tx.Height, nfttypes.QueryClassRequest{ClassId: class.FullClassId})
	if err != nil {
		return fmt.Errorf("error while handling create product class info msg: %s", err)
	}
	nftClass := nftCl.Class

	pdata := new(productstypes.ProductClassData)
	metadata, specificMetadata, err := unmarshalProductClassMetadata[*productstypes.ProductClassData](nftClass.Data, pdata)
	if err != nil {
		return fmt.Errorf("error while handling create product class info msg: %s", err)

	}

	if err := m.db.SaveOrUpdateProductClass(class, nftClass, *metadata, **specificMetadata); err != nil {
		return fmt.Errorf("error while handling create product class info msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgCreateTaskClass(index int, tx *juno.Tx, msg *productstypes.MsgCreateTaskClass) error {
	cl, err := m.src.GetProductClassInfo(tx.Height, productstypes.QueryGetProductClassInfoRequest{
		Network: msg.Network,
		Idx:     msg.Id,
	})
	if err != nil {
		return fmt.Errorf("error while handling create task class info msg: %s", err)
	}
	class := cl.ProductClassInfo

	nftCl, err := m.nftSrc.Class(tx.Height, nfttypes.QueryClassRequest{ClassId: class.FullClassId})
	if err != nil {
		return fmt.Errorf("error while handling create task class info msg: %s", err)
	}
	nftClass := nftCl.Class

	pdata := new(productstypes.TaskClassData)
	metadata, specificMetadata, err := unmarshalProductClassMetadata[*productstypes.TaskClassData](nftClass.Data, pdata)
	if err != nil {
		return fmt.Errorf("error while handling create task class info msg: %s", err)
	}

	if err := m.db.SaveOrUpdateTaskClass(class, nftClass, *metadata, **specificMetadata); err != nil {
		return fmt.Errorf("error while handling create task class info msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgFreezeClass(index int, tx *juno.Tx, msg *productstypes.MsgFreezeClass) error {
	cl, err := m.src.GetProductClassInfo(tx.Height, productstypes.QueryGetProductClassInfoRequest{
		Network: msg.Network,
		Idx:     msg.Id,
	})
	if err != nil {
		return fmt.Errorf("error while handling freeze class info msg: %s", err)
	}
	class := cl.ProductClassInfo

	if err := m.db.UpdateProductClassDisabled(class.FullClassId, true); err != nil {
		return fmt.Errorf("error while handling freeze class info msg: %s", err)
	}

	return nil
}

func (m *Module) HandleMsgUpdateClass(index int, tx *juno.Tx, msg *productstypes.MsgUpdateClass) error {
	cl, err := m.src.GetProductClassInfo(tx.Height, productstypes.QueryGetProductClassInfoRequest{
		Network: msg.Network,
		Idx:     msg.Id,
	})
	if err != nil {
		return fmt.Errorf("error while handling update class info msg: %s", err)
	}
	class := cl.ProductClassInfo

	nftCl, err := m.nftSrc.Class(tx.Height, nfttypes.QueryClassRequest{ClassId: class.FullClassId})
	if err != nil {
		return fmt.Errorf("error while handling update class info msg: %s", err)
	}
	nftClass := nftCl.Class

	switch class.ClassType {
	case productstypes.ClassType_CLASS_TYPE_PRODUCT:
		pdata := new(productstypes.ProductClassData)
		metadata, specificMetadata, err := unmarshalProductClassMetadata[*productstypes.ProductClassData](nftClass.Data, pdata)
		if err != nil {
			return fmt.Errorf("error while handling update class info msg: %s", err)
		}
		if err := m.db.SaveOrUpdateProductClass(class, nftClass, *metadata, **specificMetadata); err != nil {
			return fmt.Errorf("error while handling update class info msg: %s", err)
		}
	case productstypes.ClassType_CLASS_TYPE_TASK:
		pdata := new(productstypes.TaskClassData)
		metadata, specificMetadata, err := unmarshalProductClassMetadata[*productstypes.TaskClassData](nftClass.Data, pdata)
		if err != nil {
			return fmt.Errorf("error while handling update class info msg: %s", err)
		}
		if err := m.db.SaveOrUpdateTaskClass(class, nftClass, *metadata, **specificMetadata); err != nil {
			return fmt.Errorf("error while handling update class info msg: %s", err)
		}
	case productstypes.ClassType_CLASS_TYPE_SHIFT:
		return fmt.Errorf("shift class type not supported")
	case productstypes.ClassType_CLASS_TYPE_SERVICE:
		return fmt.Errorf("service class type not supported")
	default:
		return fmt.Errorf("unrecognized class type: %s", class.ClassType)
	}

	return nil
}

func unmarshalProductClassMetadata[T proto.Unmarshaler](data *types.Any, spData T) (*productstypes.StdClassData, *T, error) {
	stdData := new(productstypes.StdClassData)
	if len(data.Value) > 0 {
		if err := stdData.Unmarshal(data.Value); err != nil {
			return nil, nil, fmt.Errorf("error while unmarshaling std product class data: %s", err)
		}
	}

	if stdData.SpecificMetadata != nil {
		if err := spData.Unmarshal(stdData.SpecificMetadata.Value); err != nil {
			return nil, nil, fmt.Errorf("error while unmarshaling specific product class data: %s", err)
		}
	}

	return stdData, &spData, nil
}
