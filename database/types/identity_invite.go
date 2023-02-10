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
	identitytypes "github.com/villagelabsco/village/x/identity/types"
	"strings"
)

type DbIdentityInvite struct {
	Network          string `db:"network"`
	Challenge        string `db:"challenge"`
	Registered       bool   `db:"registered"`
	ConfirmedAccount string `db:"confirmed_account"`
	InviteCreator    string `db:"invite_creator"`
	HumanId          string `db:"human_id"`
	GivenRoles       string `db:"given_roles"`
}

func (fa DbIdentityInvite) FromProto(network string, i *identitytypes.Invite) DbIdentityInvite {
	return DbIdentityInvite{
		Network:          network,
		Challenge:        i.Challenge,
		Registered:       i.Registered,
		ConfirmedAccount: i.ConfirmedAccount,
		InviteCreator:    i.InviteCreator,
		HumanId:          i.HumanId,
		GivenRoles:       strings.Join(i.GivenRoles, ","),
	}
}
