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

package database

import (
	"fmt"
	"github.com/forbole/bdjuno/v3/database/types"
	econtypes "github.com/villagelabsco/villaged/x/economics/types"
)

func (db *Db) SaveOrUpdateEconomicsNetworkEnabled(network string, val bool) error {
	stmt := `
		INSERT INTO economics_network_enabled (network, active)
		VALUES ($1, $2)
		ON CONFLICT (network) DO UPDATE SET active = $2;
	`

	if _, err := db.Sql.Exec(stmt, network, val); err != nil {
		return fmt.Errorf("error while storing economics network enabled: %s", err)
	}

	return nil
}

func (db *Db) SaveEconomicsTransactionHook(hook *econtypes.TransactionHook) error {
	stmt := `
		INSERT INTO economics_transaction_hooks (network, index, trigger, type, name_id, description, params)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	h, err := types.DbEconomicsTransactionHook{}.FromProto(hook)
	if err != nil {
		return fmt.Errorf("error while converting economics transaction hook: %s", err)
	}
	if _, err := db.Sql.Exec(stmt,
		h.Network,
		h.Index,
		h.Type,
		h.NameId,
		h.Description,
		h.Params); err != nil {
		return fmt.Errorf("error while storing economics transaction hook: %s", err)
	}

	return nil
}

func (db *Db) SaveEconomicsScheduledHook(hook *econtypes.ScheduledHook) error {
	stmt := `
		INSERT INTO economics_scheduled_hooks (network, index, type, name_id, description, cron_rule, dependencies, auto_trigger, params, last_executed_timestamp, last_executed_block) 
		VALUES ($1, $2, $3, $4, $5 ,$6, $7, $8, $9, $10, $11);
	`

	h, err := types.DbEconomicsScheduledHook{}.FromProto(hook)
	if err != nil {
		return fmt.Errorf("error while converting economics scheduled hook: %s", err)
	}

	if _, err := db.Sql.Exec(stmt,
		h.Network,
		h.Index,
		h.Type,
		h.NameId,
		h.Description,
		h.CronRule,
		h.Dependencies,
		h.AutoTrigger,
		h.Params,
		h.LastExecutedTimestamp,
		h.LastExecutedBlock); err != nil {
		return fmt.Errorf("error while storing economics scheduled hook: %s", err)
	}

	return nil
}

func (db *Db) RemoveEconomicsTransactionHook(network string, idx uint64) error {
	stmt := `
		DELETE FROM economics_transaction_hooks
		WHERE network = $1 AND index = $2;
	`

	if _, err := db.Sql.Exec(stmt, network, idx); err != nil {
		return fmt.Errorf("error while removing economics transaction hook: %s", err)
	}

	return nil
}

func (db *Db) RemoveEconomicsScheduledHook(network string, idx uint64) error {
	stmt := `
		DELETE FROM economics_scheduled_hooks
		WHERE network = $1 AND index = $2;
	`

	if _, err := db.Sql.Exec(stmt, network, idx); err != nil {
		return fmt.Errorf("error while removing economics scheduled hook: %s", err)
	}

	return nil
}

func (db *Db) SaveEconomicsScheduledHookManualTrigger(msg *econtypes.MsgTriggerScheduledHooks) error {
	stmt := `
		INSERT INTO economics_scheduled_hooks_manual_triggers (creator, network, hook_idxs)
		VALUES ($1, $2, $3);
	`

	h, err := types.DbEconomicsScheduledHookManualTrigger{}.FromProto(msg)
	if err != nil {
		return fmt.Errorf("error while converting economics scheduled hook execution: %s", err)
	}

	if _, err := db.Sql.Exec(stmt,
		h.Creator,
		h.Network,
		h.HookIdxs); err != nil {
		return fmt.Errorf("error while storing economics scheduled hook execution: %s", err)
	}

	return nil
}
