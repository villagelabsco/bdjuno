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
	rbactypes "github.com/villagelabs/villaged/x/rbac/types"
)

func (db *Db) Authorization(index string) (*rbactypes.Authorizations, error) {
	q := `
		SELECT (index, messages, metadata, group_id, role_admins, role_delegates)
		FROM rbac_authorizations
		WHERE index = $1
	`

	var au rbactypes.Authorizations
	if err := db.Sqlx.Select(&au, q, index); err != nil {
		return nil, fmt.Errorf("error while getting authorization: %s", err)
	}

	return &au, nil
}

func (db *Db) SaveAuthorization(au *rbactypes.Authorizations) error {
	stmt := `
		INSERT INTO rbac_authorizations (index, messages, metadata, group_id, role_admins, role_delegates) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	if _, err := db.Sql.Exec(stmt,
		au.Index,
		au.Messages,
		au.Metadata,
		au.GroupId,
		au.RoleAdmins,
		au.RoleDelegates); err != nil {
		return fmt.Errorf("error while storing authorization: %s", err)
	}

	return nil
}
