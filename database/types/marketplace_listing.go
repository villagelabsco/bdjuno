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

package types

import (
	"encoding/base64"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	marketplacetypes "github.com/villagelabs/villaged/x/marketplace/types"
)

type DbMarketplaceListing struct {
	Network        string             `db:"network"`
	Index          string             `db:"index"`
	Reference      string             `db:"reference"`
	ProductClassId string             `db:"product_class_id"`
	ProductNftId   string             `db:"product_nft_id"`
	Attributes     sqlxtypes.JSONText `db:"attributes"`
	Creator        string             `db:"creator"`
	Active         bool               `db:"active"`
}

func (m *DbMarketplaceListing) ToProto() *marketplacetypes.Listing {
	return &marketplacetypes.Listing{
		Index:          m.Index,
		Network:        m.Network,
		Reference:      m.Reference,
		ProductClassId: m.ProductClassId,
		ProductNftId:   m.ProductNftId,
		Attributes:     base64.StdEncoding.EncodeToString(m.Attributes),
		Creator:        m.Creator,
		Active:         m.Active,
	}
}

func (DbMarketplaceListing) FromProto(l *marketplacetypes.Listing) (DbMarketplaceListing, error) {
	attrs, err := base64.StdEncoding.DecodeString(l.Attributes)
	if err != nil {
		return DbMarketplaceListing{}, err
	}

	return DbMarketplaceListing{
		Network:        l.Network,
		Index:          l.Index,
		Reference:      l.Reference,
		ProductClassId: l.ProductClassId,
		ProductNftId:   l.ProductNftId,
		Attributes:     attrs,
		Creator:        l.Creator,
		Active:         l.Active,
	}, nil
}
