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

package database

import (
	"fmt"
	"github.com/villagelabsco/bdjuno/v3/database/types"
	marketplacetypes "github.com/villagelabsco/villaged/x/marketplace/types"
)

func (db *Db) SaveOrUpdateListing(listing *marketplacetypes.Listing) error {
	stmt := `
		INSERT INTO marketplace_listings (network, index, reference, product_class_id, product_nft_id, attributes, creator, active) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (network, index) 
		    DO UPDATE SET
			reference = $3,
			attributes = $6,
			active = $8;
			
	`

	lst, err := types.DbMarketplaceListing{}.FromProto(listing)
	if err != nil {
		return fmt.Errorf("error while converting to db listing: %s", err)
	}
	if _, err := db.SQL.Exec(stmt,
		lst.Network,
		lst.Index,
		lst.Reference,
		lst.ProductClassId,
		lst.ProductNftId,
		lst.Attributes,
		lst.Creator,
		lst.Active); err != nil {
		return fmt.Errorf("error while storing listing: %s", err)
	}

	return nil
}

func (db *Db) UpdateListingActive(network, index string, active bool) error {
	stmt := `
		UPDATE marketplace_listings 
		SET 
			active = $1
		WHERE network = $2 AND index = $3
	`

	if _, err := db.SQL.Exec(stmt,
		active,
		network,
		index); err != nil {
		return fmt.Errorf("error while updating listing: %s", err)
	}

	return nil
}

func (db *Db) InsertOrder(ord *marketplacetypes.Order) error {
	stmt := `
		INSERT INTO marketplace_orders (network, index, status, timestamp, creator, attributes, items, total)
		VALUES ($1, $2, $3, $5, $6, $7, $8, $9)
	`

	o, err := types.DbMarketplaceOrder{}.FromProto(ord)
	if err != nil {
		return fmt.Errorf("error while converting to db order: %s", err)
	}

	if _, err := db.SQL.Exec(stmt,
		o.Network,
		o.Index,
		o.Status,
		o.Timestamp,
		o.Creator,
		o.Attributes,
		o.Items,
		o.Total); err != nil {
		return fmt.Errorf("error while storing order: %s", err)
	}

	return nil
}

func (db *Db) UpdateOrder(od *marketplacetypes.Order) error {
	stmt := `
		UPDATE marketplace_orders 
		SET 
			status = $1
		WHERE network = $2 AND index = $3
	`

	if _, err := db.SQL.Exec(stmt,
		od.Status,
		od.Network,
		od.Index); err != nil {
		return fmt.Errorf("error while updating order: %s", err)
	}

	return nil
}

func (db *Db) DeleteOrder(network, index string) error {
	stmt := `
		DELETE FROM marketplace_orders 
		WHERE network = $1 AND index = $2
	`

	if _, err := db.SQL.Exec(stmt,
		network,
		index); err != nil {
		return fmt.Errorf("error while deleting order: %s", err)
	}

	return nil
}
