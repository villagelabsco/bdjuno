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
	sqlxtypes "github.com/jmoiron/sqlx/types"
	econtypes "github.com/villagelabsco/villaged/x/economics/types"
)

type DbEconomicsScheduledHookManualTrigger struct {
	Id       uint64             `db:"id"`
	Creator  string             `db:"creator"`
	Network  string             `db:"network"`
	HookIdxs sqlxtypes.JSONText `db:"hook_idxs"`
}

func (DbEconomicsScheduledHookManualTrigger) FromProto(h *econtypes.MsgTriggerScheduledHooks) (DbEconomicsScheduledHookManualTrigger, error) {
	idxs, err := json.Marshal(h.HookIdxs)
	if err != nil {
		return DbEconomicsScheduledHookManualTrigger{}, err
	}

	return DbEconomicsScheduledHookManualTrigger{
		Creator:  h.Creator,
		Network:  h.Network,
		HookIdxs: idxs,
	}, nil
}
