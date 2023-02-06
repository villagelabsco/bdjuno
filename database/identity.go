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
	identitytypes "github.com/villagelabsco/villaged/x/identity/types"
	"strings"
)

func (db *Db) SaveIdentityNetwork(network *identitytypes.Network) error {
	stmt := `
	INSERT INTO identity_networks ("index", "active", "full_name", "identity_provider", "invite_only")
	VALUES ($1, $2, $3, $4, $5);`

	n := types.DbIdentityNetwork{}.FromProto(network)
	_, err := db.SQL.Exec(stmt, n.Index, n.Active, n.FullName, n.IdentityProvider, n.InviteOnly)
	if err != nil {
		return fmt.Errorf("error while inserting network: %s", err)
	}

	return nil
}

func (db *Db) UpdateIdentityNetwork(network *identitytypes.Network) error {
	stmt := `
	UPDATE identity_networks AS vn
	SET active = $2,
		full_name = $3,
		identity_provider = $4,
		invite_only = $5
	WHERE vn.index = $1;`

	n := types.DbIdentityNetwork{}.FromProto(network)
	_, err := db.SQL.Exec(stmt, n.Index, n.Active, n.FullName, n.IdentityProvider, n.InviteOnly)
	if err != nil {
		return fmt.Errorf("error while updating network: %s", err)
	}

	return nil
}

func (db *Db) IdentityAccountNetworks(index string) (*identitytypes.AccountNetworks, error) {
	q := `
	SELECT (index, networks) FROM identity_account_networks 
	WHERE index = $1
	LIMIT 1;`

	var userNetworks types.DbIdentityAccountNetworks
	err := db.SQL.Select(&userNetworks, q, index)
	if err != nil {
		return nil, fmt.Errorf("error while getting user networks: %s", err)
	}

	res, err := userNetworks.ToProto()
	if err != nil {
		return nil, fmt.Errorf("error while converting user networks: %s", err)
	}

	return res, nil
}

func (db *Db) SaveOrAppendIdentityAccountNetworks(index, network string) error {
	stmt := `
	INSERT INTO identity_account_networks as ian ("index", "networks") 
	VALUES ($1, $2) 
	ON CONFLICT (index) DO
	UPDATE
		SET
		    networks = ian.networks || $2::jsonb;
	`

	un, err := types.DbIdentityAccountNetworks{}.FromProto(
		&identitytypes.AccountNetworks{
			Index:    index,
			Networks: []string{network},
		})
	if err != nil {
		return fmt.Errorf("error while converting user networks: %s", err)
	}
	_, err = db.SQL.Exec(stmt, un.Index, un.Networks)
	if err != nil {
		return fmt.Errorf("error while inserting user networks: %s", err)
	}

	return nil
}

