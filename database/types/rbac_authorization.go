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
	"encoding/json"
	"fmt"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	rbactypes "github.com/villagelabsco/village/x/rbac/types"
)

type DbRbacAuthorization struct {
	Index         string             `db:"index"`
	Messages      sqlxtypes.JSONText `db:"messages"`
	Metadata      string             `db:"metadata"`
	GroupId       uint64             `db:"group_id"`
	RoleAdmins    sqlxtypes.JSONText `db:"role_admins"`
	RoleDelegates sqlxtypes.JSONText `db:"role_delegates"`
}

func (DbRbacAuthorization) FromProto(au *rbactypes.Authorizations) (DbRbacAuthorization, error) {
	msg, err := json.Marshal(au.Messages)
	if err != nil {
		return DbRbacAuthorization{}, fmt.Errorf("error marshalling messages: %v", err)
	}
	roleAdmins, err := json.Marshal(au.RoleAdmins)
	if err != nil {
		return DbRbacAuthorization{}, fmt.Errorf("error marshalling role admins: %v", err)
	}
	roleDelegates, err := json.Marshal(au.RoleDelegates)
	if err != nil {
		return DbRbacAuthorization{}, fmt.Errorf("error marshalling role delegates: %v", err)
	}
	return DbRbacAuthorization{
		Index:         au.Index,
		Messages:      msg,
		Metadata:      au.Metadata,
		GroupId:       au.GroupId,
		RoleAdmins:    roleAdmins,
		RoleDelegates: roleDelegates,
	}, nil
}
