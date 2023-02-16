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

import identitytypes "github.com/villagelabsco/village/x/identity/types"

type DbIdentityAccount struct {
	Index      string `db:"index"`
	HumanId    string `db:"human_id"`
	PrivateAcc bool   `db:"private_acc"`
}

func (DbIdentityAccount) FromProto(acc *identitytypes.Account) DbIdentityAccount {
	return DbIdentityAccount{
		Index:      acc.Index,
		HumanId:    acc.HumanId,
		PrivateAcc: acc.PrivateAcc,
	}
}
