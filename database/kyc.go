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
	kyctypes "github.com/villagelabs/villaged/x/kyc/types"
)

func (db *Db) InsertInvite(network string, invite *kyctypes.Invite) error {
	stmt := `
	INSERT INTO kyc_invite ("network", "challenge", "registered", "confirmed_account", "invite_creator", "human_id", "givenroles") 
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := db.Sql.Exec(stmt, network, invite.Challenge, false, invite.ConfirmedAccount, invite.InviteCreator, invite.HumanId, invite.GivenRoles)
	if err != nil {
		return fmt.Errorf("error while storing invite: %s", err)
	}

	return err
}

func (db *Db) UpdateInvite(network string, challenge string, confirmedAccount string) error {
	stmt := `
	UPDATE kyc_invite 
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
