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

package db_types

import (
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	productstypes "github.com/villagelabs/villaged/x/products/types"
)

type ProductClassInfo struct {
	Network          string                  `db:"network"`
	ClassId          string                  `db:"class_id"`
	FullClassId      string                  `db:"full_class_id"`
	ClassType        productstypes.ClassType `db:"class_type"`
	Name             string                  `db:"name"`
	Description      string                  `db:"description"`
	Metadata         sqlxtypes.JSONText      `db:"metadata"`
	SpecificMetadata sqlxtypes.JSONText      `db:"specific_metadata"`
}

func (p ProductClassInfo) FromProto(info productstypes.ProductClassInfo, nftClass *nfttypes.Class, metadata []byte, specificMetadata []byte) ProductClassInfo {
	return ProductClassInfo{
		Network:          info.Network,
		ClassId:          info.ClassId,
		FullClassId:      info.FullClassId,
		ClassType:        info.ClassType,
		Name:             nftClass.Name,
		Description:      nftClass.Description,
		Metadata:         metadata,
		SpecificMetadata: specificMetadata,
	}
}

func (p ProductClassInfo) ToProto() productstypes.ProductClassInfo {
	return productstypes.ProductClassInfo{
		Network:     p.Network,
		ClassId:     p.ClassId,
		FullClassId: p.FullClassId,
		ClassType:   p.ClassType,
		Name:        p.Name,
	}
}
