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
	"time"
)

type DbIdentityNetworkKyb struct {
	Index     string    `db:"index"`
	Status    uint64    `db:"status"`
	DataHash  string    `db:"data_hash"`
	Timestamp time.Time `db:"timestamp"`
	Metadata  string    `db:"metadata"`
}

func (DbIdentityNetworkKyb) FromProto(kyb *identitytypes.NetworkKyb) DbIdentityNetworkKyb {
	return DbIdentityNetworkKyb{
		Index:     kyb.Index,
		Status:    kyb.Status,
		DataHash:  kyb.DataHash,
		Timestamp: time.Unix(int64(kyb.Timestamp), 0),
		Metadata:  kyb.Metadata,
	}
}
