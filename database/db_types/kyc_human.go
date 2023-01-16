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

import sqlxtypes "github.com/jmoiron/sqlx/types"

type DbKycHuman struct {
	Index                string             `db:"index"`
	VnsDomain            string             `db:"vns_domain"`
	Accounts             sqlxtypes.JSONText `db:"accounts"`
	NetworkPrimaryWallet sqlxtypes.JSONText `db:"network_primary_wallet"`
}
