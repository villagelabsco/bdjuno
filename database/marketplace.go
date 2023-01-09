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
	marketplacetypes "github.com/villagelabs/villaged/x/marketplace/types"
)

func (db *Db) Listing(network, index string) (*marketplacetypes.Listing, error) {
	q := `
		SELECT (network, index, product, nft, attributes, creator, active) FROM marketplace_listings
		WHERE network = $1 AND index = $2
	`

	var listing marketplacetypes.Listing
	if err := db.Sqlx.Select(&listing, q, network, index); err != nil {
		return nil, fmt.Errorf("error while getting listing: %s", err)
	}

	return &listing, nil
}

func (db *Db) InsertListing(lst *marketplacetypes.Listing) error {
	stmt := `
		INSERT INTO marketplace_listings (network, index, product, nft, attributes, creator, active) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	if _, err := db.Sql.Exec(stmt,
		lst.Network,
		lst.Index,
		lst.Product,
		lst.Nft,
		lst.Attributes,
		lst.Creator,
		lst.Active); err != nil {
		return fmt.Errorf("error while storing listing: %s", err)
	}

	return nil
}

func (db *Db) UpdateListing(lst *marketplacetypes.Listing) error {
	stmt := `
		UPDATE marketplace_listings 
		SET 
			product = $1,
			attributes = $2
		WHERE network = $3 AND index = $4
	`

	if _, err := db.Sql.Exec(stmt,
		lst.Product,
		lst.Attributes,
		lst.Network,
		lst.Index); err != nil {
		return fmt.Errorf("error while updating listing: %s", err)
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

	if _, err := db.Sql.Exec(stmt,
		active,
		network,
		index); err != nil {
		return fmt.Errorf("error while updating listing: %s", err)
	}

	return nil
}

func (db *Db) Order(network, index string) (*marketplacetypes.Order, error) {
	q := `
		SELECT (network, creator, index, attributes, items) FROM marketplace_orders
		WHERE network = $1 AND index = $2
	`

	var order marketplacetypes.Order
	if err := db.Sqlx.Select(&order, q, network, index); err != nil {
		return nil, fmt.Errorf("error while getting order: %s", err)
	}

	return &order, nil
}

func (db *Db) InsertOrder(od *marketplacetypes.Order) error {
	stmt := `
		INSERT INTO marketplace_orders (network, creator, index, attributes, items)
		VALUES ($1, $2, $3, $5, $6)
	`

	if _, err := db.Sql.Exec(stmt,
		od.Network,
		od.Creator,
		od.Index,
		od.Attributes,
		od.Items); err != nil {
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

	if _, err := db.Sql.Exec(stmt,
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

	if _, err := db.Sql.Exec(stmt,
		network,
		index); err != nil {
		return fmt.Errorf("error while deleting order: %s", err)
	}

	return nil
}
