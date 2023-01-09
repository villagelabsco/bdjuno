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
	classestypes "github.com/villagelabs/villaged/x/classes/types"
)

func (db *Db) InsertClass(c *classestypes.Class) error {
	stmt := `
		INSERT INTO classes_classes (network, index, name, description, parent, parent_chain, has_children, creator) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	if _, err := db.Sql.Exec(stmt,
		c.NetworkID,
		c.Index,
		c.Name,
		c.Description,
		c.Parent,
		c.ParentChain,
		c.HasChildren,
		c.Creator); err != nil {
		return fmt.Errorf("error while storing class: %s", err)
	}
	return nil
}

func (db *Db) UpdateClass(c *classestypes.Class) error {
	stmt := `
		UPDATE classes_classes 
		SET 
			name = $1,
			description = $2
		WHERE network = $3 AND index = $4`

	if _, err := db.Sql.Exec(stmt,
		c.Name,
		c.Description,
		c.NetworkID,
		c.Index); err != nil {
		return fmt.Errorf("error while updating class: %s", err)
	}
	return nil
}

func (db *Db) UpdateClassHasChildren(network, index string, val bool) error {
	stmt := `
		UPDATE classes_classes 
		SET 
			has_children = $1
		WHERE network = $2 AND index = $3`

	if _, err := db.Sql.Exec(stmt, val, network, index); err != nil {
		return fmt.Errorf("error while updating class has children: %s", err)
	}
	return nil
}

func (db *Db) DeleteClass(network, index string) error {
	stmt := `
		DELETE FROM classes_classes 
		WHERE network = $1 AND index = $2`

	if _, err := db.Sql.Exec(stmt, network, index); err != nil {
		return fmt.Errorf("error while deleting class: %s", err)
	}
	return nil
}

func (db *Db) Class(network, index string) (classestypes.Class, error) {
	q := `
		SELECT (network, index, name, description, parent, parent_chain, has_children, creator) FROM classes_classes
		WHERE network = $1 AND index = $2
	`

	var c classestypes.Class
	if err := db.Sqlx.Select(&c, q, network, index); err != nil {
		return classestypes.Class{}, fmt.Errorf("error while getting class: %s", err)
	}

	return c, nil
}
