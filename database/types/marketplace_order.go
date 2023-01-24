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
	"encoding/json"
	"fmt"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	marketplacetypes "github.com/villagelabs/villaged/x/marketplace/types"
	"time"
)

type DbMarketplaceOrder struct {
	Network    string    `db:"network"`
	Index      string    `db:"index"`
	Status     int64     `db:"status"`
	Timestamp  time.Time `db:"timestamp"`
	Creator    string    `db:"creator"`
	Attributes sqlxtypes.JSONText
	Items      sqlxtypes.JSONText
	Total      sqlxtypes.JSONText
}

func (DbMarketplaceOrder) FromProto(ord *marketplacetypes.Order) (DbMarketplaceOrder, error) {
	attrs, err := base64.StdEncoding.DecodeString(ord.Attributes)
	if err != nil {
		return DbMarketplaceOrder{}, fmt.Errorf("error decoding order attributes: %s", err)
	}

	items, err := json.Marshal(ord.Items)
	if err != nil {
		return DbMarketplaceOrder{}, fmt.Errorf("error marshalling order items: %s", err)
	}

	total, err := json.Marshal(ord.Total)
	if err != nil {
		return DbMarketplaceOrder{}, fmt.Errorf("error marshalling order total: %s", err)
	}

	return DbMarketplaceOrder{
		Network:    ord.Network,
		Index:      ord.Index,
		Status:     ord.Status,
		Timestamp:  time.Unix(0, ord.Timestamp),
		Creator:    ord.Creator,
		Attributes: attrs,
		Items:      items,
		Total:      total,
	}, nil
}
