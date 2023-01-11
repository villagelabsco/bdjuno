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
	productstypes "github.com/villagelabs/villaged/x/products/types"
)

func (db *Db) DeleteProduct(network, index string) error {
	stmt := `
		UPDATE products_products
		SET
			active = false
		WHERE network = $1 AND index = $2
	`

	if _, err := db.Sql.Exec(stmt, network, index); err != nil {
		return fmt.Errorf("error while setting product as inactive: %s", err)
	}

	return nil
}

func (db *Db) UpdateProductHasChildren(network, index string, val bool) error {
	stmt := `
	UPDATE products_products
	SET
	    has_children = $1
	WHERE network = $2 AND index = $3
	`

	if _, err := db.Sql.Exec(stmt, val, network, index); err != nil {
		return fmt.Errorf("error while updating product has_children: %s", err)
	}

	return nil
}

func (db *Db) UpdateProductActive(network, index string, val bool) error {
	stmt := `
	UPDATE products_products
	SET
	    active = $1
	WHERE network = $2 AND index = $3
	`

	if _, err := db.Sql.Exec(stmt, val, network, index); err != nil {
		return fmt.Errorf("error while updating product active: %s", err)
	}

	return nil
}

func (db *Db) SaveProductClass(pc productstypes.ProductClassInfo) error {

	return nil
}
