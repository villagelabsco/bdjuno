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
	"encoding/json"
	"fmt"
	"github.com/villagelabsco/bdjuno/v4/database/types"
	rbactypes "github.com/villagelabsco/village/x/rbac/types"
)

func (db *Db) SaveRbacAuthorization(au *rbactypes.Authorizations, members []*rbactypes.MemberUpdates) error {
	stmt := `
		INSERT INTO rbac_authorizations (index, messages, metadata, group_id, role_admins, role_delegates, members) 
		VALUES ($1, $2, $3, $4, $5, $6, $7);      
	`

	dbau, err := types.DbRbacAuthorization{}.FromProto(au, members)
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
		dbau.Members,
	); err != nil {
		return fmt.Errorf("error while storing authorization: %s", err)
	}

	return nil
}

func (db *Db) UpdateRbacRoleAdmins(au *rbactypes.Authorizations) error {
	stmt := `
		UPDATE rbac_authorizations SET role_admins = $2 WHERE index = $1;
	`

	roleAdmins, err := json.Marshal(au.RoleAdmins)
	if err != nil {
		return fmt.Errorf("error marshalling role admins: %v", err)
	}

	if _, err := db.SQL.Exec(stmt,
		au.Index,
		roleAdmins,
	); err != nil {
		return fmt.Errorf("error while updating role admins: %s", err)
	}

	return nil
}

func (db *Db) UpdateRbacRoleDelegates(au *rbactypes.Authorizations) error {
	stmt := `
		UPDATE rbac_authorizations SET role_delegates = $2 WHERE index = $1;
	`

	roleDelegates, err := json.Marshal(au.RoleDelegates)
	if err != nil {
		return fmt.Errorf("error marshalling role delegates: %v", err)
	}

	if _, err := db.SQL.Exec(stmt,
		au.Index,
		roleDelegates,
	); err != nil {
		return fmt.Errorf("error while updating role delegates: %s", err)
	}

	return nil
}

func (db *Db) UpdateRbacRoleMessages(au *rbactypes.Authorizations) error {
	stmt := `
		UPDATE rbac_authorizations SET messages = $2 WHERE index = $1;
	`

	messages, err := json.Marshal(au.Messages)
	if err != nil {
		return fmt.Errorf("error marshalling messages: %v", err)
	}

	if _, err := db.SQL.Exec(stmt,
		au.Index,
		messages,
	); err != nil {
		return fmt.Errorf("error while updating messages: %s", err)
	}

	return nil
}

func (db *Db) UpdateRbacRoleMembers(index string, members []*rbactypes.MemberUpdates) error {
	stmt := `
		UPDATE rbac_authorizations SET members = $2 WHERE index = $1;
	`

	membersJson, err := json.Marshal(members)
	if err != nil {
		return fmt.Errorf("error marshalling members: %v", err)
	}

	if _, err := db.SQL.Exec(stmt,
		index,
		membersJson,
	); err != nil {
		return fmt.Errorf("error while updating members: %s", err)
	}

	return nil
}
