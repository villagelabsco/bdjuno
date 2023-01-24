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
	identitytypes "github.com/villagelabs/villaged/x/identity/types"
)

type DbIdentityNetwork struct {
	Index            string `db:"index"`
	Active           bool   `db:"active"`
	FullName         string `db:"full_name"`
	IdentityProvider string `db:"identity_provider"`
	InviteOnly       bool   `db:"invite_only"`
}

func (vn DbIdentityNetwork) FromProto(n *identitytypes.Network) DbIdentityNetwork {
	return DbIdentityNetwork{
		Index:            n.Index,
		Active:           n.Active,
		FullName:         n.FullName,
		IdentityProvider: n.IdentityProvider,
		InviteOnly:       n.InviteOnly,
	}
}
