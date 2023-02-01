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
	econtypes "github.com/villagelabsco/villaged/x/economics/types"
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

func (p DbEconomicsScheduledHook) FromProto(hook *econtypes.ScheduledHook) (DbEconomicsScheduledHook, error) {
	dps, err := json.Marshal(hook.HookDependencies)
	if err != nil {
		return DbEconomicsScheduledHook{}, fmt.Errorf("error while marshalling hook dependencies: %v", err)
	}

	params, err := json.Marshal(hook.HookParams)
	if err != nil {
		return DbEconomicsScheduledHook{}, fmt.Errorf("error while marshalling hook params: %v", err)
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
		Params:                params,
		LastExecutedTimestamp: time.Unix(0, int64(hook.LastExecutedTimestamp)),
		LastExecutedBlock:     hook.LastExecutedBlock,
	}, nil
}
