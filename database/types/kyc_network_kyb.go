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

import kyctypes "github.com/villagelabs/villaged/x/kyc/types"

type DbKycNetworkKyb struct {
	Index     string `db:"index"`
	Status    uint64 `db:"status"`
	DataHash  string `db:"data_hash"`
	Timestamp uint64 `db:"timestamp"`
	Metadata  string `db:"metadata"`
}

func (DbKycNetworkKyb) FromProto(kyb *kyctypes.NetworkKyb) DbKycNetworkKyb {
	return DbKycNetworkKyb{
		Index:     kyb.Index,
		Status:    kyb.Status,
		DataHash:  kyb.DataHash,
		Timestamp: kyb.Timestamp,
		Metadata:  kyb.Metadata,
	}
}