func (db *Db) SaveIdentityInvite(network string, invite *identitytypes.Invite) error {
	stmt := `
	INSERT INTO identity_invites ("network", "challenge", "registered", "confirmed_account", "invite_creator", "human_id", "given_roles") 
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	inv := types.DbIdentityInvite{}.FromProto(network, invite)
	_, err := db.SQL.Exec(stmt, network,
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

func (db *Db) SaveMultipleIdentityInvites(network string, invites []*identitytypes.Invite) error {
	stmt := `
		INSERT INTO identity_invites ("network", "challenge", "registered", "confirmed_account", "invite_creator", "human_id", "given_roles")
		VALUES %s;
		`

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
	_, err := db.SQL.Exec(stmt)
	if err != nil {
		return fmt.Errorf("error while storing invites: %s", err)
	}

	return nil
}

func (db *Db) UpdateIdentityInvite(network, challenge, confirmedAccount string, registered bool) error {
	stmt := `
	UPDATE identity_invites
	SET 
	    "registered" = $1, 
	    "confirmed_account" = $2 
	WHERE "network" = $3 AND "challenge" = $4;`

	_, err := db.SQL.Exec(stmt, true, confirmedAccount, network, challenge)
	if err != nil {
		return fmt.Errorf("error while updating invite: %s", err)
	}

	return nil
}

func (db *Db) DeleteIdentityInvite(network string, challenge string) error {
	stmt := `
		DELETE FROM identity_invites
		WHERE "network" = $1 AND "challenge" = $2;
	`

	_, err := db.SQL.Exec(stmt, network, challenge)
	if err != nil {
		return fmt.Errorf("error while deleting invite: %s", err)
	}

	return nil
}

func (db *Db) SaveOrUpdateIdentityProvider(ip identitytypes.IdentityProvider) error {
	stmt := `
		INSERT INTO identity_providers (index, admin_accounts, provider_accounts, asset_minter_accounts, asset_burner_accounts) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (index) DO
		UPDATE 
		    SET
		        admin_accounts = $2,
		        provider_accounts = $3,
		        asset_minter_accounts = $4,
		        asset_burner_accounts = $5;
	`

	dbip, err := types.DbIdentityProvider{}.FromProto(ip)
	if err != nil {
		return fmt.Errorf("error while converting identity provider: %s", err)
	}

	_, err = db.SQL.Exec(stmt,
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

func (db *Db) SaveOrUpdateIdentityNetworkKyb(kyb *identitytypes.NetworkKyb) error {
	stmt := `
		INSERT INTO identity_network_kyb (index, status, data_hash, timestamp, metadata) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (index) DO
		UPDATE 
		    SET
		        status = $2,
		        data_hash = $3,
		        timestamp = $4,
		        metadata = $5;
	`

	dbkyb := types.DbIdentityNetworkKyb{}.FromProto(kyb)
	_, err := db.SQL.Exec(stmt,
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

func (db *Db) SaveOrUpdateIdentityHuman(human *identitytypes.Human) error {
	stmt := `
		INSERT INTO identity_humans (index, vns_domain, accounts, networks, network_primary_wallet) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (index) DO
		UPDATE
		    SET
		        vns_domain = $2,
		        accounts = $3,
		        networks = $4,
		        network_primary_wallet = $5;
	`

	dbHuman, err := types.DbIdentityHuman{}.FromProto(human)
	if err != nil {
		return fmt.Errorf("error while converting human: %s", err)
	}
	_, err = db.SQL.Exec(stmt,
		dbHuman.Index,
		dbHuman.VnsDomain,
		dbHuman.Accounts,
		dbHuman.Networks,
		dbHuman.NetworkPrimaryWallet,
	)
	if err != nil {
		return fmt.Errorf("error while storing human: %s", err)
	}

	return nil
}

func (db *Db) SaveOrUpdateIdentityAccount(account *identitytypes.Account) error {
	stmt := `
		INSERT INTO identity_accounts (index, human_id, private_acc) 
		VALUES ($1, $2, $3)
		ON CONFLICT (index) DO 
		UPDATE
			SET
			    human_id = $2,
			    private_acc = $3
	`

	dbAcc := types.DbIdentityAccount{}.FromProto(account)
	_, err := db.SQL.Exec(stmt, dbAcc.Index, dbAcc.HumanId, dbAcc.PrivateAcc)
	if err != nil {
		return fmt.Errorf("error while storing account: %s", err)
	}

	return nil
}

func (db *Db) SaveOrUpdateKycStatus(provider string, status *identitytypes.KycStatus) error {
	stmt := `
		INSERT INTO identity_kyc_statuses (human_id, identity_provider, status, data_hash, timestamp) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (human_id) DO
		UPDATE
		    SET
		        human_id = $1,
		        identity_provider = $2,
		        status = $3,
		        data_hash = $4,
		        timestamp = $5;
	`

	dbSt := types.DbIdentityKycStatus{}.FromProto(provider, status)
	_, err := db.SQL.Exec(stmt,
		dbSt.HumanId,
		dbSt.IdentityProvider,
		dbSt.Status,
		dbSt.DataHash,
		dbSt.Timestamp,
	)
	if err != nil {
		return fmt.Errorf("error while storing kyc status: %s", err)
	}

	return nil
}
