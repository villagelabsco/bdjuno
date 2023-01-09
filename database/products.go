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

func (db *Db) InsertProduct(pr *productstypes.Product) error {
	stmt := `
		INSERT INTO products_products (network, index, parent, parent_chain, has_children, name, description, attributes, images, tags, p_type, class, creator, active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`

	if _, err := db.Sql.Exec(stmt,
		pr.Network,
		pr.Index,
		pr.Parent,
		pr.ParentChain,
		pr.HasChildren,
		pr.Name,
		pr.Description,
		pr.Attributes,
		pr.Images,
		pr.Tags,
		pr.Ptype,
		pr.ClassID,
		pr.Creator,
		pr.Active); err != nil {
		return fmt.Errorf("error while inserting product: %s", err)
	}

	return nil
}

func (db *Db) UpdateProduct(pr *productstypes.Product) error {
	stmt := `
		UPDATE products_products
		SET
		    name = $1,
		    description = $2,
		    attributes = $3,
		    images = $4,
		    tags = $5,
		    p_type = $6,
		    class = $7
		WHERE network = $8 AND index = $9
	`

	if _, err := db.Sql.Exec(stmt,
		pr.Name,
		pr.Description,
		pr.Attributes,
		pr.Images,
		pr.Tags,
		pr.Ptype,
		pr.ClassID,
		pr.Network,
		pr.Index); err != nil {
		return fmt.Errorf("error while updating product: %s", err)
	}

	return nil
}

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

func (db *Db) Product(network, index string) (productstypes.Product, error) {
	q := `
		SELECT
			network,
			index,
			parent,
			parent_chain,
			has_children,
			name,
			description,
			attributes,
			images,
			tags,
			p_type,
			class,
			creator,
			active
		FROM products_products
		WHERE network = $1 AND index = $2
	`

	var pr productstypes.Product
	if err := db.Sqlx.Select(&pr, q, network, index); err != nil {
		return productstypes.Product{}, fmt.Errorf("error while getting product: %s", err)
	}

	return pr, nil
}
