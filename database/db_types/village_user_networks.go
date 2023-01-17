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

package db_types

import (
	"encoding/json"
	"fmt"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	villagetypes "github.com/villagelabs/villaged/x/village/types"
)

type DbVillageUserNetworks struct {
	Index    string             `db:"index"`
	Networks sqlxtypes.JSONText `db:"networks"`
}

func (vn DbVillageUserNetworks) FromProto(n *villagetypes.UserNetworks) (DbVillageUserNetworks, error) {
	networks, err := json.Marshal(n.Networks)
	if err != nil {
		return DbVillageUserNetworks{}, fmt.Errorf("error marshalling networks: %v", err)
	}

	return DbVillageUserNetworks{
		Index:    n.Index,
		Networks: networks,
	}, nil
}
