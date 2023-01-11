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
	villagetypes "github.com/villagelabs/villaged/x/village/types"
)

func (db *Db) SaveNetwork(network *villagetypes.Network) error {
	stmt := `
	INSERT INTO village_network ("index", "active", "full_name", "identity_provider", "invite_only")
	VALUES ($1, $2, $3, $4, $5);`

	_, err := db.Sql.Exec(stmt, network.Index, network.Active, network.FullName, network.IdentityProvider, network.InviteOnly)
	if err != nil {
		return fmt.Errorf("error while inserting network: %s", err)
	}

	return nil
}

func (db *Db) UpdateNetwork(network *villagetypes.Network) error {
	stmt := `
	UPDATE village_network AS vn
	SET active = $2,
		full_name = $3,
		identity_provider = $4,
		invite_only = $5
	WHERE vn.index = $1;`

	_, err := db.Sql.Exec(stmt, network.Index, network.Active, network.FullName, network.IdentityProvider, network.InviteOnly)
	if err != nil {
		return fmt.Errorf("error while updating network: %s", err)
	}

	return nil
}

func (db *Db) UserNetworks(index string) (*villagetypes.UserNetworks, error) {
	q := `
	SELECT ("index", "networks") FROM village_user_networks 
	WHERE "index" = $1;`

	var userNetworks villagetypes.UserNetworks
	err := db.Sqlx.Select(&userNetworks, q, index)
	if err != nil {
		return nil, fmt.Errorf("error while getting user networks: %s", err)
	}

	return &userNetworks, nil
}

func (db *Db) SaveUserNetworks(userNetworks *villagetypes.UserNetworks) error {
	stmt := `
	INSERT INTO village_user_networks ("index", "networks")
	VALUES ($1, $2);`

	_, err := db.Sql.Exec(stmt, userNetworks.Index, userNetworks.Networks)
	if err != nil {
		return fmt.Errorf("error while inserting user networks: %s", err)
	}

	return nil
}

func (db *Db) UpdateUserNetworks(userNetworks *villagetypes.UserNetworks) error {
	stmt := `
	UPDATE village_user_networks AS vun
	SET networks = $2
	WHERE vun.index = $1;`

	_, err := db.Sql.Exec(stmt, userNetworks.Index, userNetworks.Networks)
	if err != nil {
		return fmt.Errorf("error while updating user networks: %s", err)
	}

	return nil
}
