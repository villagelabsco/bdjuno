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

type DbIdentityAccountNetworks struct {
	Index    string             `db:"index"`
	Networks sqlxtypes.JSONText `db:"networks"`
}

func (vn DbIdentityAccountNetworks) FromProto(n *identitytypes.AccountNetworks) (DbIdentityAccountNetworks, error) {
	networks, err := json.Marshal(n.Networks)
	if err != nil {
		return DbIdentityAccountNetworks{}, fmt.Errorf("error marshalling networks: %v", err)
	}

	return DbIdentityAccountNetworks{
		Index:    n.Index,
		Networks: networks,
	}, nil
}

func (vn DbIdentityAccountNetworks) ToProto() (*identitytypes.AccountNetworks, error) {
	var networks []string
	err := json.Unmarshal(vn.Networks, &networks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling networks: %v", err)
	}

	return &identitytypes.AccountNetworks{
		Index:    vn.Index,
		Networks: networks,
	}, nil
}
