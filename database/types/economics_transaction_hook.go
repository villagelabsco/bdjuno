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
)

type DbEconomicsTransactionHook struct {
	Network     string                        `db:"network"`
	Index       uint64                        `db:"index"`
	Trigger     econtypes.TriggerType         `db:"trigger"`
	Type        econtypes.TransactionHookType `db:"type"`
	NameId      string                        `db:"name_id"`
	Description string                        `db:"description"`
	Params      sqlxtypes.JSONText            `db:"params"`
}

func (DbEconomicsTransactionHook) FromProto(h *econtypes.TransactionHook) (DbEconomicsTransactionHook, error) {
	bParams, err := json.Marshal(h.HookParams)
	if err != nil {
		return DbEconomicsTransactionHook{}, fmt.Errorf("error marshalling hook params: %v", err)
	}

	return DbEconomicsTransactionHook{
		Network:     h.Network,
		Index:       h.HookIdx,
		Trigger:     h.Trigger,
		Type:        h.HookType,
		NameId:      h.NameId,
		Description: h.HookDescription,
		Params:      bParams,
	}, nil
}
