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
	"github.com/gogo/protobuf/proto"
	sqlxtypes "github.com/jmoiron/sqlx/types"
	econtypes "github.com/villagelabsco/village/x/economics/types"
	"time"
)

type DbEconomicsScheduledHook struct {
	Network               string                      `db:"network"`
	Index                 uint64                      `db:"index"`
	Type                  econtypes.ScheduledHookType `db:"type"`
	NameId                string                      `db:"name_id"`
	Description           string                      `db:"description"`
	CronRule              string                      `db:"cron_rule"`
	Dependencies          sqlxtypes.JSONText          `db:"dependencies"`
	AutoTrigger           bool                        `db:"auto_trigger"`
	Params                sqlxtypes.JSONText          `db:"params"`
	LastExecutedTimestamp time.Time                   `db:"last_executed_timestamp"`
	LastExecutedBlock     uint64                      `db:"last_executed_block"`
}

func (h DbEconomicsScheduledHook) FromProto(hook *econtypes.ScheduledHook) (DbEconomicsScheduledHook, error) {
	var p proto.Unmarshaler
	if hook.HookParams != nil {
		switch hook.HookType {
		case econtypes.ScheduledHookType_SCHEDHOOKTYPE_TRANSFER_DENOM_TO_SHAREHOLDERS:
			p = new(econtypes.HookParamsTransferDenomToShareholders)
		case econtypes.ScheduledHookType_SCHEDHOOKTYPE_AUTO_SWAP_DENOM:
			p = new(econtypes.HookParamsAutoSwapDenom)
		case econtypes.ScheduledHookType_SCHEDHOOKTYPE_DECAY_DENOM_FOR_INACTIVE_ACCOUNT:
			p = new(econtypes.HookParamsDecayDenomForInactiveAccounts)
		case econtypes.ScheduledHookType_SCHEDHOOKTYPE_RECURRING_MINT_TOKEN:
			p = new(econtypes.HookParamsRecurringMintToken)
		}
	}

	bParams, err := json.Marshal(p)
	if err != nil {
		return DbEconomicsScheduledHook{}, fmt.Errorf("error while marshalling hook params: %v", err)
	}

	dps, err := json.Marshal(hook.HookDependencies)
	if err != nil {
		return DbEconomicsScheduledHook{}, fmt.Errorf("error while marshalling hook dependencies: %v", err)
	}

	return DbEconomicsScheduledHook{
		Network:               hook.Network,
		Index:                 hook.HookIdx,
		Type:                  hook.HookType,
		NameId:                hook.NameId,
		Description:           hook.HookDescription,
		CronRule:              hook.CronRule,
		Dependencies:          dps,
		AutoTrigger:           hook.AutoTrigger,
		Params:                bParams,
		LastExecutedTimestamp: time.Unix(0, int64(hook.LastExecutedTimestamp)),
		LastExecutedBlock:     hook.LastExecutedBlock,
	}, nil
}
