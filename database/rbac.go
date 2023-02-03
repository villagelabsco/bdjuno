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
	rbactypes "github.com/villagelabsco/villaged/x/rbac/types"
)

func (db *Db) SaveOrUpdateAuthorization(au *rbactypes.Authorizations) error {
	stmt := `
		INSERT INTO rbac_authorizations (index, messages, metadata, group_id, role_admins, role_delegates) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	dbau, err := types.DbRbacAuthorization{}.FromProto(au)
	if err != nil {
		return fmt.Errorf("error converting rbac authorizations from proto: %v", err)
	}
	if _, err := db.SQL.Exec(stmt,
		dbau.Index,
		dbau.Messages,
		dbau.Metadata,
		dbau.GroupId,
		dbau.RoleAdmins,
		dbau.RoleDelegates,
	); err != nil {
		return fmt.Errorf("error while storing authorization: %s", err)
	}

	return nil
}
