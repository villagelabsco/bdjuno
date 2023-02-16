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
	identitytypes "github.com/villagelabsco/village/x/identity/types"
)

type DbIdentityProvider struct {
	Index               string             `db:"index"`
	AdminAccounts       sqlxtypes.JSONText `db:"admin_accounts"`
	ProviderAccounts    sqlxtypes.JSONText `db:"provider_accounts"`
	AssetMinterAccounts sqlxtypes.JSONText `db:"asset_minter_accounts"`
	AssetBurnerAccounts sqlxtypes.JSONText `db:"asset_burner_accounts"`
}

func (d DbIdentityProvider) FromProto(provider identitytypes.IdentityProvider) (DbIdentityProvider, error) {
	aa, err := json.Marshal(provider.AdminAccounts)
	if err != nil {
		return DbIdentityProvider{}, fmt.Errorf("error marshalling AdminAccounts: %s", err)
	}
	pa, err := json.Marshal(provider.ProviderAccounts)
	if err != nil {
		return DbIdentityProvider{}, fmt.Errorf("error marshalling ProviderAccounts: %s", err)
	}
	ama, err := json.Marshal(provider.AssetMinterAccounts)
	if err != nil {
		return DbIdentityProvider{}, fmt.Errorf("error marshalling AssetMinterAccounts: %s", err)
	}
	aba, err := json.Marshal(provider.AssetBurnerAccounts)
	if err != nil {
		return DbIdentityProvider{}, fmt.Errorf("error marshalling AssetBurnerAccounts: %s", err)
	}

	return DbIdentityProvider{
		Index:               provider.Index,
		AdminAccounts:       aa,
		ProviderAccounts:    pa,
		AssetMinterAccounts: ama,
		AssetBurnerAccounts: aba,
	}, nil
}
