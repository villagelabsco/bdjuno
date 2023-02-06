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
	identitytypes "github.com/villagelabsco/villaged/x/identity/types"
	"time"
)

type DbIdentityKycStatus struct {
	HumanId          string    `db:"human_id"`
	IdentityProvider string    `db:"identity_provider"`
	Status           time.Time `db:"status"`
	DataHash         string    `db:"data_hash"`
	Timestamp        uint64    `db:"timestamp"`
}

func (DbIdentityKycStatus) FromProto(provider string, st *identitytypes.KycStatus) DbIdentityKycStatus {
	return DbIdentityKycStatus{
		HumanId:          st.HumanId,
		IdentityProvider: provider,
		Status:           time.Unix(int64(st.Status), 0),
		DataHash:         st.DataHash,
		Timestamp:        st.Timestamp,
	}
}
