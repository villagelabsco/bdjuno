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
	"github.com/forbole/bdjuno/v3/database/db_types"
	kyctypes "github.com/villagelabs/villaged/x/kyc/types"
	"strings"
)

func (db *Db) SaveInvite(network string, invite *kyctypes.Invite) error {
	stmt := `
	INSERT INTO kyc_invites ("network", "challenge", "registered", "confirmed_account", "invite_creator", "human_id", "given_roles") 
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	inv := db_types.DbKycInvite{}.FromProto(network, invite)
	_, err := db.Sql.Exec(stmt, network,
		inv.Challenge,
		inv.Registered,
		inv.ConfirmedAccount,
		inv.InviteCreator,
		inv.HumanId,
		inv.GivenRoles,
	)
	if err != nil {
		return fmt.Errorf("error while storing invite: %s", err)
	}

	return err
}

func (db *Db) SaveMultipleInvites(network string, invites []*kyctypes.Invite) error {
	stmt := `
	INSERT INTO kyc_invites ("network", "challenge", "registered", "confirmed_account", "invite_creator", "human_id", "given_roles")
	VALUES %s;`

	values := make([]string, 0, len(invites))
	for _, invite := range invites {
		values = append(values,
			fmt.Sprintf("('%s', '%s', '%t', '%s', '%s', '%s', '%s')",
				network,
				invite.Challenge,
				false,
				invite.ConfirmedAccount,
				invite.InviteCreator,
				invite.HumanId,
				strings.Join(invite.GivenRoles, ","),
			))
	}

	stmt = fmt.Sprintf(stmt, strings.Join(values, ","))
	_, err := db.Sql.Exec(stmt)
	if err != nil {
		return fmt.Errorf("error while storing invites: %s", err)
	}

	return nil
}

func (db *Db) UpdateInvite(network string, challenge string, confirmedAccount string) error {
	stmt := `
	UPDATE kyc_invites
	SET 
	    "registered" = $1, 
	    "confirmed_account" = $2 
	WHERE "network" = $3 AND "challenge" = $4;`

	_, err := db.Sql.Exec(stmt, true, confirmedAccount, network, challenge)
	if err != nil {
		return fmt.Errorf("error while updating invite: %s", err)
	}

	return nil
}

func (db *Db) DeleteInvite(network string, challenge string) error {
	stmt := `
		DELETE FROM kyc_invites
		WHERE "network" = $1 AND "challenge" = $2;
	`

	_, err := db.Sql.Exec(stmt, network, challenge)
	if err != nil {
		return fmt.Errorf("error while deleting invite: %s", err)
	}

	return nil
}

func (db *Db) SaveOrUpdateIdentityProvider(ip kyctypes.IdentityProvider) error {
	stmt := `
		INSERT INTO kyc_identity_provider (index, admin_accounts, provider_accounts, asset_minter_accounts, asset_burner_accounts) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (index) DO
		UPDATE 
		    SET
		        admin_accounts = $2,
		        provider_accounts = $3,
		        asset_minter_accounts = $4,
		        asset_burner_accounts = $5;
	`

	dbip, err := db_types.DbKycIdentityProvider{}.FromProto(ip)
	if err != nil {
		return fmt.Errorf("error while converting identity provider: %s", err)
	}

	_, err = db.Sql.Exec(stmt,
		dbip.Index,
		dbip.AdminAccounts,
		dbip.ProviderAccounts,
		dbip.AssetMinterAccounts,
		dbip.AssetBurnerAccounts,
	)
	if err != nil {
		return fmt.Errorf("error while storing identity provider: %s", err)
	}

	return nil
}

func (db *Db) SaveOrUpdateNetworkKyb(kyb *kyctypes.NetworkKyb) error {
	stmt := `
		INSERT INTO kyc_network_kyb (index, status, data_hash, timestamp, metadata) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (index) DO
		UPDATE 
		    SET
		        status = $2,
		        data_hash = $3,
		        timestamp = $4,
		        metadata = $5;
	`

	dbkyb := db_types.DbKycNetworkKyb{}.FromProto(kyb)
	_, err := db.Sql.Exec(stmt,
		dbkyb.Index,
		dbkyb.Status,
		dbkyb.DataHash,
		dbkyb.Timestamp,
		dbkyb.Metadata,
	)
	if err != nil {
		return fmt.Errorf("error while storing network kyb: %s", err)
	}

	return nil
}
