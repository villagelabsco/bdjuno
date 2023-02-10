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

type DbIdentityAccountLinkProposal struct {
	Index                        string             `db:"index"`
	ProposerAccount              string             `db:"proposer_account"`
	HumanId                      string             `db:"human_id"`
	SetAsPrimaryWalletForNetwork string             `db:"set_as_primary_wallet_for_network"`
	Deposit                      sqlxtypes.JSONText `db:"deposit"`
}

func (DbIdentityAccountLinkProposal) FromProto(alp *identitytypes.AccountLinkProposal) (DbIdentityAccountLinkProposal, error) {
	deposit, err := json.Marshal(alp.Deposit)
	if err != nil {
		return DbIdentityAccountLinkProposal{}, fmt.Errorf("error while marshaling deposit: %s", err)
	}
	
	return DbIdentityAccountLinkProposal{
		Index:                        alp.Index,
		ProposerAccount:              alp.ProposerAccount,
		HumanId:                      alp.HumanId,
		SetAsPrimaryWalletForNetwork: alp.SetAsPrimaryWalletForNetwork,
		Deposit:                      deposit,
	}, nil
}
